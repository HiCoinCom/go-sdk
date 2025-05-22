package transfer

type SyncWithdrawResult struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data []SyncWithdrawData `json:"data"`
}

type SyncWithdrawData struct {
	Symbol          string `json:"symbol"`
	Amount          string `json:"amount"`
	RealFee         string `json:"real_fee"`
	WithdrawSource  int64  `json:"withdraw_source"`
	Fee             string `json:"fee"`
	AddressTo       string `json:"address_to"`
	Memo            string `json:"memo"`
	CreatedAt       int64  `json:"created_at"`
	Txid            string `json:"txid"`
	Confirmations   int64  `json:"confirmations"`
	ContractAddress string `json:"contract_address"`
	SubWalletId     int64  `json:"sub_wallet_id"`
	AddressFrom     string `json:"address_from"`
	FeeSymbol       string `json:"fee_symbol"`
	UpdatedAt       int64  `json:"updated_at"`
	BaseSymbol      string `json:"base_symbol"`
	Id              int64  `json:"id"`
	RequestId       string `json:"request_id"`
	Status          int64  `json:"status"`
	TxHeight        int64  `json:"tx_height"`
}
