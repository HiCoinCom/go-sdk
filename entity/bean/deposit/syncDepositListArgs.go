package deposit

// SyncDepositListArgs struct
type SyncDepositListArgs struct {
	MaxId string `json:"max_id"`
}

func (p SyncDepositListArgs) Validate() bool {
	return p.MaxId == ""
}
