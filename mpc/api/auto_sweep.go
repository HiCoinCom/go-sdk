// Package api provides MPC API implementations
package api

import (
	"errors"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"

	"chainup.com/go-sdk/mpc/types"
)

// AutoSweepAPI provides auto-sweep operations
type AutoSweepAPI struct {
	*MpcBaseAPI
}

// NewAutoSweepAPI creates a new AutoSweepAPI instance
func NewAutoSweepAPI(config MpcConfigProvider) *AutoSweepAPI {
	return &AutoSweepAPI{
		MpcBaseAPI: NewMpcBaseAPI(config),
	}
}

// AutoCollectSubWallets automatically collects from sub-wallets
// walletIDs: List of wallet IDs
// symbol: Coin symbol
func (a *AutoSweepAPI) AutoCollectSubWallets(walletIDs []int64, symbol string) (*types.AutoCollectResult, error) {
	if len(walletIDs) == 0 {
		return nil, errors.New("parameter \"sub_wallet_ids\" is required")
	}
	if symbol == "" {
		return nil, errors.New("parameter \"symbol\" is required")
	}

	// Convert wallet IDs to comma-separated string
	idStrs := make([]string, len(walletIDs))
	for i, id := range walletIDs {
		idStrs[i] = strconv.FormatInt(id, 10)
	}

	params := map[string]interface{}{
		"sub_wallet_ids": strings.Join(idStrs, ","),
		"symbol":         symbol,
	}

	response, err := a.Post("/api/mpc/auto_collect/sub_wallets", params)
	if err != nil {
		return nil, err
	}

	result, err := a.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var collectResult types.AutoCollectResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &collectResult); err != nil {
		return nil, err
	}

	return &collectResult, nil
}

// SetAutoCollectSymbol sets auto-collection symbol configuration
// args: Auto collect symbol arguments (symbol, collectMin, fuelingLimit)
func (a *AutoSweepAPI) SetAutoCollectSymbol(args *types.SetAutoCollectSymbolArgs) (bool, error) {
	if args == nil {
		return false, errors.New("args cannot be nil")
	}
	if args.Symbol == "" {
		return false, errors.New("parameter \"symbol\" is required")
	}
	if args.CollectMin.LessThanOrEqual(decimal.Zero) {
		return false, errors.New("parameter \"collect_min\" is required")
	}

	if args.FuelingLimit.LessThanOrEqual(decimal.Zero) {
		return false, errors.New("parameter \"fueling_limit\" is required")
	}

	params := map[string]interface{}{
		"symbol":        args.Symbol,
		"collect_min":   args.CollectMin,
		"fueling_limit": args.FuelingLimit,
	}

	response, err := a.Post("/api/mpc/auto_collect/symbol/set", params)
	if err != nil {
		return false, err
	}

	result, err := a.ValidateResponse(response)
	if err != nil {
		return false, err
	}

	var collectResult types.AutoCollectResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &collectResult); err != nil {
		return false, err
	}

	return collectResult.Code == "0", nil
}

// SyncAutoCollectRecords syncs auto-collection records
// maxID: Starting record ID, default is 0
func (a *AutoSweepAPI) SyncAutoCollectRecords(maxID int64) (*types.AutoCollectRecordResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := a.Get("/api/mpc/billing/sync_auto_collect_list", params)
	if err != nil {
		return nil, err
	}

	result, err := a.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var recordResult types.AutoCollectRecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &recordResult); err != nil {
		return nil, err
	}

	return &recordResult, nil
}
