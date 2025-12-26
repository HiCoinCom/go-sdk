// Package mpc provides MPC (Multi-Party Computation) client implementation.
package mpc

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"chainup.com/go-sdk/utils"
)

// Config holds the configuration for MPC client.
type Config struct {
	// Domain is the API domain URL (required).
	Domain string

	// AppID is the application ID (required).
	AppID string

	// RsaPrivateKey is the PEM-encoded RSA private key for request signing.
	RsaPrivateKey string

	// WaasPublicKey is the PEM-encoded RSA public key for response verification.
	WaasPublicKey string

	// ApiKey is the API key for authentication.
	ApiKey string

	// SignPrivateKey is the PEM-encoded RSA private key for transaction signing.
	SignPrivateKey string

	// Debug enables debug mode for logging.
	Debug bool

	// Timeout is the HTTP request timeout in seconds.
	Timeout int

	// CryptoProvider is a custom crypto provider (optional).
	CryptoProvider utils.CryptoProvider

	// Cached parsed sign private key
	signPrivateKey *rsa.PrivateKey
}

// MpcConfig is an alias for Config for backward compatibility.
type MpcConfig = Config

// Validate checks if the configuration is valid.
func (c *Config) Validate() error {
	if c.Domain == "" {
		c.Domain = utils.DefaultDomain
	}
	if c.AppID == "" {
		return errors.New("app_id is required")
	}

	if c.Timeout <= 0 {
		c.Timeout = utils.DefaultTimeout
	}

	if c.CryptoProvider == nil && c.RsaPrivateKey == "" {
		return errors.New("rsa_private_key is required (or provide crypto_provider)")
	}

	// Parse sign private key if provided
	if c.SignPrivateKey != "" {
		key, err := utils.ParsePrivateKey(c.SignPrivateKey)
		if err != nil {
			return fmt.Errorf("failed to parse sign private key: %w", err)
		}
		c.signPrivateKey = key
	}

	if c.CryptoProvider == nil && c.RsaPrivateKey != "" {
		// Use NewRSACryptoProviderWithSignKey to support separate signing key
		provider, err := utils.NewRSACryptoProviderWithSignKey(c.RsaPrivateKey, c.WaasPublicKey, c.SignPrivateKey, "UTF-8")
		if err != nil {
			return fmt.Errorf("failed to create crypto provider: %w", err)
		}
		c.CryptoProvider = provider
	} else if c.CryptoProvider != nil && c.signPrivateKey != nil {
		// If CryptoProvider is already set but we have a signPrivateKey, try to set it
		if rsaProvider, ok := c.CryptoProvider.(*utils.RSACryptoProvider); ok {
			rsaProvider.SetSignPrivateKey(c.signPrivateKey)
		}
	}

	return nil
}

// GetDomain returns the API domain.
func (c *Config) GetDomain() string {
	if len(c.Domain) == 0 {
		return utils.DefaultDomain
	}

	return c.Domain
}

// GetAppID returns the application ID.
func (c *Config) GetAppID() string {
	return c.AppID
}

// GetApiKey returns the API key.
func (c *Config) GetApiKey() string {
	return c.ApiKey
}

// IsDebug returns the debug flag.
func (c *Config) IsDebug() bool {
	return c.Debug
}

// GetCryptoProvider returns the crypto provider.
func (c *Config) GetCryptoProvider() utils.CryptoProvider {
	return c.CryptoProvider
}

// GetSignPrivateKey returns the parsed RSA private key for transaction signing.
func (c *Config) GetSignPrivateKey() *rsa.PrivateKey {
	return c.signPrivateKey
}

// ConfigBuilder helps build Config with a fluent interface.
type ConfigBuilder struct {
	config *Config
}

// MpcConfigBuilder is an alias for ConfigBuilder for backward compatibility.
type MpcConfigBuilder = ConfigBuilder

// NewMpcConfigBuilder creates a new ConfigBuilder.
func NewMpcConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: &Config{},
	}
}

// SetDomain sets the API domain URL.
func (b *ConfigBuilder) SetDomain(domain string) *ConfigBuilder {
	b.config.Domain = domain
	return b
}

// SetAppID sets the application ID.
func (b *ConfigBuilder) SetAppID(appID string) *ConfigBuilder {
	b.config.AppID = appID
	return b
}

// SetRsaPrivateKey sets the RSA private key for request signing.
func (b *ConfigBuilder) SetRsaPrivateKey(rsaPrivateKey string) *ConfigBuilder {
	b.config.RsaPrivateKey = rsaPrivateKey
	return b
}

// SetWaasPublicKey sets the WaaS server public key for response verification.
func (b *ConfigBuilder) SetWaasPublicKey(waasPublicKey string) *ConfigBuilder {
	b.config.WaasPublicKey = waasPublicKey
	return b
}

// SetApiKey sets the API key for authentication.
func (b *ConfigBuilder) SetApiKey(apiKey string) *ConfigBuilder {
	b.config.ApiKey = apiKey
	return b
}

// SetSignPrivateKey sets the signing private key for transaction signatures.
func (b *ConfigBuilder) SetSignPrivateKey(signPrivateKey string) *ConfigBuilder {
	b.config.SignPrivateKey = signPrivateKey
	return b
}

// SetDebug enables or disables debug mode.
func (b *ConfigBuilder) SetDebug(debug bool) *ConfigBuilder {
	b.config.Debug = debug
	return b
}

// SetTimeout sets the HTTP request timeout in seconds.
func (b *ConfigBuilder) SetTimeout(timeout int) *ConfigBuilder {
	b.config.Timeout = timeout
	return b
}

// SetCryptoProvider sets a custom crypto provider.
func (b *ConfigBuilder) SetCryptoProvider(provider utils.CryptoProvider) *ConfigBuilder {
	b.config.CryptoProvider = provider
	return b
}

// Build creates and validates the Config.
func (b *ConfigBuilder) Build() (*Config, error) {
	if err := b.config.Validate(); err != nil {
		return nil, err
	}
	return b.config, nil
}
