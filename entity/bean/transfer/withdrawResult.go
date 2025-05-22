package transfer

type WithdrawResult struct {
	Code string `json:"code"`
	Data struct {
		WithdrawId int64 `json:"withdraw_id"`
	} `json:"data"`
	Msg string `json:"msg"`
}
