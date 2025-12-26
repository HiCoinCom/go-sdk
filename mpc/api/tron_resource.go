// Package api provides MPC API implementations
package api

import (
	"errors"
	"strings"

	"chainup.com/go-sdk/mpc/types"
)

// TronResourceAPI provides Tron resource operations
type TronResourceAPI struct {
	*MpcBaseAPI
}

// NewTronResourceAPI creates a new TronResourceAPI instance
func NewTronResourceAPI(config MpcConfigProvider) *TronResourceAPI {
	return &TronResourceAPI{
		MpcBaseAPI: NewMpcBaseAPI(config),
	}
}

// CreateTronDelegate creates a Tron delegate (buy resource)
// https://custodydocs-zh.chainup.com/api-references/mpc-apis/apis/tron/delegate-create
func (t *TronResourceAPI) CreateTronDelegate(args *types.TronBuyResourceArgs) (*types.TronBuyResourceResult, error) {
	if args == nil {
		return nil, errors.New("args cannot be nil")
	}
	if args.RequestID == "" {
		return nil, errors.New("parameter \"request_id\" is required")
	}
	if args.AddressFrom == "" {
		return nil, errors.New("parameter \"address_from\" is required")
	}
	if args.ServiceChargeType == "" {
		return nil, errors.New("parameter \"service_charge_type\" is required")
	}

	if args.BuyType == 0 || args.BuyType ==2 {
		if len(args.AddressTo) == 0 || len(args.ContractAddress) == 0 {
			return nil, errors.New("parameter \"address_to and contract_address\" is required")
		}
	}

	params := map[string]interface{}{
		"request_id":          args.RequestID,
		"buy_type":            args.BuyType,
		"resource_type":       args.ResourceType,
		"service_charge_type": args.ServiceChargeType,
		"address_from":        args.AddressFrom,
	}

	if args.EnergyNum > 0 {
		params["energy_num"] = args.EnergyNum
	}
	if args.NetNum > 0 {
		params["net_num"] = args.NetNum
	}
	if args.AddressTo != "" {
		params["address_to"] = args.AddressTo
	}
	if args.ContractAddress != "" {
		params["contract_address"] = args.ContractAddress
	}

	response, err := t.Post("/api/mpc/tron/delegate", params)
	if err != nil {
		return nil, err
	}

	result, err := t.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var buyResult types.TronBuyResourceResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &buyResult); err != nil {
		return nil, err
	}

	return &buyResult, nil
}

// GetBuyResourceRecords gets Tron resource purchase records
// https://custodydocs-zh.chainup.com/api-references/mpc-apis/apis/tron/delegate-record-list
func (t *TronResourceAPI) GetBuyResourceRecords(requestIds []string) (*types.TronBuyResourceRecordResult, error) {
	if len(requestIds) == 0 {
		return nil, errors.New("parameter \"request_ids\" is required")
	}

	params := map[string]interface{}{
		"ids": strings.Join(requestIds, ","),
	}

	response, err := t.Post("/api/mpc/tron/delegate/trans_list", params)
	if err != nil {
		return nil, err
	}

	result, err := t.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var recordResult types.TronBuyResourceRecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &recordResult); err != nil {
		return nil, err
	}

	return &recordResult, nil
}

// SyncBuyResourceRecords syncs Tron resource purchase records
// https://custodydocs-zh.chainup.com/api-references/mpc-apis/apis/tron/delegate-record-sync-list
func (t *TronResourceAPI) SyncBuyResourceRecords(maxId int) (*types.TronBuyResourceRecordResult, error) {
	params := map[string]interface{}{
		"max_id": maxId,
	}

	response, err := t.Post("/api/mpc/tron/delegate/sync_trans_list", params)
	if err != nil {
		return nil, err
	}

	result, err := t.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var recordResult types.TronBuyResourceRecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &recordResult); err != nil {
		return nil, err
	}

	return &recordResult, nil
}