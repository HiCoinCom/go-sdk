package workstation

// SubWalletGetAddressListResult struct
type SubWalletGetAddressListResult struct {
	Code string                     `json:"code"`
	Data []SubWalletAddressListData `json:"data"`
	Msg  string                     `json:"msg"`
}

// SubWalletAddressListData struct
type SubWalletAddressListData struct {
	Address  string `json:"address"`
	AddrType int64  `json:"addr_type"`
	Memo     string `json:"memo"`
}
