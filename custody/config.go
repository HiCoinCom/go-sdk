// Package custody provides WaaS (Wallet-as-a-Service) client implementation.
package custody

import (
	"errors"
	"fmt"

	"chainup.com/go-sdk/utils"
)

// Config holds the configuration for WaaS client.
type Config struct {
	// Host is the API host URL (required).
	Host string

	// AppID is the application ID (required).
	AppID string

	// PrivateKey is the PEM-encoded RSA private key for request signing (required).
	PrivateKey string

	// PublicKey is the PEM-encoded RSA public key for response verification (required).
	PublicKey string

	// Version is the API version (default: v1).
	Version string

	// Charset is the character encoding (default: UTF-8).
	Charset string

	// Debug enables debug mode for logging.
	Debug bool

	// Timeout is the HTTP request timeout in seconds.
	Timeout int

	// CryptoProvider is a custom crypto provider (optional).
	CryptoProvider utils.CryptoProvider
}

// WaasConfig is an alias for Config for backward compatibility.
type WaasConfig = Config

// Validate checks if the configuration is valid.
func (c *Config) Validate() error {
	if c.Host == "" {
		c.Host = utils.DefaultDomain
	}
	if c.AppID == "" {
		return errors.New("app_id is required")
	}
	if c.PrivateKey == "" {
		return errors.New("private_key is required")
	}
	if c.PublicKey == "" {
		return errors.New("public_key is required")
	}

	if c.Version == "" {
		c.Version = utils.DefaultVersion
	}
	if c.Charset == "" {
		c.Charset = utils.DefaultCharset
	}
	if c.Timeout <= 0 {
		c.Timeout = utils.DefaultTimeout
	}

	if c.CryptoProvider == nil {
		provider, err := utils.NewRSACryptoProvider(c.PrivateKey, c.PublicKey, c.Charset)
		if err != nil {
			return fmt.Errorf("failed to create crypto provider: %w", err)
		}
		c.CryptoProvider = provider
	}

	return nil
}

// GetHost returns the API host.
func (c *Config) GetHost() string {
	if len(c.Host) == 0 {
		return utils.DefaultDomain
	}

	return c.Host
}

// GetAppID returns the application ID.
func (c *Config) GetAppID() string {
	return c.AppID
}

// GetCharset returns the charset.
func (c *Config) GetCharset() string {
	return c.Charset
}

// GetDebug returns the debug flag.
func (c *Config) GetDebug() bool {
	return c.Debug
}

// GetTimeout returns the timeout.
func (c *Config) GetTimeout() int {
	return c.Timeout
}

// GetCryptoProvider returns the crypto provider.
func (c *Config) GetCryptoProvider() utils.CryptoProvider {
	return c.CryptoProvider
}

// ConfigBuilder helps build Config with a fluent interface.
type ConfigBuilder struct {
	config *Config
}

// WaasConfigBuilder is an alias for ConfigBuilder for backward compatibility.
type WaasConfigBuilder = ConfigBuilder

// NewWaasConfigBuilder creates a new ConfigBuilder.
func NewWaasConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{
		config: &Config{},
	}
}

// SetHost sets the API host URL.
func (b *ConfigBuilder) SetHost(host string) *ConfigBuilder {
	b.config.Host = host
	return b
}

// SetAppID sets the application ID.
func (b *ConfigBuilder) SetAppID(appID string) *ConfigBuilder {
	b.config.AppID = appID
	return b
}

// SetPrivateKey sets the RSA private key.
func (b *ConfigBuilder) SetPrivateKey(privateKey string) *ConfigBuilder {
	b.config.PrivateKey = privateKey
	return b
}

// SetPublicKey sets the ChainUp public key.
func (b *ConfigBuilder) SetPublicKey(publicKey string) *ConfigBuilder {
	b.config.PublicKey = publicKey
	return b
}

// SetVersion sets the API version.
func (b *ConfigBuilder) SetVersion(version string) *ConfigBuilder {
	b.config.Version = version
	return b
}

// SetCharset sets the charset encoding.
func (b *ConfigBuilder) SetCharset(charset string) *ConfigBuilder {
	b.config.Charset = charset
	return b
}

// SetDebug enables or disables debug mode.
func (b *ConfigBuilder) SetDebug(debug bool) *ConfigBuilder {
	b.config.Debug = debug
	return b
}

// SetTimeout sets the HTTP request timeout.
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
