package deposit

type DepositListResult struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []DepositListData `json:"data"`
}

type DepositListData struct {
	Symbol          string `json:"symbol"`
	Amount          string `json:"amount"`
	AddressTo       string `json:"address_to"`
	Memo            string `json:"memo"`
	CreatedAt       int64  `json:"created_at"`
	Txid            string `json:"txid"`
	Confirmations   int64  `json:"confirmations"`
	ContractAddress string `json:"contract_address"`
	SubWalletId     int64  `json:"sub_wallet_id"`
	AddressFrom     string `json:"address_from"`
	UpdatedAt       int64  `json:"updated_at"`
	BaseSymbol      string `json:"base_symbol"`
	Id              int64  `json:"id"`
	Status          int64  `json:"status"`
	TxHeight        int64  `json:"tx_height"`
}
