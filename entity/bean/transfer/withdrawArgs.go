package transfer

type WithdrawArgs struct {
	SubWalletId int64  `json:"sub_wallet_id"`
	Symbol      string `json:"symbol"`
	From        string `json:"from"`
	AddressTo   string `json:"address_to"`
	Memo        string `json:"memo"`
	Amount      string `json:"amount"`
	RequestId   string `json:"request_id"`
	Remark      string `json:"remark"`
	Sign        string `json:"sign"`
}

func (p WithdrawArgs) Validate() bool {
	return p.SubWalletId == 0 || p.Symbol == "" || p.AddressTo == "" || p.Amount == "" || p.RequestId == ""
}
