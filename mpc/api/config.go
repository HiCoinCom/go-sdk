// Package api provides MPC API implementations.
package api

import (
	"crypto/rsa"

	"chainup.com/go-sdk/utils"
)

// MpcConfigProvider defines the interface for accessing MPC configuration.
// Implementations of this interface provide the necessary credentials and settings
// for authenticating and communicating with the MPC API.
type MpcConfigProvider interface {
	// GetDomain returns the API base URL.
	GetDomain() string

	// GetAppID returns the application ID for authentication.
	GetAppID() string

	// GetApiKey returns the API key for authentication.
	GetApiKey() string

	// IsDebug returns whether debug mode is enabled.
	IsDebug() bool

	// GetCryptoProvider returns the crypto provider for encryption/decryption.
	GetCryptoProvider() utils.CryptoProvider

	// GetSignPrivateKey returns the private key for transaction signing.
	GetSignPrivateKey() *rsa.PrivateKey
}
