// Package utils provides utility functions and constants for ChainUp Custody SDK.
package utils

// API configuration constants.
const (
	// DefaultCharset is the default character encoding for API requests.
	DefaultCharset = "UTF-8"

	// DefaultVersion is the default API version.
	DefaultVersion = "v1"

	// DefaultTimeout is the default HTTP request timeout in seconds.
	DefaultTimeout = 30

	DefaultDomain = "https://openapi.chainup.com"
)

// HTTP method constants.
const (
	// HTTPMethodGet represents the HTTP GET method.
	HTTPMethodGet = "GET"

	// HTTPMethodPost represents the HTTP POST method.
	HTTPMethodPost = "POST"
)

// HTTP content type constants.
const (
	// ContentTypeFormURLEncoded is the form-urlencoded content type.
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"

	// ContentTypeJSON is the JSON content type.
	ContentTypeJSON = "application/json"
)

// Response code constants.
const (
	// ResponseCodeSuccess indicates a successful API response.
	ResponseCodeSuccess = 0

	// ResponseCodeSuccessStr is the string version of success code.
	ResponseCodeSuccessStr = "0"
)
