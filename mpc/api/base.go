// Package api provides MPC API implementations
package api

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"chainup.com/go-sdk/utils"
)

// MpcConfigProvider defines the interface for accessing MPC configuration
type MpcConfigProvider interface {
	GetDomain() string
	GetAppID() string
	GetApiKey() string
	IsDebug() bool
	GetCryptoProvider() utils.CryptoProvider
	GetSignPrivateKey() *rsa.PrivateKey
}

// MpcBaseAPI provides common functionality for all MPC API implementations
type MpcBaseAPI struct {
	config         MpcConfigProvider
	httpClient     *utils.MpcHTTPClient
	cryptoProvider utils.CryptoProvider
}

// NewMpcBaseAPI creates a new MpcBaseAPI instance
func NewMpcBaseAPI(config MpcConfigProvider) *MpcBaseAPI {
	return &MpcBaseAPI{
		config:         config,
		httpClient:     utils.NewMpcHTTPClient(config.GetDomain(), config.GetAppID(), config.GetApiKey(), utils.DefaultTimeout, config.IsDebug()),
		cryptoProvider: config.GetCryptoProvider(),
	}
}

// buildRequestArgs builds the request args JSON with common parameters
func (m *MpcBaseAPI) buildRequestArgs(data map[string]interface{}) (string, error) {
	if data == nil {
		data = make(map[string]interface{})
	}

	// Add common parameters
	data["time"] = time.Now().UnixMilli()
	data["charset"] = "utf-8"

	// Convert to JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data: %w", err)
	}

	return string(jsonBytes), nil
}

// executeRequest executes an MPC API request
func (m *MpcBaseAPI) executeRequest(method, path string, data map[string]interface{}) (map[string]interface{}, error) {
	// Step 1: Build request args JSON
	rawJSON, err := m.buildRequestArgs(data)
	if err != nil {
		return nil, err
	}

	if m.config.IsDebug() {
		fmt.Printf("[MPC Request Args]: %s\n", rawJSON)
	}

	// Step 2: Encrypt with private key
	encryptedData := ""
	if m.cryptoProvider != nil {
		encrypted, err := m.cryptoProvider.EncryptWithPrivateKey(rawJSON)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt request data: %w", err)
		}
		encryptedData = encrypted

		if m.config.IsDebug() {
			if len(encryptedData) > 100 {
				fmt.Printf("[MPC Encrypted Data]: %s...\n", encryptedData[:100])
			} else {
				fmt.Printf("[MPC Encrypted Data]: %s\n", encryptedData)
			}
		}
	}

	// Step 3: Send request
	requestData := map[string]interface{}{
		"data": encryptedData,
		"app_id" : m.config.GetAppID(),
	}

	var response string

	if method == utils.HTTPMethodPost {
		response, err = m.httpClient.Post(path, requestData)
	} else {
		response, err = m.httpClient.Get(path, requestData)
	}

	if err != nil {
		return nil, err
	}

	if m.config.IsDebug() {
		fmt.Printf("[MPC Response]: %s\n", response)
	}

	// Step 4: Parse response and decrypt data if needed
	var parsedResponse map[string]interface{}
	if err := json.Unmarshal([]byte(response), &parsedResponse); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	// Check if response has encrypted data field and decrypt
	if dataField, ok := parsedResponse["data"].(string); ok && dataField != "" {
		if m.cryptoProvider != nil {
			decrypted, err := m.cryptoProvider.DecryptWithPublicKey(dataField)
			if err != nil {
				if m.config.IsDebug() {
					fmt.Printf("[MPC Decrypt Error]: %v\n", err)
				}
				// If decryption fails, might be an error response, return as-is
				return parsedResponse, nil
			}

			if m.config.IsDebug() {
				fmt.Printf("[MPC Decrypted]: %s\n", decrypted)
			}

			// Parse decrypted JSON - the decrypted data IS the full response structure
			// containing code, data, msg fields
			var decryptedResponse map[string]interface{}
			if err := json.Unmarshal([]byte(decrypted), &decryptedResponse); err != nil {
				return nil, fmt.Errorf("failed to parse decrypted data: %w", err)
			}

			// Return the decrypted response directly (it contains code, data, msg)
			return decryptedResponse, nil
		}
	}

	return parsedResponse, nil
}

// Post executes a POST request
func (m *MpcBaseAPI) Post(path string, data map[string]interface{}) (map[string]interface{}, error) {
	return m.executeRequest(utils.HTTPMethodPost, path, data)
}

// Get executes a GET request
func (m *MpcBaseAPI) Get(path string, data map[string]interface{}) (map[string]interface{}, error) {
	return m.executeRequest(utils.HTTPMethodGet, path, data)
}

// ValidateResponse validates response and handles errors
func (m *MpcBaseAPI) ValidateResponse(response map[string]interface{}) (interface{}, error) {
	// Check for error code
	var code interface{}
	var ok bool

	if code, ok = response["code"]; !ok {
		return response, nil
	}

	// Convert code to int for comparison
	var codeInt int
	switch v := code.(type) {
	case float64:
		codeInt = int(v)
	case int:
		codeInt = v
	case string:
		if v == "0" {
			codeInt = 0
		} else {
			return nil, fmt.Errorf("API Error [%s]: %v", v, response["msg"])
		}
	default:
		codeInt = -1
	}

	if codeInt != utils.ResponseCodeSuccess {
		msg := "Unknown error"
		if msgField, ok := response["msg"]; ok {
			msg = fmt.Sprintf("%v", msgField)
		}
		return nil, fmt.Errorf("API Error [%d]: %s", codeInt, msg)
	}

	return response, nil
}

// SafeUnmarshalResponse safely unmarshals response, handling cases where data field is bool
func SafeUnmarshalResponse(response map[string]interface{}, result interface{}) error {
	// Check if data field exists and is not a valid object (e.g., false when error)
	if data, ok := response["data"]; ok {
		switch data.(type) {
		case bool:
			// When data is false, set it to nil for proper unmarshaling
			response["data"] = nil
		}
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

// ResponseError represents an API error response
type ResponseError struct {
	Code    int
	Message string
}

// Error implements the error interface
func (e *ResponseError) Error() string {
	return fmt.Sprintf("API Error [%d]: %s", e.Code, e.Message)
}

// NewResponseError creates a new ResponseError
func NewResponseError(code int, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

// IsResponseError checks if an error is a ResponseError
func IsResponseError(err error) bool {
	var respErr *ResponseError
	return errors.As(err, &respErr)
}
