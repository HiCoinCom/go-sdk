// Package api provides API implementations for WaaS operations
package api

import (
	"fmt"
	"strings"

	"chainup.com/go-sdk/custody/types"
	"github.com/shopspring/decimal"
)

// BillingAPI provides deposit, withdrawal and miner fee operations
type BillingAPI struct {
	*BaseAPI
}

// NewBillingAPI creates a new BillingAPI instance
func NewBillingAPI(config ConfigProvider) *BillingAPI {
	return &BillingAPI{
		BaseAPI: NewBaseAPI(config),
	}
}

// WithdrawArgs represents withdrawal request parameters
type WithdrawArgs struct {
	RequestID string          `json:"request_id"` // Unique request ID (merchant generated)
	FromUID   int64           `json:"from_uid"`   // Source user ID
	ToAddress string          `json:"to_address"` // Destination address
	Amount    decimal.Decimal `json:"amount"`     // Withdrawal amount
	Symbol    string          `json:"symbol"`     // Cryptocurrency symbol (e.g., 'BTC', 'ETH')
	CheckSum  string          `json:"check_sum"`  // Notify withdraw args
}

// Withdraw creates a withdrawal request
// Parameters:
//   - args: Withdrawal request arguments
//
// Returns: Withdrawal result
func (b *BillingAPI) Withdraw(args *WithdrawArgs) (*types.WithdrawResult, error) {
	params := map[string]interface{}{
		"request_id": args.RequestID,
		"from_uid":   args.FromUID,
		"to_address": args.ToAddress,
		"amount":     args.Amount.String(),
		"symbol":     args.Symbol,
	}

	response, err := b.Post("/billing/withdraw", params)
	if err != nil {
		return nil, err
	}

	_, err = b.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.WithdrawResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WithdrawList gets withdrawal records by request IDs
// Parameters:
//   - requestIDs: List of request IDs
//
// Returns: Withdrawal records
func (b *BillingAPI) WithdrawList(requestIDs []string) (*types.WithdrawListResult, error) {
	params := map[string]interface{}{
		"ids": strings.Join(requestIDs, ","),
	}

	response, err := b.Post("/billing/withdrawList", params)
	if err != nil {
		return nil, err
	}

	_, err = b.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.WithdrawListResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SyncWithdrawList syncs withdrawal records by max ID (pagination)
// Parameters:
//   - maxID: Maximum transaction ID for pagination
//
// Returns: Synced withdrawal records
func (b *BillingAPI) SyncWithdrawList(maxID int64) (*types.WithdrawListResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := b.Post("/billing/syncWithdrawList", params)
	if err != nil {
		return nil, err
	}

	_, err = b.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.WithdrawListResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DepositList gets deposit records by WaaS IDs
// Parameters:
//   - ids: List of WaaS deposit IDs
//
// Returns: Deposit records
func (b *BillingAPI) DepositList(ids []int64) (*types.DepositListResult, error) {
	var strIDs []string
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}

	params := map[string]interface{}{
		"ids": strings.Join(strIDs, ","),
	}

	response, err := b.Post("/billing/depositList", params)
	if err != nil {
		return nil, err
	}

	_, err = b.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.DepositListResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SyncDepositList syncs deposit records by max ID (pagination)
// Parameters:
//   - maxID: Maximum transaction ID for pagination
//
// Returns: Synced deposit records
func (b *BillingAPI) SyncDepositList(maxID int64) (*types.DepositListResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := b.Post("/billing/syncDepositList", params)
	if err != nil {
		return nil, err
	}

	_, err = b.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.DepositListResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// MinerFeeList gets miner fee records by WaaS IDs
// Parameters:
//   - ids: List of WaaS transaction IDs
//
// Returns: Miner fee records
func (b *BillingAPI) MinerFeeList(ids []int64) (*types.MinerFeeListResult, error) {
	var strIDs []string
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}

	params := map[string]interface{}{
		"ids": strings.Join(strIDs, ","),
	}

	response, err := b.Post("/billing/minerFeeList", params)
	if err != nil {
		return nil, err
	}

	_, err = b.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.MinerFeeListResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SyncMinerFeeList syncs miner fee records by max ID (pagination)
// Parameters:
//   - maxID: Maximum transaction ID for pagination
//
// Returns: Synced miner fee records
func (b *BillingAPI) SyncMinerFeeList(maxID int64) (*types.MinerFeeListResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := b.Post("/billing/syncMinerFeeList", params)
	if err != nil {
		return nil, err
	}

	_, err = b.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.MinerFeeListResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
