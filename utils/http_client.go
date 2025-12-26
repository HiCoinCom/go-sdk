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

// HTTPClient provides HTTP request functionality for API communication.
type HTTPClient struct {
	client  *http.Client
	baseURL string
	debug   bool
}

// NewHTTPClient creates a new HTTP client.
func NewHTTPClient(baseURL string, timeout int, debug bool) *HTTPClient {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	return &HTTPClient{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		baseURL: baseURL,
		debug:   debug,
	}
}

// Request executes an HTTP request.
func (h *HTTPClient) Request(method, path string, data map[string]interface{}, headers map[string]string) (string, error) {
	fullURL := h.baseURL + path

	var req *http.Request
	var err error

	switch method {
	case HTTPMethodPost:
		formData := url.Values{}
		for key, value := range data {
			formData.Set(key, fmt.Sprintf("%v", value))
		}

		req, err = http.NewRequest(method, fullURL, strings.NewReader(formData.Encode()))
		if err != nil {
			return "", fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", ContentTypeFormURLEncoded)

	case HTTPMethodGet:
		req, err = http.NewRequest(method, fullURL, nil)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %w", err)
		}

		q := req.URL.Query()
		for key, value := range data {
			q.Add(key, fmt.Sprintf("%v", value))
		}
		req.URL.RawQuery = q.Encode()

	default:
		return "", fmt.Errorf("unsupported HTTP method: %s", method)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	if h.debug {
		fmt.Printf("[HTTP Request] %s %s\n", method, fullURL)
		fmt.Printf("[HTTP Data] %+v\n", data)
	}

	resp, err := h.client.Do(req)
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

	if h.debug {
		fmt.Printf("[HTTP Response] %s\n", string(body))
	}

	return string(body), nil
}

// Post executes a POST request.
func (h *HTTPClient) Post(path string, data map[string]interface{}) (string, error) {
	return h.Request(HTTPMethodPost, path, data, nil)
}

// Get executes a GET request.
func (h *HTTPClient) Get(path string, data map[string]interface{}) (string, error) {
	return h.Request(HTTPMethodGet, path, data, nil)
}

// MpcHTTPClient provides HTTP request functionality for MPC API communication.
type MpcHTTPClient struct {
	client  *http.Client
	baseURL string
	appID   string
	apiKey  string
	debug   bool
}

// NewMpcHTTPClient creates a new MPC HTTP client.
func NewMpcHTTPClient(baseURL, appID, apiKey string, timeout int, debug bool) *MpcHTTPClient {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	return &MpcHTTPClient{
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		baseURL: baseURL,
		appID:   appID,
		apiKey:  apiKey,
		debug:   debug,
	}
}

// Request executes an HTTP request for MPC API.
func (m *MpcHTTPClient) Request(method, path string, data map[string]interface{}) (string, error) {
	fullURL := m.baseURL + path

	if data == nil {
		data = make(map[string]interface{})
	}
	data["app_id"] = m.appID

	var req *http.Request
	var err error

	switch method {
	case HTTPMethodPost:
		formData := url.Values{}
		for key, value := range data {
			formData.Set(key, fmt.Sprintf("%v", value))
		}

		req, err = http.NewRequest(method, fullURL, strings.NewReader(formData.Encode()))
		if err != nil {
			return "", fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Content-Type", ContentTypeFormURLEncoded)

	case HTTPMethodGet:
		req, err = http.NewRequest(method, fullURL, nil)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %w", err)
		}

		q := req.URL.Query()
		for key, value := range data {
			q.Add(key, fmt.Sprintf("%v", value))
		}
		req.URL.RawQuery = q.Encode()

	default:
		return "", fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if m.apiKey != "" {
		req.Header.Set("API-KEY", m.apiKey)
	}

	if m.debug {
		fmt.Printf("[MPC HTTP Request] %s %s\n", method, fullURL)
		fmt.Printf("[MPC HTTP Data] %+v\n", data)
	}

	resp, err := m.client.Do(req)
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

	if m.debug {
		fmt.Printf("[MPC HTTP Response] %s\n", string(body))
	}

	return string(body), nil
}

// Post executes a POST request.
func (m *MpcHTTPClient) Post(path string, data map[string]interface{}) (string, error) {
	return m.Request(HTTPMethodPost, path, data)
}

// Get executes a GET request.
func (m *MpcHTTPClient) Get(path string, data map[string]interface{}) (string, error) {
	return m.Request(HTTPMethodGet, path, data)
}
