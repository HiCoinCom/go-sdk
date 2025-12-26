// Package api provides API implementations for WaaS operations
package api

import (
	"strings"

	"chainup.com/go-sdk/custody/types"
	"github.com/shopspring/decimal"
)

// TransferAPI provides transfer operations between accounts
type TransferAPI struct {
	*BaseAPI
}

// NewTransferAPI creates a new TransferAPI instance
func NewTransferAPI(config ConfigProvider) *TransferAPI {
	return &TransferAPI{
		BaseAPI: NewBaseAPI(config),
	}
}

// TransferArgs represents transfer request parameters
type TransferArgs struct {
	RequestID string          `json:"request_id"` // Unique request ID (merchant generated)
	FromUID   int64           `json:"from_uid"`   // Source user ID
	ToUID     int64           `json:"to_uid"`     // Destination user ID
	Symbol    string          `json:"symbol"`     // Cryptocurrency symbol (e.g., 'BTC', 'ETH')
	Amount    decimal.Decimal `json:"amount"`     // Transfer amount
	Remark    string          `json:"remark"`     // (Optional) Additional remark
}

// AccountTransfer performs an account transfer
// Parameters:
//   - args: Transfer request arguments
//
// Returns: Transfer result
func (t *TransferAPI) AccountTransfer(args *TransferArgs) (*types.TransferResult, error) {
	params := map[string]interface{}{
		"request_id": args.RequestID,
		"from_uid":   args.FromUID,
		"to_uid":     args.ToUID,
		"symbol":     args.Symbol,
		"amount":     args.Amount.String(),
	}

	if args.Remark != "" {
		params["remark"] = args.Remark
	}

	response, err := t.Post("/account/transfer", params)
	if err != nil {
		return nil, err
	}

	_, err = t.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.TransferResult](response)
}

// GetAccountTransferList gets account transfer list by request IDs
// Parameters:
//   - requestIDs: List of request IDs
//
// Returns: Transfer records
func (t *TransferAPI) GetAccountTransferList(requestIDs []string) (*types.TransferListResult, error) {
	params := map[string]interface{}{
		"ids": strings.Join(requestIDs, ","),
	}

	response, err := t.Post("/account/transferList", params)
	if err != nil {
		return nil, err
	}

	_, err = t.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.TransferListResult](response)
}

// SyncAccountTransferList syncs account transfer list by max ID (pagination)
// Parameters:
//   - maxID: Maximum transaction ID for pagination
//
// Returns: Synced transfer records
func (t *TransferAPI) SyncAccountTransferList(maxID int64) (*types.TransferListResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := t.Post("/account/syncTransferList", params)
	if err != nil {
		return nil, err
	}

	_, err = t.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.TransferListResult](response)
}
