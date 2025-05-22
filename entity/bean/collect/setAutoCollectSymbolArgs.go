package collect

import "github.com/shopspring/decimal"

// SetCollectSubWalletArgs struct
type SetCollectSubWalletArgs struct {
	Symbol       string `json:"symbol"`
	CollectMin   string `json:"collect_min"`
	FuelingLimit string `json:"fueling_limit"`
}

func (p SetCollectSubWalletArgs) Validate() bool {
	if p.Symbol == "" || p.CollectMin == "" || p.FuelingLimit == "" {
		return false
	}
	maxInfo, _ := decimal.NewFromString("9999999999999999")
	flag := checkAmount(p.CollectMin, maxInfo)
	flag = checkAmount(p.FuelingLimit, maxInfo)
	return flag
}

func checkAmount(amount string, maxAmount decimal.Decimal) bool {
	collectMin, err := decimal.NewFromString(amount)
	if err != nil {
		return false
	}
	if collectMin.Cmp(maxAmount) > 0 {
		return false
	}
	if collectMin.Cmp(decimal.Zero) < 0 {
		return false
	}
	// 返回 -小数位数。最多不超过6位小数
	if collectMin.Exponent() < -6 {
		return false
	}
	return true
}
