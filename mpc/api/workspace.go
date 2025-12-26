// Package api provides MPC API implementations
package api

import (
	"errors"

	"chainup.com/go-sdk/mpc/types"
)

// WorkSpaceAPI provides workspace operations
type WorkSpaceAPI struct {
	*MpcBaseAPI
}

// NewWorkSpaceAPI creates a new WorkSpaceAPI instance
func NewWorkSpaceAPI(config MpcConfigProvider) *WorkSpaceAPI {
	return &WorkSpaceAPI{
		MpcBaseAPI: NewMpcBaseAPI(config),
	}
}

// GetSupportMainChain gets supported main chains
func (w *WorkSpaceAPI) GetSupportMainChain() (*types.SupportMainChainResult, error) {
	response, err := w.Get("/api/mpc/wallet/open_coin", nil)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var chainResult types.SupportMainChainResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &chainResult); err != nil {
		return nil, err
	}

	return &chainResult, nil
}

// GetCoinDetails gets coin details
// args: Coin details query arguments (symbol, contractAddress, showBalance, maxId, limit)
func (w *WorkSpaceAPI) GetCoinDetails(args *types.GetCoinDetailsArgs) (*types.CoinDetailsResult, error) {
	params := map[string]interface{}{}

	if args != nil {
		if args.Symbol != "" {
			params["symbol"] = args.Symbol
		}
		if args.ContractAddress != "" {
			params["contract_address"] = args.ContractAddress
		}
		if args.ShowBalance {
			params["show_balance"] = true
		}
		if args.MaxID > 0 {
			params["max_id"] = args.MaxID
		}
		if args.Limit > 0 {
			params["limit"] = args.Limit
		}
	}

	response, err := w.Get("/api/mpc/coin_list", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var coinResult types.CoinDetailsResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &coinResult); err != nil {
		return nil, err
	}

	return &coinResult, nil
}

// GetLastBlockHeight gets the latest block height
// symbol: Main chain symbol (e.g., "ETH")
func (w *WorkSpaceAPI) GetLastBlockHeight(symbol string) (*types.BlockHeightResult, error) {
	if symbol == "" {
		return nil, errors.New("parameter \"symbol\" is required")
	}

	params := map[string]interface{}{
		"base_symbol": symbol,
	}

	response, err := w.Get("/api/mpc/chain_height", params)
	if err != nil {
		return nil, err
	}

	result, err := w.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	var heightResult types.BlockHeightResult
	if err := SafeUnmarshalResponse(result.(map[string]interface{}), &heightResult); err != nil {
		return nil, err
	}

	return &heightResult, nil
}
