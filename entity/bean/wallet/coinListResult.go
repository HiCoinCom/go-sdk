package wallet

type CoinListResult struct {
	Code string         `json:"code"`
	Data []CoinListData `json:"data"`
	Msg  string         `json:"msg"`
}

type CoinListData struct {
	AddressRegex        string `json:"address_regex"`
	AddressTagRegex     string `json:"address_tag_regex"`
	BaseSymbol          string `json:"base_symbol"`
	CoinNet             string `json:"coin_net"`
	ContractAddress     string `json:"contract_address"`
	Decimals            string `json:"decimals"`
	DepositConfirmation string `json:"deposit_confirmation"`
	AddressLink         string `json:"address_link"`
	TxidLink            string `json:"txid_link"`
	Icon                string `json:"icon"`
	IfOpenChain         bool   `json:"if_open_chain"`
	RealSymbol          string `json:"real_symbol"`
	SupportMemo         string `json:"support_memo"`
	SupportToken        string `json:"support_token"`
	Symbol              string `json:"symbol"`
	SymbolAlias         string `json:"symbol_alias"`
	SupportAcceleration bool   `json:"support_acceleration"`
	SupportMultiAddr    bool   `json:"support_multi_addr"`
	MergeAddressSymbol  string `json:"merge_address_symbol"`
}
