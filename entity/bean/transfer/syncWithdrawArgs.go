package transfer

type SyncWithdrawArgs struct {
	MaxId string `json:"max_id"`
}

func (p SyncWithdrawArgs) Validate() bool {
	return p.MaxId == ""
}
