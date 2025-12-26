// Package api provides MPC API implementations
package api

import (
	"errors"
	"fmt"
	"strings"

	"chainup.com/go-sdk/mpc/types"
	"chainup.com/go-sdk/utils/mpcsign"
)

// WithdrawAPI provides withdrawal operations
type WithdrawAPI struct {
	*MpcBaseAPI
}

// NewWithdrawAPI creates a new WithdrawAPI instance
func NewWithdrawAPI(config MpcConfigProvider) *WithdrawAPI {
	return &WithdrawAPI{
		MpcBaseAPI: NewMpcBaseAPI(config),
	}
}

// Withdraw initiates a withdrawal
// req: Withdrawal request parameters
// needTransactionSign: Whether to sign the transaction (requires signPrivateKey in config)
func (w *WithdrawAPI) Withdraw(req *types.WithdrawRequest, needTransactionSign bool) (*types.WithdrawResponse, error) {
	if req == nil {
		return nil, errors.New("withdraw request is required")
	}

	// Build params map
	params := make(map[string]interface{})
	params["request_id"] = req.RequestID
	params["sub_wallet_id"] = req.WalletID
	params["symbol"] = req.Symbol
	params["amount"] = req.Amount.String()
	params["address_to"] = req.AddressTo

	if req.From != "" {
		params["from"] = req.From
	}
	if req.Memo != "" {
		params["memo"] = req.Memo
	}
	if req.Remark != "" {
		params["remark"] = req.Remark
	}
	if req.Outputs != "" {
		params["outputs"] = req.Outputs
	}

	// Generate signature if needed
	if needTransactionSign {
		signProvider := w.config.GetCryptoProvider()
		if signProvider == nil {
			return nil, fmt.Errorf("crypto provider is required when needTransactionSign is true")
		}

		// Build sign params
		signParams := map[string]string{
			"request_id":    req.RequestID,
			"sub_wallet_id": fmt.Sprintf("%d", req.WalletID),
			"symbol":        req.Symbol,
			"address_to":    req.AddressTo,
			"amount":        req.Amount.String(),
		}

		if req.Memo != "" {
			signParams["memo"] = req.Memo
		}
		if req.Outputs != "" {
			signParams["outputs"] = req.Outputs
		}

		signature, err := mpcsign.GenerateWithdrawSign(signParams, signProvider)
		if err != nil {
			return nil, fmt.Errorf("failed to generate signature: %w", err)
		}

		params["sign"] = signature
	}

	response, err := w.Post("/api/mpc/billing/withdraw", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var withdrawResp types.WithdrawResponse
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &withdrawResp); err != nil {
		return nil, err
	}

	return &withdrawResp, nil
}

// GetWithdrawRecords gets withdrawal records by request IDs
// requestIDs: List of request IDs
func (w *WithdrawAPI) GetWithdrawRecords(requestIDs []string) (*types.WithdrawRecordResult, error) {
	if len(requestIDs) == 0 {
		return nil, errors.New("parameter \"request_ids\" is required and must be a non-empty array")
	}

	params := map[string]interface{}{
		"ids": strings.Join(requestIDs, ","),
	}

	response, err := w.Get("/api/mpc/billing/withdraw_list", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var withdrawResult types.WithdrawRecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &withdrawResult); err != nil {
		return nil, err
	}

	return &withdrawResult, nil
}

// SyncWithdrawRecords syncs withdrawal records by max ID
// maxID: Withdrawal record initial ID, default is 0
func (w *WithdrawAPI) SyncWithdrawRecords(maxID int64) (*types.WithdrawRecordResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := w.Get("/api/mpc/billing/sync_withdraw_list", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var withdrawResult types.WithdrawRecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &withdrawResult); err != nil {
		return nil, err
	}

	return &withdrawResult, nil
}
