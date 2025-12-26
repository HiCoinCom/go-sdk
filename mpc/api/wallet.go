// Package api provides MPC API implementations
package api

import (
	"errors"
	"strconv"
	"strings"

	"chainup.com/go-sdk/mpc/types"
)

// WalletAPI provides MPC wallet management operations
type WalletAPI struct {
	*MpcBaseAPI
}

// NewWalletAPI creates a new WalletAPI instance
func NewWalletAPI(config MpcConfigProvider) *WalletAPI {
	return &WalletAPI{
		MpcBaseAPI: NewMpcBaseAPI(config),
	}
}

// CreateWallet creates a new wallet
// walletName: Wallet name (max 50 characters)
// showStatus: Display status: 1 (show), 2 (hide, default)
func (w *WalletAPI) CreateWallet(walletName string, showStatus types.AppShowStatus) (*types.WalletCreateResult, error) {
	if walletName == "" {
		return nil, errors.New("parameter \"sub_wallet_name\" is required")
	}

	if len(walletName) > 50 {
		return nil, errors.New("wallet name cannot be longer than 50 characters")
	}

	if showStatus == 0 {
		showStatus = types.AppShowStatusHidden
	}

	params := map[string]interface{}{
		"sub_wallet_name": walletName,
		"app_show_status": int(showStatus),
	}

	response, err := w.Post("/api/mpc/sub_wallet/create", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	// Parse response
	var walletResult types.WalletCreateResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &walletResult); err != nil {
		return nil, err
	}

	return &walletResult, nil
}

// CreateWalletAddress creates a wallet address
// walletID: Wallet ID
// symbol: Unique identifier for the coin (e.g., "ETH")
func (w *WalletAPI) CreateWalletAddress(walletID int64, symbol string) (*types.WalletAddressResult, error) {
	if walletID == 0 {
		return nil, errors.New("parameter \"sub_wallet_id\" is required")
	}

	if symbol == "" {
		return nil, errors.New("parameter \"symbol\" is required")
	}

	params := map[string]interface{}{
		"sub_wallet_id": walletID,
		"symbol":        symbol,
	}

	response, err := w.Post("/api/mpc/sub_wallet/create/address", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var addressResult types.WalletAddressResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &addressResult); err != nil {
		return nil, err
	}

	return &addressResult, nil
}

// QueryWalletAddress queries wallet addresses
// args: Query arguments (walletID, symbol, maxID)
func (w *WalletAPI) QueryWalletAddress(args *types.QueryWalletAddressArgs) (*types.WalletAddressListResult, error) {
	if args == nil {
		return nil, errors.New("args cannot be nil")
	}
	if args.WalletID == 0 {
		return nil, errors.New("parameter \"sub_wallet_id\" is required")
	}

	if args.Symbol == "" {
		return nil, errors.New("parameter \"symbol\" is required")
	}

	params := map[string]interface{}{
		"sub_wallet_id": args.WalletID,
		"symbol":        args.Symbol,
		"max_id":        args.MaxID,
	}

	response, err := w.Post("/api/mpc/sub_wallet/get/address/list", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var listResult types.WalletAddressListResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &listResult); err != nil {
		return nil, err
	}

	return &listResult, nil
}

// GetWalletAssets gets wallet assets
// walletID: Wallet ID
// symbol: Unique identifier for the coin (e.g., "ETH")
func (w *WalletAPI) GetWalletAssets(walletID int64, symbol string) (*types.WalletAssetsResult, error) {
	if walletID == 0 {
		return nil, errors.New("parameter \"sub_wallet_id\" is required")
	}

	if symbol == "" {
		return nil, errors.New("parameter \"symbol\" is required")
	}

	params := map[string]interface{}{
		"sub_wallet_id": walletID,
		"symbol":        symbol,
	}

	response, err := w.Get("/api/mpc/sub_wallet/assets", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var assetsResult types.WalletAssetsResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &assetsResult); err != nil {
		return nil, err
	}

	return &assetsResult, nil
}

// ChangeWalletShowStatus modifies the wallet display status
// walletIDs: List of wallet IDs
// showStatus: Display status: 1 (show), 2 (hide)
func (w *WalletAPI) ChangeWalletShowStatus(walletIDs []int64, showStatus types.AppShowStatus) (bool, error) {
	if len(walletIDs) == 0 {
		return false, errors.New("parameter \"sub_wallet_ids\" is required")
	}

	if showStatus != types.AppShowStatusShow && showStatus != types.AppShowStatusHidden {
		return false, errors.New("parameter \"app_show_status\" must be 1 or 2")
	}

	// Convert wallet IDs to comma-separated string
	idStrs := make([]string, len(walletIDs))
	for i, id := range walletIDs {
		idStrs[i] = strconv.FormatInt(id, 10)
	}

	params := map[string]interface{}{
		"sub_wallet_ids":  strings.Join(idStrs, ","),
		"app_show_status": int(showStatus),
	}

	response, err := w.Post("/api/mpc/sub_wallet/change_show_status", params)
	if err != nil {
		return false, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return false, err
	}

	// Check if code is "0"
	if resultMap, ok := result.(map[string]interface{}); ok {
		if code, ok := resultMap["code"].(string); ok && code == "0" {
			return true, nil
		}
	}

	return false, nil
}

// WalletAddressInfo gets wallet address info
// address: Any address
// memo: If it's a Memo type, input the memo
func (w *WalletAPI) WalletAddressInfo(address, memo string) (*types.WalletAddressInfoResult, error) {
	if address == "" {
		return nil, errors.New("parameter \"address\" is required")
	}

	params := map[string]interface{}{
		"address": address,
	}
	if memo != "" {
		params["memo"] = memo
	}

	response, err := w.Get("/api/mpc/sub_wallet/address/info", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var infoResult types.WalletAddressInfoResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &infoResult); err != nil {
		return nil, err
	}

	return &infoResult, nil
}
