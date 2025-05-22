package workstation

// SubWalletAddressInfoArgs struct
type SubWalletAddressInfoArgs struct {
	Address string `json:"address"`
	Memo    string `json:"memo"`
}

func (p SubWalletAddressInfoArgs) Validate() bool {
	return p.Address == ""
}
