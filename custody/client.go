// Package custody provides WaaS (Wallet-as-a-Service) client implementation.
package custody

import (
	"chainup.com/go-sdk/custody/api"
)

// Client is the main entry point for WaaS API operations.
// It provides factory methods for creating API instances.
type Client struct {
	config *Config
}

// WaasClient is an alias for Client for backward compatibility.
type WaasClient = Client

// NewWaasClient creates a new Client with the given configuration.
func NewWaasClient(config *Config) (*Client, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		config: config,
	}, nil
}

// GetUserAPI returns UserAPI instance for user management.
func (c *Client) GetUserAPI() *api.UserAPI {
	return api.NewUserAPI(c.config)
}

// GetAccountAPI returns AccountAPI instance for account management.
func (c *Client) GetAccountAPI() *api.AccountAPI {
	return api.NewAccountAPI(c.config)
}

// GetBillingAPI returns BillingAPI instance for billing operations.
func (c *Client) GetBillingAPI() *api.BillingAPI {
	return api.NewBillingAPI(c.config)
}

// GetCoinAPI returns CoinAPI instance for coin information.
func (c *Client) GetCoinAPI() *api.CoinAPI {
	return api.NewCoinAPI(c.config)
}

// GetTransferAPI returns TransferAPI instance for transfer operations.
func (c *Client) GetTransferAPI() *api.TransferAPI {
	return api.NewTransferAPI(c.config)
}

// GetAsyncNotifyAPI returns AsyncNotifyAPI instance for notification handling.
func (c *Client) GetAsyncNotifyAPI() *api.AsyncNotifyAPI {
	return api.NewAsyncNotifyAPI(c.config)
}

// ClientBuilder helps build Client with a fluent interface.
type ClientBuilder struct {
	configBuilder *ConfigBuilder
}

// WaasClientBuilder is an alias for ClientBuilder for backward compatibility.
type WaasClientBuilder = ClientBuilder

// NewWaasClientBuilder creates a new ClientBuilder.
func NewWaasClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		configBuilder: NewWaasConfigBuilder(),
	}
}

// SetHost sets the API host URL.
func (b *ClientBuilder) SetHost(host string) *ClientBuilder {
	b.configBuilder.SetHost(host)
	return b
}

// SetAppID sets the application ID.
func (b *ClientBuilder) SetAppID(appID string) *ClientBuilder {
	b.configBuilder.SetAppID(appID)
	return b
}

// SetPrivateKey sets the RSA private key.
func (b *ClientBuilder) SetPrivateKey(privateKey string) *ClientBuilder {
	b.configBuilder.SetPrivateKey(privateKey)
	return b
}

// SetPublicKey sets the ChainUp public key.
func (b *ClientBuilder) SetPublicKey(publicKey string) *ClientBuilder {
	b.configBuilder.SetPublicKey(publicKey)
	return b
}

// SetDebug enables or disables debug mode.
func (b *ClientBuilder) SetDebug(debug bool) *ClientBuilder {
	b.configBuilder.SetDebug(debug)
	return b
}

// SetTimeout sets the HTTP request timeout.
func (b *ClientBuilder) SetTimeout(timeout int) *ClientBuilder {
	b.configBuilder.SetTimeout(timeout)
	return b
}

// Build creates and returns a configured Client instance.
func (b *ClientBuilder) Build() (*Client, error) {
	config, err := b.configBuilder.Build()
	if err != nil {
		return nil, err
	}

	return NewWaasClient(config)
}
