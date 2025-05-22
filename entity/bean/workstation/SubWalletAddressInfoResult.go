package workstation

// SubWalletAddressInfoResult struct
type SubWalletAddressInfoResult struct {
	Code string `json:"code"`
	Data struct {
		SubWalletId        int64  `json:"sub_wallet_id"`
		AddrType           int64  `json:"addr_type"`
		MergeAddressSymbol string `json:"merge_address_symbol"`
	} `json:"data"`
	Msg string `json:"msg"`
}
