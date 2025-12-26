package api

import (
	"encoding/json"
	"fmt"

	"chainup.com/go-sdk/utils"
)

// ValidateResponse validates an API response and returns the result or an error.
// It checks the response code and returns an error if the API call was unsuccessful.
func ValidateResponse(response map[string]interface{}) (interface{}, error) {
	code, ok := response["code"]
	if !ok {
		return response, nil
	}

	codeInt, codeStr := parseResponseCode(code)
	if codeInt != utils.ResponseCodeSuccess {
		msg := extractErrorMessage(response)
		if codeStr != "" {
			return nil, fmt.Errorf("API Error [%s]: %s", codeStr, msg)
		}
		return nil, NewResponseError(codeInt, msg)
	}

	return response, nil
}

// parseResponseCode converts a response code to an integer and original string.
func parseResponseCode(code interface{}) (int, string) {
	switch v := code.(type) {
	case float64:
		return int(v), ""
	case int:
		return v, ""
	case string:
		if v == "0" {
			return 0, ""
		}
		return -1, v
	default:
		return -1, ""
	}
}

// extractErrorMessage extracts the error message from a response.
func extractErrorMessage(response map[string]interface{}) string {
	if msg, ok := response["msg"]; ok {
		return fmt.Sprintf("%v", msg)
	}
	return "Unknown error"
}

// SafeUnmarshalResponse safely unmarshals a response map into a result struct.
// It handles special cases where the data field might be a boolean instead of an object.
func SafeUnmarshalResponse(response map[string]interface{}, result interface{}) error {
	// Handle case where data field is a boolean (e.g., false on error)
	if data, ok := response["data"]; ok {
		if _, isBool := data.(bool); isBool {
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
