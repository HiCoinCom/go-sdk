package deposit

type DepositListArgs struct {
	Ids string `json:"ids"`
}

func (p DepositListArgs) Validate() bool {
	return p.Ids == ""
}
