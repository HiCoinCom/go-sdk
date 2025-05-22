package web3

// Web3TransCreateArgs struct
type Web3TransCreateArgs struct {
	SubWalletId         int64  `json:"sub_wallet_id"`
	RequestId           string `json:"request_id"`
	MainChainSymbol     string `json:"main_chain_symbol"`
	From                string `json:"from"`
	InteractiveContract string `json:"interactive_contract"`
	Amount              string `json:"amount"`
	GasPrice            string `json:"gas_price"`
	GasLimit            string `json:"gas_limit"`
	InputData           string `json:"input_data"`
	TransType           string `json:"trans_type"`
	DappName            string `json:"dapp_name"`
	DappUrl             string `json:"dapp_url"`
	DappImg             string `json:"dapp_img"`
	Sign                string `json:"sign"`
}

func (p Web3TransCreateArgs) Validate() bool {
	return p.SubWalletId == 0 || p.RequestId == "" || p.MainChainSymbol == "" || p.InteractiveContract == "" || p.Amount == "" || p.GasPrice == "" || p.GasLimit == "" || p.InputData == "" || p.TransType == ""
}
