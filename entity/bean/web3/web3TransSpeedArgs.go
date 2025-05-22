package web3

// Web3TransSpeedArgs struct
type Web3TransSpeedArgs struct {
	TransId  int64  `json:"trans_id"`
	GasPrice string `json:"gas_price"`
	GasLimit string `json:"gas_limit"`
}

func (p Web3TransSpeedArgs) Validate() bool {
	return p.TransId == 0
}
