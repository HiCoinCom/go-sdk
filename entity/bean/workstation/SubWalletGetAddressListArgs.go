package workstation

// SubWalletGetAddressListArgs struct
type SubWalletGetAddressListArgs struct {
	SubWalletId int64  `json:"sub_wallet_id"`
	Symbol      string `json:"symbol"`
	MaxId       int64  `json:"max_id"`
}

func (p SubWalletGetAddressListArgs) Validate() bool {
	return p.SubWalletId == 0 || p.Symbol == "" || p.MaxId < 0
}
