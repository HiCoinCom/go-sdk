// Package api provides MPC API implementations
package api

import (
	"errors"
	"strconv"
	"strings"

	"chainup.com/go-sdk/mpc/types"
)

// DepositAPI provides deposit record operations
type DepositAPI struct {
	*MpcBaseAPI
}

// NewDepositAPI creates a new DepositAPI instance
func NewDepositAPI(config MpcConfigProvider) *DepositAPI {
	return &DepositAPI{
		MpcBaseAPI: NewMpcBaseAPI(config),
	}
}

// GetDepositRecords gets deposit records by IDs
// ids: List of deposit IDs (up to 100)
func (d *DepositAPI) GetDepositRecords(ids []int64) (*types.DepositRecordResult, error) {
	if len(ids) == 0 {
		return nil, errors.New("parameter \"ids\" is required and must be a non-empty array")
	}

	// Convert IDs to comma-separated string
	idStrs := make([]string, len(ids))
	for i, id := range ids {
		idStrs[i] = strconv.FormatInt(id, 10)
	}

	params := map[string]interface{}{
		"ids": strings.Join(idStrs, ","),
	}

	response, err := d.Get("/api/mpc/billing/deposit_list", params)
	if err != nil {
		return nil, err
	}

	result, err := d.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var depositResult types.DepositRecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &depositResult); err != nil {
		return nil, err
	}

	return &depositResult, nil
}

// SyncDepositRecords syncs deposit records by max ID
// maxID: Deposit record initial ID, default is 0
func (d *DepositAPI) SyncDepositRecords(maxID int64) (*types.DepositRecordResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := d.Get("/api/mpc/billing/sync_deposit_list", params)
	if err != nil {
		return nil, err
	}

	result, err := d.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var depositResult types.DepositRecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &depositResult); err != nil {
		return nil, err
	}

	return &depositResult, nil
}
