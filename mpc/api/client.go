package api

import (
	"encoding/json"
	"fmt"
	"time"

	"chainup.com/go-sdk/utils"
)

// MpcBaseAPI provides common functionality for all MPC API implementations.
// It handles request building, encryption, and response parsing.
type MpcBaseAPI struct {
	config         MpcConfigProvider
	httpClient     *utils.MpcHTTPClient
	cryptoProvider utils.CryptoProvider
}

// NewMpcBaseAPI creates a new MpcBaseAPI instance.
func NewMpcBaseAPI(config MpcConfigProvider) *MpcBaseAPI {
	return &MpcBaseAPI{
		config: config,
		httpClient: utils.NewMpcHTTPClient(
			config.GetDomain(),
			config.GetAppID(),
			config.GetApiKey(),
			utils.DefaultTimeout,
			config.IsDebug(),
		),
		cryptoProvider: config.GetCryptoProvider(),
	}
}

// Post executes a POST request to the specified path with the given data.
func (m *MpcBaseAPI) Post(path string, data map[string]interface{}) (map[string]interface{}, error) {
	return m.executeRequest(utils.HTTPMethodPost, path, data)
}

// Get executes a GET request to the specified path with the given data.
func (m *MpcBaseAPI) Get(path string, data map[string]interface{}) (map[string]interface{}, error) {
	return m.executeRequest(utils.HTTPMethodGet, path, data)
}

// ValidateResponse validates response and handles errors.
// This is a convenience method that delegates to the package-level ValidateResponse.
func (m *MpcBaseAPI) ValidateResponse(response map[string]interface{}) (interface{}, error) {
	return ValidateResponse(response)
}

// executeRequest executes an MPC API request with encryption and decryption.
func (m *MpcBaseAPI) executeRequest(method, path string, data map[string]interface{}) (map[string]interface{}, error) {
	// Build and encrypt request
	encryptedData, err := m.buildEncryptedRequest(data)
	if err != nil {
		return nil, err
	}

	// Send request
	response, err := m.sendRequest(method, path, encryptedData)
	if err != nil {
		return nil, err
	}

	// Parse and decrypt response
	return m.parseResponse(response)
}

// buildEncryptedRequest builds and encrypts the request data.
func (m *MpcBaseAPI) buildEncryptedRequest(data map[string]interface{}) (string, error) {
	rawJSON, err := m.buildRequestArgs(data)
	if err != nil {
		return "", err
	}

	m.debugLog("[MPC Request Args]: %s", rawJSON)

	if m.cryptoProvider == nil {
		return "", nil
	}

	encrypted, err := m.cryptoProvider.EncryptWithPrivateKey(rawJSON)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt request data: %w", err)
	}

	m.debugLogTruncated("[MPC Encrypted Data]", encrypted, 100)

	return encrypted, nil
}

// buildRequestArgs builds the request args JSON with common parameters.
func (m *MpcBaseAPI) buildRequestArgs(data map[string]interface{}) (string, error) {
	if data == nil {
		data = make(map[string]interface{})
	}

	data["time"] = time.Now().UnixMilli()
	data["charset"] = "utf-8"

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data: %w", err)
	}

	return string(jsonBytes), nil
}

// sendRequest sends the HTTP request.
func (m *MpcBaseAPI) sendRequest(method, path, encryptedData string) (string, error) {
	requestData := map[string]interface{}{
		"data":   encryptedData,
		"app_id": m.config.GetAppID(),
	}

	var response string
	var err error

	switch method {
	case utils.HTTPMethodPost:
		response, err = m.httpClient.Post(path, requestData)
	default:
		response, err = m.httpClient.Get(path, requestData)
	}

	if err != nil {
		return "", err
	}

	m.debugLog("[MPC Response]: %s", response)

	return response, nil
}

// parseResponse parses and decrypts the response.
func (m *MpcBaseAPI) parseResponse(response string) (map[string]interface{}, error) {
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(response), &parsed); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	// Check if response has encrypted data field
	dataField, ok := parsed["data"].(string)
	if !ok || dataField == "" || m.cryptoProvider == nil {
		return parsed, nil
	}

	// Decrypt the data
	decrypted, err := m.cryptoProvider.DecryptWithPublicKey(dataField)
	if err != nil {
		m.debugLog("[MPC Decrypt Error]: %v", err)
		return parsed, nil
	}

	m.debugLog("[MPC Decrypted]: %s", decrypted)

	// Parse decrypted JSON
	var decryptedResponse map[string]interface{}
	if err := json.Unmarshal([]byte(decrypted), &decryptedResponse); err != nil {
		return nil, fmt.Errorf("failed to parse decrypted data: %w", err)
	}

	return decryptedResponse, nil
}

// debugLog logs a message if debug mode is enabled.
func (m *MpcBaseAPI) debugLog(format string, args ...interface{}) {
	if m.config.IsDebug() {
		fmt.Printf(format+"\n", args...)
	}
}

// debugLogTruncated logs a truncated message if debug mode is enabled.
func (m *MpcBaseAPI) debugLogTruncated(prefix, value string, maxLen int) {
	if !m.config.IsDebug() {
		return
	}
	if len(value) > maxLen {
		fmt.Printf("%s: %s...\n", prefix, value[:maxLen])
	} else {
		fmt.Printf("%s: %s\n", prefix, value)
	}
}
