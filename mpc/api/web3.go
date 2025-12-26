// Package api provides MPC API implementations
package api

import (
	"errors"
	"fmt"
	"strings"

	"chainup.com/go-sdk/mpc/types"
	"chainup.com/go-sdk/utils/mpcsign"
)

// Web3API provides Web3 transaction operations
type Web3API struct {
	*MpcBaseAPI
}

// NewWeb3API creates a new Web3API instance
func NewWeb3API(config MpcConfigProvider) *Web3API {
	return &Web3API{
		MpcBaseAPI: NewMpcBaseAPI(config),
	}
}

// CreateWeb3Trans creates a Web3 transaction
// req: Web3 transaction request parameters
// needTransactionSign: Whether to sign the transaction (requires signPrivateKey in config)
func (w *Web3API) CreateWeb3Trans(req *types.Web3TransRequest, needTransactionSign bool) (*types.Web3TransResponse, error) {
	if req == nil {
		return nil, errors.New("web3 transaction request is required")
	}

	// Build params map
	params := make(map[string]interface{})
	params["request_id"] = req.RequestID
	params["sub_wallet_id"] = req.WalletID
	params["main_chain_symbol"] = req.MainChainSymbol
	params["interactive_contract"] = req.InteractiveContract
	params["amount"] = req.Amount.String()
	params["gas_price"] = req.GasPrice.String()
	params["gas_limit"] = req.GasLimit
	params["input_data"] = req.InputData
	params["trans_type"] = req.TransType

	if req.From != "" {
		params["from"] = req.From
	}
	if req.DappName != "" {
		params["dapp_name"] = req.DappName
	}
	if req.DappURL != "" {
		params["dapp_url"] = req.DappURL
	}
	if req.DappImg != "" {
		params["dapp_img"] = req.DappImg
	}

	// Generate signature if needed
	if needTransactionSign {
		signProvider := w.config.GetCryptoProvider()
		if signProvider == nil {
			return nil, fmt.Errorf("crypto provider is required when needTransactionSign is true")
		}

		// Build sign params
		signParams := map[string]string{
			"request_id":           req.RequestID,
			"sub_wallet_id":        fmt.Sprintf("%d", req.WalletID),
			"main_chain_symbol":    req.MainChainSymbol,
			"interactive_contract": req.InteractiveContract,
			"amount":               req.Amount.String(),
			"input_data":           req.InputData,
		}

		signature, err := mpcsign.GenerateWeb3Sign(signParams, signProvider)
		if err != nil {
			return nil, fmt.Errorf("failed to generate signature: %w", err)
		}

		params["sign"] = signature
	}

	response, err := w.Post("/api/mpc/web3/trans/create", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var web3Resp types.Web3TransResponse
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &web3Resp); err != nil {
		return nil, err
	}

	return &web3Resp, nil
}

// AccelerationWeb3Trans accelerates a Web3 transaction
// args: Acceleration arguments (request_id, gas_price)
func (w *Web3API) AccelerationWeb3Trans(args *types.Web3AccelerationArgs) (bool, error) {
	if args == nil {
		return false, errors.New("acceleration args is required")
	}

	params := map[string]interface{}{
		"request_id": args.RequestID,
		"gas_price":  args.GasPrice.String(),
	}

	response, err := w.Post("/api/mpc/web3/acceleration_transaction", params)
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

// GetWeb3Records gets Web3 transaction records by request IDs
// requestIDs: List of request IDs
func (w *Web3API) GetWeb3Records(requestIDs []string) (*types.Web3RecordResult, error) {
	if len(requestIDs) == 0 {
		return nil, errors.New("parameter \"request_ids\" is required and must be a non-empty array")
	}

	params := map[string]interface{}{
		"ids": strings.Join(requestIDs, ","),
	}

	response, err := w.Get("/api/mpc/web3/trans_list", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var web3Result types.Web3RecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &web3Result); err != nil {
		return nil, err
	}

	return &web3Result, nil
}

// SyncWeb3Records syncs Web3 transaction records by max ID
// maxID: Web3 record initial ID, default is 0
func (w *Web3API) SyncWeb3Records(maxID int64) (*types.Web3RecordResult, error) {
	params := map[string]interface{}{
		"max_id": maxID,
	}

	response, err := w.Get("/api/mpc/web3/sync_trans_list", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var web3Result types.Web3RecordResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &web3Result); err != nil {
		return nil, err
	}

	return &web3Result, nil
}
