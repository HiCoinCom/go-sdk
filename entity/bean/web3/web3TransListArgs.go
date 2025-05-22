package web3

// Web3TransListArgs struct
type Web3TransListArgs struct {
	Ids string `json:"ids"`
}

func (p Web3TransListArgs) Validate() bool {
	return p.Ids == ""
}
