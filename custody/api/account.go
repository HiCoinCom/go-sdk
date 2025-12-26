// Package api provides API implementations for WaaS operations
package api

import (
	"chainup.com/go-sdk/custody/types"
)

// AccountAPI provides account and balance management operations
type AccountAPI struct {
	*BaseAPI
}

// NewAccountAPI creates a new AccountAPI instance
func NewAccountAPI(config ConfigProvider) *AccountAPI {
	return &AccountAPI{
		BaseAPI: NewBaseAPI(config),
	}
}

// GetUserAccount gets user account balance for a specific cryptocurrency
// Parameters:
//   - uid: User ID
//   - symbol: Cryptocurrency symbol (e.g., 'BTC', 'ETH')
//
// Returns: Account balance information
func (a *AccountAPI) GetUserAccount(uid int64, symbol string) (*types.AccountResult, error) {
	params := map[string]interface{}{
		"uid":    uid,
		"symbol": symbol,
	}

	response, err := a.Post("/account/getByUidAndSymbol", params)
	if err != nil {
		return nil, err
	}

	_, err = a.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.AccountResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserAddress gets user deposit address for a specific cryptocurrency
// Parameters:
//   - uid: User ID
//   - symbol: Cryptocurrency symbol (e.g., 'BTC', 'ETH')
//
// Returns: Deposit address information
func (a *AccountAPI) GetUserAddress(uid int64, symbol string) (*types.UserAddressResult, error) {
	params := map[string]interface{}{
		"uid":    uid,
		"symbol": symbol,
	}

	response, err := a.Post("/account/getDepositAddress", params)
	if err != nil {
		return nil, err
	}

	_, err = a.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.UserAddressResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCompanyAccount gets company (merchant) account balance for a specific cryptocurrency
// Parameters:
//   - symbol: Cryptocurrency symbol (e.g., 'BTC', 'ETH')
//
// Returns: Company account information
func (a *AccountAPI) GetCompanyAccount(symbol string) (*types.CompanyAccountResult, error) {
	params := map[string]interface{}{
		"symbol": symbol,
	}

	response, err := a.Post("/account/getCompanyBySymbol", params)
	if err != nil {
		return nil, err
	}

	_, err = a.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.CompanyAccountResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserAddressInfo gets user address information by address
// Parameters:
//   - address: Blockchain address to query
//
// Returns: Address details
func (a *AccountAPI) GetUserAddressInfo(address string) (*types.UserAddressResult, error) {
	params := map[string]interface{}{
		"address": address,
	}

	response, err := a.Post("/account/getDepositAddressInfo", params)
	if err != nil {
		return nil, err
	}

	_, err = a.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.UserAddressResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SyncUserAddressList syncs user address list by max ID (pagination)
// Parameters:
//   - maxID: Maximum address ID for pagination (0 for first sync)
//
// Returns: Synced user address list with id, uid, address, symbol
func (a *AccountAPI) SyncUserAddressList(maxID int64) (*types.UserAddressListResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := a.Post("/address/syncList", params)
	if err != nil {
		return nil, err
	}

	_, err = a.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var result types.UserAddressListResult
	if err := unmarshalResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
