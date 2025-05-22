package web3

// SyncWeb3TransListResult struct
type SyncWeb3TransListResult struct {
	Code string                  `json:"code"`
	Msg  string                  `json:"msg"`
	Data []SyncWeb3TransListData `json:"data"`
}

// SyncWeb3TransListData struct
type SyncWeb3TransListData struct {
	Id                  int64  `json:"id"`
	RequestId           string `json:"request_id"`
	SubWalletId         int64  `json:"sub_wallet_id"`
	Txid                string `json:"txid"`
	Symbol              string `json:"symbol"`
	MainChainSymbol     string `json:"main_chain_symbol"`
	Amount              string `json:"amount"`
	FeeSymbol           string `json:"fee_symbol"`
	Fee                 string `json:"fee"`
	RealFee             string `json:"real_fee"`
	CreatedAt           int64  `json:"created_at"`
	UpdatedAt           int64  `json:"updated_at"`
	AddressFrom         string `json:"address_from"`
	AddressTo           string `json:"address_to"`
	Confirmations       int64  `json:"confirmations"`
	InputData           string `json:"input_data"`
	InteractiveContract string `json:"interactive_contract"`
	Status              int64  `json:"status"`
	TransSource         int64  `json:"trans_source"`
	TxHeight            int64  `json:"tx_height"`
}
