package transfer

type WithdrawListArgs struct {
	Ids string `json:"ids"`
}

func (p WithdrawListArgs) Validate() bool {
	return p.Ids == ""
}
