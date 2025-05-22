package workstation

// SubWalletChangeShowArgs struct
type SubWalletChangeShowArgs struct {
	SubWalletIds  string `json:"sub_wallet_ids"`
	AppShowStatus string `json:"app_show_status"`
}

func (p SubWalletChangeShowArgs) Validate() bool {
	return p.SubWalletIds == "" || p.AppShowStatus == ""
}
