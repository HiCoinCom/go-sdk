package wallet

type ChainHeightResult struct {
	Code string          `json:"code"`
	Data ChainHeightData `json:"data"`
	Msg  string          `json:"msg"`
}

type ChainHeightData struct {
	Height int64 `json:"height"`
}
