package workstation

type CreateWalletResult struct {
	Code string `json:"code"`
	Data struct {
		SubWalletId int64 `json:"sub_wallet_id"`
	}
	Msg string `json:"msg"`
}
