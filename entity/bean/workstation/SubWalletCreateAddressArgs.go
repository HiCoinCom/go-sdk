package workstation

// CreateWalletArgs struct
type SubWalletCreateAddressArgs struct {
	SubWalletId int64  `json:"sub_wallet_id"`
	Symbol      string `json:"symbol"`
}

func (p SubWalletCreateAddressArgs) Validate() bool {
	return p.SubWalletId == 0 || p.Symbol == ""
}
