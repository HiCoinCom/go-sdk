// Package api provides API implementations for WaaS operations
package api

import (
	"encoding/json"

	"chainup.com/go-sdk/custody/types"
)

// UserAPI provides user management and registration operations
type UserAPI struct {
	*BaseAPI
}

// NewUserAPI creates a new UserAPI instance
func NewUserAPI(config ConfigProvider) *UserAPI {
	return &UserAPI{
		BaseAPI: NewBaseAPI(config),
	}
}

// mapToResult converts map response to typed result
func mapToResult[T any](response map[string]interface{}) (*T, error) {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	var result T
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RegisterMobileUser registers a new user using mobile phone
// Parameters:
//   - country: Country code (e.g., '86')
//   - mobile: Mobile phone number
//
// Returns: User registration result containing uid
func (u *UserAPI) RegisterMobileUser(country, mobile string) (*types.UserInfoResult, error) {
	params := map[string]interface{}{
		"country": country,
		"mobile":  mobile,
	}

	response, err := u.Post("/user/createUser", params)
	if err != nil {
		return nil, err
	}

	_, err = u.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.UserInfoResult](response)
}

// RegisterEmailUser registers a new user using email
// Parameters:
//   - email: Email address
//
// Returns: User registration result containing uid
func (u *UserAPI) RegisterEmailUser(email string) (*types.UserInfoResult, error) {
	params := map[string]interface{}{
		"email": email,
	}

	response, err := u.Post("/user/registerEmail", params)
	if err != nil {
		return nil, err
	}

	_, err = u.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.UserInfoResult](response)
}

// GetMobileUser gets user information by mobile phone
// Parameters:
//   - country: Country code (e.g., '86')
//   - mobile: Mobile phone number
//
// Returns: User information
func (u *UserAPI) GetMobileUser(country, mobile string) (*types.UserInfoResult, error) {
	params := map[string]interface{}{
		"country": country,
		"mobile":  mobile,
	}

	response, err := u.Post("/user/info", params)
	if err != nil {
		return nil, err
	}

	_, err = u.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.UserInfoResult](response)
}

// GetEmailUser gets user information by email
// Parameters:
//   - email: User email
//
// Returns: User information
func (u *UserAPI) GetEmailUser(email string) (*types.UserInfoResult, error) {
	params := map[string]interface{}{
		"email": email,
	}

	response, err := u.Post("/user/info", params)
	if err != nil {
		return nil, err
	}

	_, err = u.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.UserInfoResult](response)
}

// SyncUserList syncs user list by max ID (pagination)
// Parameters:
//   - maxID: Maximum user ID for pagination (0 for first sync)
//
// Returns: Synced user list
func (u *UserAPI) SyncUserList(maxID int64) (*types.UserListResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := u.Post("/user/syncList", params)
	if err != nil {
		return nil, err
	}

	_, err = u.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.UserListResult](response)
}
