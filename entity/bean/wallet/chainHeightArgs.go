package wallet

type ChainHeightArgs struct {
	BaseSymbol string `json:"base_symbol"`
}

func (p ChainHeightArgs) Validate() bool {
	return p.BaseSymbol == ""
}
