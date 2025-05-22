package bean

import "encoding/json"

type Result struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (r *Result) ToJson() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (r *Result) IsSuccess() bool {
	return r.Code == 0
}
