package workstation

type SubWalletAssetsResult struct {
	Code string `json:"code"`
	Data struct {
		CollectingBalance string `json:"collecting_balance"`
		NormalBalance     string `json:"normal_balance"`
		LockBalance       string `json:"lock_balance"`
	}
	Msg string `json:"msg"`
}
