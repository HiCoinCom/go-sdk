// Package api provides MPC API implementations
package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"chainup.com/go-sdk/mpc/types"
)

// NotifyAPI provides notification operations
type NotifyAPI struct {
	*MpcBaseAPI
}

// NewNotifyAPI creates a new NotifyAPI instance
func NewNotifyAPI(config MpcConfigProvider) *NotifyAPI {
	return &NotifyAPI{
		MpcBaseAPI: NewMpcBaseAPI(config),
	}
}

// NotifyRequest decrypts deposit and withdrawal notification parameters.
// cipher: Encrypted notification data
func (n *NotifyAPI) NotifyRequest(cipher string) (*types.NotifyData, error) {
	if cipher == "" {
		return nil, errors.New("parameter \"cipher\" is required")
	}

	// Decrypt the cipher using crypto provider
	cryptoProvider := n.config.GetCryptoProvider()
	if cryptoProvider == nil {
		return nil, errors.New("crypto provider is required for notification decryption")
	}

	decrypted, err := cryptoProvider.DecryptWithPublicKey(cipher)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt notification: %w", err)
	}

	var notifyData types.NotifyData
	if err := json.Unmarshal([]byte(decrypted), &notifyData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal notification: %w", err)
	}

	return &notifyData, nil
}
