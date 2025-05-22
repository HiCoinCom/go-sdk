package web3

type Web3TransCreateResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TransId int64 `json:"trans_id"`
	}
}
