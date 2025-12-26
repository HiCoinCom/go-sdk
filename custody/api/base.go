// Package api provides API implementations for WaaS operations
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"chainup.com/go-sdk/utils"
)

// ConfigProvider defines the interface for accessing configuration
type ConfigProvider interface {
	GetHost() string
	GetAppID() string
	GetCharset() string
	GetDebug() bool
	GetTimeout() int
	GetCryptoProvider() utils.CryptoProvider
}

// BaseAPI provides common functionality for all WaaS API implementations
type BaseAPI struct {
	host           string
	appID          string
	charset        string
	debug          bool
	httpClient     *utils.HTTPClient
	cryptoProvider utils.CryptoProvider
}

// WaaS API version prefix
const waasAPIPrefix = "/api/v2"

// NewBaseAPI creates a new BaseAPI instance
func NewBaseAPI(config ConfigProvider) *BaseAPI {
	baseURL := config.GetHost() + waasAPIPrefix
	return &BaseAPI{
		host:           baseURL,
		appID:          config.GetAppID(),
		charset:        config.GetCharset(),
		debug:          config.GetDebug(),
		httpClient:     utils.NewHTTPClient(baseURL, config.GetTimeout(), config.GetDebug()),
		cryptoProvider: config.GetCryptoProvider(),
	}
}

// buildRequestArgs builds the request args JSON with common parameters
func (b *BaseAPI) buildRequestArgs(data map[string]interface{}) (string, error) {
	if data == nil {
		data = make(map[string]interface{})
	}

	// Add common parameters
	data["time"] = time.Now().UnixMilli()
	data["charset"] = b.charset

	// Convert to JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data: %w", err)
	}

	return string(jsonBytes), nil
}

// executeRequest executes an API request with signing and encryption
func (b *BaseAPI) executeRequest(method, path string, data map[string]interface{}) (map[string]interface{}, error) {
	// Step 1: Build request args JSON
	rawJSON, err := b.buildRequestArgs(data)
	if err != nil {
		return nil, err
	}

	if b.debug {
		fmt.Printf("[WaaS Request Args]: %s\n", rawJSON)
	}

	// Step 2: Encrypt with private key
	encryptedData := ""
	if b.cryptoProvider != nil {
		encrypted, err := b.cryptoProvider.EncryptWithPrivateKey(rawJSON)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt request data: %w", err)
		}
		encryptedData = encrypted

		if b.debug {
			if len(encryptedData) > 100 {
				fmt.Printf("[WaaS Encrypted Data]: %s...\n", encryptedData[:100])
			} else {
				fmt.Printf("[WaaS Encrypted Data]: %s\n", encryptedData)
			}
		}
	}

	// Step 3: Send request with only app_id and data
	requestData := map[string]interface{}{
		"app_id": b.appID,
		"data":   encryptedData,
	}

	var response string
	if method == utils.HTTPMethodPost {
		response, err = b.httpClient.Post(path, requestData)
	} else {
		response, err = b.httpClient.Get(path, requestData)
	}

	if err != nil {
		return nil, err
	}

	if b.debug {
		fmt.Printf("[WaaS Response]: %s\n", response)
	}

	// Step 4: Parse response and decrypt data if needed
	var parsedResponse map[string]interface{}
	if err := json.Unmarshal([]byte(response), &parsedResponse); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	// Check if response has encrypted data field and decrypt
	if dataField, ok := parsedResponse["data"].(string); ok && dataField != "" {
		if b.cryptoProvider != nil {
			decrypted, err := b.cryptoProvider.DecryptWithPublicKey(dataField)
			if err != nil {
				if b.debug {
					fmt.Printf("[WaaS Decrypt Error]: %v\n", err)
				}
				// If decryption fails, might be an error response, return as-is
				return parsedResponse, nil
			}

			if b.debug {
				fmt.Printf("[WaaS Decrypted]: %s\n", decrypted)
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
func (b *BaseAPI) Post(path string, data map[string]interface{}) (map[string]interface{}, error) {
	return b.executeRequest(utils.HTTPMethodPost, path, data)
}

// Get executes a GET request
func (b *BaseAPI) Get(path string, data map[string]interface{}) (map[string]interface{}, error) {
	return b.executeRequest(utils.HTTPMethodGet, path, data)
}

// ValidateResponse validates response and handles errors
func (b *BaseAPI) ValidateResponse(response map[string]interface{}) (interface{}, error) {
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

	// Return data field if exists, otherwise return whole response
	if dataField, ok := response["data"]; ok {
		return dataField, nil
	}

	return response, nil
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
