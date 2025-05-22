package wallet

type OpenCoinArgs struct {
	Time    int64  `json:"time"`
	Charset string `json:"charset"`
}

func (p OpenCoinArgs) Validate() bool {
	return p.Time == 0 || p.Charset == ""
}
