package workstation

// SubWalletAssetsArgs struct
type SubWalletAssetsArgs struct {
	Symbol      string `json:"symbol"`
	SubWalletId int64  `json:"sub_wallet_id"`
}

func (p SubWalletAssetsArgs) Validate() bool {
	return p.Symbol == "" || p.SubWalletId == 0
}
