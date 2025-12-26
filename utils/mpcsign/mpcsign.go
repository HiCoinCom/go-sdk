// Package mpcsign provides signature utilities for MPC API requests
package mpcsign

import (
	"crypto/md5"
	"crypto/rsa"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"chainup.com/go-sdk/utils"
)

// SignProvider defines the interface for signing operations
type SignProvider interface {
	// SignWithPrivateKey signs data using RSA-SHA256
	SignWithPrivateKey(data string) (string, error)
}

// GenerateWithdrawSign generates signature for withdraw request using CryptoProvider
func GenerateWithdrawSign(params map[string]string, provider SignProvider) (string, error) {
	// Sort params and format
	sortedStr := ParamsSort(params)
	hash := md5.Sum([]byte(strings.ToLower(sortedStr)))
	md5Sign := fmt.Sprintf("%x", hash)
	// Sign the sorted string using provider
	return provider.SignWithPrivateKey(md5Sign)
}

// GenerateWithdrawSignWithKey generates signature for withdraw request using RSA private key (legacy)
func GenerateWithdrawSignWithKey(params map[string]string, privateKey *rsa.PrivateKey) (string, error) {
	// Create a temporary provider with the private key
	tempProvider := &rsaKeyProvider{privateKey: privateKey}
	return GenerateWithdrawSign(params, tempProvider)
}

// GenerateWeb3Sign generates signature for Web3 transaction request using CryptoProvider
func GenerateWeb3Sign(params map[string]string, provider SignProvider) (string, error) {
	// Sort params and format
	sortedStr := ParamsSort(params)
	hash := md5.Sum([]byte(strings.ToLower(sortedStr)))
	md5Sign := fmt.Sprintf("%x", hash)
	// Sign the sorted string using provider
	return provider.SignWithPrivateKey(md5Sign)
}

// GenerateWeb3SignWithKey generates signature for Web3 transaction request using RSA private key (legacy)
func GenerateWeb3SignWithKey(params map[string]string, privateKey *rsa.PrivateKey) (string, error) {
	tempProvider := &rsaKeyProvider{privateKey: privateKey}
	return GenerateWeb3Sign(params, tempProvider)
}

// rsaKeyProvider is a temporary provider that wraps a raw RSA private key
type rsaKeyProvider struct {
	privateKey *rsa.PrivateKey
}

func (p *rsaKeyProvider) SignWithPrivateKey(data string) (string, error) {
	provider, err := utils.NewRSACryptoProviderWithKeys(p.privateKey, nil, "")
	if err != nil {
		return "", err
	}
	return provider.SignWithPrivateKey(data)
}

// ParamsSort sorts parameters and formats them as required for signing
func ParamsSort(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}

	// Get keys and sort
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build sorted string
	var parts []string
	for _, k := range keys {
		v := params[k]
		if v == "" {
			continue
		}
		// Remove trailing zeros for numeric values
		//v = removeTrailingZeros(v)
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}

	str := strings.Join(parts, "&")
	return str
}

// removeTrailingZeros removes trailing zeros from numeric strings
func removeTrailingZeros(s string) string {
	// Check if it's a numeric value with decimal point
	if matched, _ := regexp.MatchString(`^\d+\.\d+$`, s); !matched {
		return s
	}

	// Remove trailing zeros
	s = strings.TrimRight(s, "0")
	// Remove trailing decimal point
	s = strings.TrimRight(s, ".")

	return s
}
