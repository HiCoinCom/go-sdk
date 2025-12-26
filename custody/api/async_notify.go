// Package api provides API implementations for WaaS operations
package api

import (
	"encoding/json"
	"errors"

	"chainup.com/go-sdk/custody/types"
)

// AsyncNotifyAPI provides async notification operations
type AsyncNotifyAPI struct {
	*BaseAPI
}

// NewAsyncNotifyAPI creates a new AsyncNotifyAPI instance
func NewAsyncNotifyAPI(config ConfigProvider) *AsyncNotifyAPI {
	return &AsyncNotifyAPI{
		BaseAPI: NewBaseAPI(config),
	}
}

// NotifyRequest decrypts and processes an async notification (alias for VerifyRequest)
// Parameters:
//   - cipher: Encrypted notification data
//
// Returns: Decrypted notification data as AsyncNotifyArgs
func (a *AsyncNotifyAPI) NotifyRequest(cipher string) (*types.AsyncNotifyArgs, error) {
	if cipher == "" {
		return nil, errors.New("VerifyRequest: cipher cannot be empty")
	}

	if a.cryptoProvider == nil {
		return nil, NewResponseError(-1, "crypto provider not set")
	}

	// Decrypt the cipher text with public key
	decrypted, err := a.cryptoProvider.DecryptWithPublicKey(cipher)
	if err != nil {
		return nil, err
	}

	if a.debug {
		println("[AsyncNotify VerifyRequest Decrypted]:", decrypted)
	}

	if decrypted == "" {
		return nil, errors.New("VerifyRequest: decode cipher returned empty")
	}

	// Parse the decrypted JSON into AsyncNotifyArgs
	var args types.AsyncNotifyArgs
	if err := json.Unmarshal([]byte(decrypted), &args); err != nil {
		return nil, err
	}

	return &args, nil
}

// VerifyRequest decrypts and verifies an async notification request
// Parameters:
//   - cipher: Encrypted notification data from webhook callback
//
// Returns: Decrypted and parsed WithdrawArgs
func (a *AsyncNotifyAPI) VerifyRequest(cipher string) (*WithdrawArgs, error) {
	if cipher == "" {
		return nil, errors.New("VerifyRequest: cipher cannot be empty")
	}

	if a.cryptoProvider == nil {
		return nil, NewResponseError(-1, "crypto provider not set")
	}

	// Decrypt the cipher text with public key
	decrypted, err := a.cryptoProvider.DecryptWithPublicKey(cipher)
	if err != nil {
		return nil, err
	}

	if a.debug {
		println("[AsyncNotify VerifyRequest Decrypted]:", decrypted)
	}

	if decrypted == "" {
		return nil, errors.New("VerifyRequest: decode cipher returned empty")
	}

	var args WithdrawArgs
	if err := json.Unmarshal([]byte(decrypted), &args); err != nil {
		return nil, err
	}

	return &args, nil
}

// VerifyResponse encrypts the WithdrawArgs for response
// Parameters:
//   - args: WithdrawArgs to encrypt and return
//
// Returns: Encrypted response string
func (a *AsyncNotifyAPI) VerifyResponse(args *WithdrawArgs) (string, error) {
	if args == nil {
		return "", errors.New("VerifyResponse: args cannot be nil")
	}

	if a.cryptoProvider == nil {
		return "", NewResponseError(-1, "crypto provider not set")
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(args)
	if err != nil {
		return "", err
	}

	if a.debug {
		println("[AsyncNotify VerifyResponse JSON]:", string(jsonData))
	}

	// Encrypt with private key
	encrypted, err := a.cryptoProvider.EncryptWithPrivateKey(string(jsonData))
	if err != nil {
		return "", err
	}

	return encrypted, nil
}
