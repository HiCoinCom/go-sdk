package collect

// AutoCollectSubWalletResult struct
type AutoCollectSubWalletResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CollectSubWalletId int64 `json:"collect_sub_wallet_id"`
		FuelingSubWalletId int64 `json:"fueling_sub_wallet_id"`
	}
}
