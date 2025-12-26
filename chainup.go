// Package chainup provides the ChainUp Custody SDK for Go.
//
// This SDK provides a comprehensive interface to interact with ChainUp's
// custody services, including both WaaS (Wallet as a Service) and MPC
// (Multi-Party Computation) APIs.
//
// # Quick Start
//
// For WaaS API:
//
//	client, err := chainup.NewWaasClientBuilder().
//	    SetHost("https://api.custody.chainup.com").
//	    SetAppID("your-app-id").
//	    SetPrivateKey(privateKeyPEM).
//	    SetPublicKey(publicKeyPEM).
//	    Build()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	userAPI := client.GetUserAPI()
//
// For MPC API:
//
//	client, err := chainup.NewMpcClientBuilder().
//	    SetDomain("https://mpc-api.custody.chainup.com").
//	    SetAppID("your-app-id").
//	    SetRsaPrivateKey(privateKeyPEM).
//	    SetWaasPublicKey(publicKeyPEM).
//	    SetApiKey("your-api-key").
//	    Build()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	walletAPI := client.GetWalletAPI()
//
// # Architecture
//
// The SDK is organized into the following packages:
//   - chainup: Main entry point with type aliases and convenience constructors
//   - custody: WaaS API client and configuration
//   - custody/api: WaaS API implementations
//   - custody/types: WaaS data structures
//   - mpc: MPC API client and configuration
//   - mpc/api: MPC API implementations
//   - mpc/types: MPC data structures
//   - utils: Shared utilities (HTTP, crypto, constants)
package chainup

import (
	"chainup.com/go-sdk/custody"
	"chainup.com/go-sdk/mpc"
	"chainup.com/go-sdk/utils"
)

// WaaS client type aliases for convenient access from the main package.
type (
	// WaasClient is the main client for WaaS API operations.
	WaasClient = custody.WaasClient

	// WaasConfig holds the configuration for WaaS client.
	WaasConfig = custody.WaasConfig

	// WaasClientBuilder helps build WaasClient instances using a fluent interface.
	WaasClientBuilder = custody.WaasClientBuilder
)

// MPC client type aliases for convenient access from the main package.
type (
	// MpcClient is the main client for MPC API operations.
	MpcClient = mpc.MpcClient

	// MpcConfig holds the configuration for MPC client.
	MpcConfig = mpc.MpcConfig

	// MpcClientBuilder helps build MpcClient instances using a fluent interface.
	MpcClientBuilder = mpc.MpcClientBuilder
)

// Utility type aliases.
type (
	// CryptoProvider defines the interface for cryptographic operations.
	CryptoProvider = utils.CryptoProvider

	// RSACryptoProvider implements RSA encryption/decryption.
	RSACryptoProvider = utils.RSACryptoProvider
)

// NewWaasClientBuilder creates a new WaaS client builder.
func NewWaasClientBuilder() *custody.WaasClientBuilder {
	return custody.NewWaasClientBuilder()
}

// NewMpcClientBuilder creates a new MPC client builder.
func NewMpcClientBuilder() *mpc.MpcClientBuilder {
	return mpc.NewMpcClientBuilder()
}

// NewRSACryptoProvider creates a new RSA crypto provider.
// privateKeyPEM: PEM-encoded RSA private key for signing requests
// publicKeyPEM: PEM-encoded RSA public key for verifying responses
// charset: Character encoding (default: UTF-8)
func NewRSACryptoProvider(privateKeyPEM, publicKeyPEM, charset string) (*utils.RSACryptoProvider, error) {
	return utils.NewRSACryptoProvider(privateKeyPEM, publicKeyPEM, charset)
}

// SDK Constants
const (
	// Version is the current SDK version.
	Version = "1.0.0"

	// DefaultCharset is the default character encoding.
	DefaultCharset = utils.DefaultCharset

	// DefaultTimeout is the default HTTP request timeout in seconds.
	DefaultTimeout = utils.DefaultTimeout
)
