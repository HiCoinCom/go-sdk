// Package utils provides utility functions and constants for ChainUp Custody SDK.
package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// RequestOption defines a function type for configuring HTTP requests.
type RequestOption func(*http.Request)

// WithHeader adds a header to the request.
func WithHeader(key, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

// WithHeaders adds multiple headers to the request.
func WithHeaders(headers map[string]string) RequestOption {
	return func(req *http.Request) {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
}

// BaseHTTPClient provides common HTTP request functionality.
type BaseHTTPClient struct {
	client  *http.Client
	baseURL string
	debug   bool
}

// NewBaseHTTPClient creates a new base HTTP client.
func NewBaseHTTPClient(baseURL string, timeout int, debug bool) *BaseHTTPClient {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	return &BaseHTTPClient{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		baseURL: baseURL,
		debug:   debug,
	}
}

// buildRequest creates an HTTP request based on method and data.
func (b *BaseHTTPClient) buildRequest(method, fullURL string, data map[string]interface{}) (*http.Request, error) {
	switch method {
	case HTTPMethodPost:
		formData := url.Values{}
		for key, value := range data {
			formData.Set(key, fmt.Sprintf("%v", value))
		}
		req, err := http.NewRequest(method, fullURL, strings.NewReader(formData.Encode()))
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", ContentTypeFormURLEncoded)
		return req, nil

	case HTTPMethodGet:
		req, err := http.NewRequest(method, fullURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		if len(data) > 0 {
			q := req.URL.Query()
			for key, value := range data {
				q.Add(key, fmt.Sprintf("%v", value))
			}
			req.URL.RawQuery = q.Encode()
		}
		return req, nil

	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}
}

// execute performs the HTTP request and returns the response body.
func (b *BaseHTTPClient) execute(req *http.Request, logPrefix string) (string, error) {
	if b.debug {
		fmt.Printf("[%s Request] %s %s\n", logPrefix, req.Method, req.URL.String())
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	if b.debug {
		fmt.Printf("[%s Response] %s\n", logPrefix, string(body))
	}

	return string(body), nil
}

// Request executes an HTTP request with optional configurations.
func (b *BaseHTTPClient) Request(method, path string, data map[string]interface{}, opts ...RequestOption) (string, error) {
	fullURL := b.baseURL + path

	req, err := b.buildRequest(method, fullURL, data)
	if err != nil {
		return "", err
	}

	// Apply request options
	for _, opt := range opts {
		opt(req)
	}

	if b.debug {
		fmt.Printf("[HTTP Data] %+v\n", data)
	}

	return b.execute(req, "HTTP")
}

// Post executes a POST request.
func (b *BaseHTTPClient) Post(path string, data map[string]interface{}, opts ...RequestOption) (string, error) {
	return b.Request(HTTPMethodPost, path, data, opts...)
}

// Get executes a GET request.
func (b *BaseHTTPClient) Get(path string, data map[string]interface{}, opts ...RequestOption) (string, error) {
	return b.Request(HTTPMethodGet, path, data, opts...)
}

// HTTPClient provides HTTP request functionality for API communication.
// It embeds BaseHTTPClient for common functionality.
type HTTPClient struct {
	*BaseHTTPClient
}

// NewHTTPClient creates a new HTTP client.
func NewHTTPClient(baseURL string, timeout int, debug bool) *HTTPClient {
	return &HTTPClient{
		BaseHTTPClient: NewBaseHTTPClient(baseURL, timeout, debug),
	}
}

// Request executes an HTTP request with headers support.
func (h *HTTPClient) Request(method, path string, data map[string]interface{}, headers map[string]string) (string, error) {
	var opts []RequestOption
	if len(headers) > 0 {
		opts = append(opts, WithHeaders(headers))
	}
	return h.BaseHTTPClient.Request(method, path, data, opts...)
}

// MpcHTTPClient provides HTTP request functionality for MPC API communication.
type MpcHTTPClient struct {
	*BaseHTTPClient
	appID  string
	apiKey string
}

// NewMpcHTTPClient creates a new MPC HTTP client.
func NewMpcHTTPClient(baseURL, appID, apiKey string, timeout int, debug bool) *MpcHTTPClient {
	return &MpcHTTPClient{
		BaseHTTPClient: NewBaseHTTPClient(baseURL, timeout, debug),
		appID:          appID,
		apiKey:         apiKey,
	}
}

// Request executes an HTTP request for MPC API.
func (m *MpcHTTPClient) Request(method, path string, data map[string]interface{}) (string, error) {
	// Ensure data map exists and add app_id
	if data == nil {
		data = make(map[string]interface{})
	}
	data["app_id"] = m.appID

	// Build request options
	var opts []RequestOption
	if m.apiKey != "" {
		opts = append(opts, WithHeader("API-KEY", m.apiKey))
	}

	fullURL := m.baseURL + path

	req, err := m.buildRequest(method, fullURL, data)
	if err != nil {
		return "", err
	}

	// Apply options
	for _, opt := range opts {
		opt(req)
	}

	if m.debug {
		fmt.Printf("[MPC HTTP Data] %+v\n", data)
	}

	return m.execute(req, "MPC HTTP")
}

// Post executes a POST request.
func (m *MpcHTTPClient) Post(path string, data map[string]interface{}) (string, error) {
	return m.Request(HTTPMethodPost, path, data)
}

// Get executes a GET request.
func (m *MpcHTTPClient) Get(path string, data map[string]interface{}) (string, error) {
	return m.Request(HTTPMethodGet, path, data)
}
