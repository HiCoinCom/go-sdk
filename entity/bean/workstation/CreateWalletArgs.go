package workstation

type CreateWalletArgs struct {
	SubWalletName string `json:"sub_wallet_name"`
	AppShowStatus string `json:"app_show_status"`
}

func (p CreateWalletArgs) Validate() bool {
	return p.SubWalletName == ""
}
