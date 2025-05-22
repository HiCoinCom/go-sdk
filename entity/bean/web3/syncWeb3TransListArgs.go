package web3

// SyncWeb3TransListArgs struct
type SyncWeb3TransListArgs struct {
	MaxId string `json:"max_id"`
}

func (p SyncWeb3TransListArgs) Validate() bool {
	return p.MaxId == ""
}
