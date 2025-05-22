package wallet

type CoinListArgs struct {
	Time       int64  `json:"time"`
	Charset    string `json:"charset"`
	Symbol     string `json:"symbol"`
	BaseSymbol string `json:"base_symbol"`
	OpenChain  bool   `json:"open_chain"`
}

func (p CoinListArgs) Validate() bool {
	return p.Time == 0 || p.Charset == ""
}
