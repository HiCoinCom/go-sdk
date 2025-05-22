package collect

// SetAutoCollectSymbolResult struct
type SetAutoCollectSymbolResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Symbol string `json:"symbol"`
	}
}
