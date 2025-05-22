package workstation

// SubWalletCreateAddressResult struct
type SubWalletCreateAddressResult struct {
	Code string `json:"code"`
	Data struct {
		Address     string `json:"address"`
		AddressType int64  `json:"address_type"`
		Memo        string `json:"memo"`
	}
	Msg string `json:"msg"`
}
