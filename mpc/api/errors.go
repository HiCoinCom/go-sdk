package api

import (
	"errors"
	"fmt"
)

// ResponseError represents an API error response.
type ResponseError struct {
	Code    int
	Message string
}

// Error implements the error interface.
func (e *ResponseError) Error() string {
	return fmt.Sprintf("API Error [%d]: %s", e.Code, e.Message)
}

// NewResponseError creates a new ResponseError.
func NewResponseError(code int, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

// IsResponseError checks if an error is a ResponseError.
func IsResponseError(err error) bool {
	var respErr *ResponseError
	return errors.As(err, &respErr)
}

// AsResponseError attempts to convert an error to a ResponseError.
// Returns nil if the error is not a ResponseError.
func AsResponseError(err error) *ResponseError {
	var respErr *ResponseError
	if errors.As(err, &respErr) {
		return respErr
	}
	return nil
}
