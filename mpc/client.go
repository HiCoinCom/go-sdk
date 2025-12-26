// Package mpc provides MPC (Multi-Party Computation) client implementation.
package mpc

import (
	"chainup.com/go-sdk/mpc/api"
)

// Client is the main entry point for MPC API operations.
// It provides factory methods for creating API instances.
type Client struct {
	config *Config
}

// MpcClient is an alias for Client for backward compatibility.
type MpcClient = Client

// NewMpcClient creates a new Client with the given configuration.
func NewMpcClient(config *Config) (*Client, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		config: config,
	}, nil
}

// GetWalletAPI returns WalletAPI instance for wallet operations.
func (c *Client) GetWalletAPI() *api.WalletAPI {
	return api.NewWalletAPI(c.config)
}

// GetDepositAPI returns DepositAPI instance for deposit operations.
func (c *Client) GetDepositAPI() *api.DepositAPI {
	return api.NewDepositAPI(c.config)
}

// GetWithdrawAPI returns WithdrawAPI instance for withdrawal operations.
func (c *Client) GetWithdrawAPI() *api.WithdrawAPI {
	return api.NewWithdrawAPI(c.config)
}

// GetWeb3API returns Web3API instance for Web3 operations.
func (c *Client) GetWeb3API() *api.Web3API {
	return api.NewWeb3API(c.config)
}

// GetAutoSweepAPI returns AutoSweepAPI instance for auto-sweep operations.
func (c *Client) GetAutoSweepAPI() *api.AutoSweepAPI {
	return api.NewAutoSweepAPI(c.config)
}

// GetNotifyAPI returns NotifyAPI instance for notification operations.
func (c *Client) GetNotifyAPI() *api.NotifyAPI {
	return api.NewNotifyAPI(c.config)
}

// GetWorkSpaceAPI returns WorkSpaceAPI instance for workspace operations.
func (c *Client) GetWorkSpaceAPI() *api.WorkSpaceAPI {
	return api.NewWorkSpaceAPI(c.config)
}

// GetTronResourceAPI returns TronResourceAPI instance for TRON resource operations.
func (c *Client) GetTronResourceAPI() *api.TronResourceAPI {
	return api.NewTronResourceAPI(c.config)
}

// ClientBuilder helps build Client with a fluent interface.
type ClientBuilder struct {
	configBuilder *ConfigBuilder
}

// MpcClientBuilder is an alias for ClientBuilder for backward compatibility.
type MpcClientBuilder = ClientBuilder

// NewMpcClientBuilder creates a new ClientBuilder.
func NewMpcClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		configBuilder: NewMpcConfigBuilder(),
	}
}

// SetDomain sets the API domain URL.
func (b *ClientBuilder) SetDomain(domain string) *ClientBuilder {
	b.configBuilder.SetDomain(domain)
	return b
}

// SetAppID sets the application ID.
func (b *ClientBuilder) SetAppID(appID string) *ClientBuilder {
	b.configBuilder.SetAppID(appID)
	return b
}

// SetRsaPrivateKey sets the RSA private key.
func (b *ClientBuilder) SetRsaPrivateKey(rsaPrivateKey string) *ClientBuilder {
	b.configBuilder.SetRsaPrivateKey(rsaPrivateKey)
	return b
}

// SetWaasPublicKey sets the WaaS server public key.
func (b *ClientBuilder) SetWaasPublicKey(waasPublicKey string) *ClientBuilder {
	b.configBuilder.SetWaasPublicKey(waasPublicKey)
	return b
}

// SetApiKey sets the API key.
func (b *ClientBuilder) SetApiKey(apiKey string) *ClientBuilder {
	b.configBuilder.SetApiKey(apiKey)
	return b
}

// SetSignPrivateKey sets the signing private key.
func (b *ClientBuilder) SetSignPrivateKey(signPrivateKey string) *ClientBuilder {
	b.configBuilder.SetSignPrivateKey(signPrivateKey)
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

	return NewMpcClient(config)
}
