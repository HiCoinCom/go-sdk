// Package api provides API implementations for WaaS operations
package api

import (
	"chainup.com/go-sdk/custody/types"
)

// CoinAPI provides coin and blockchain information operations
type CoinAPI struct {
	*BaseAPI
}

// NewCoinAPI creates a new CoinAPI instance
func NewCoinAPI(config ConfigProvider) *CoinAPI {
	return &CoinAPI{
		BaseAPI: NewBaseAPI(config),
	}
}

// GetCoinList gets the list of supported cryptocurrencies
// Returns: List of coin information
func (c *CoinAPI) GetCoinList() (*types.CoinInfoListResult, error) {
	params := make(map[string]interface{})

	response, err := c.Post("/user/getCoinList", params)
	if err != nil {
		return nil, err
	}

	_, err = c.ValidateResponse(response)
	if err != nil {
		return nil, err
	}

	return mapToResult[types.CoinInfoListResult](response)
}
