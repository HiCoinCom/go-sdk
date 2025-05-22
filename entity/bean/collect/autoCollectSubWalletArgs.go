package collect

// AutoCollectSubWalletArgs struct
type AutoCollectSubWalletArgs struct {
	Symbol string `json:"symbol"`
}

func (p AutoCollectSubWalletArgs) Validate() bool {
	return p.Symbol == ""
}
