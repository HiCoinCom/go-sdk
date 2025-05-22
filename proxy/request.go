package proxy

import (
	"chainup.com/go-sdk/entity/bean"
	"chainup.com/go-sdk/entity/common"
	"chainup.com/go-sdk/entity/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeout = 120 * time.Second

func Request(key utils.DataCrypto, data string, result interface{}, uri string, config common.WaasConfig) (interface{}, error) {
	if config.AppId == "" {
		fmt.Println("appId is empty")
		return "", errors.New("appId is empty")
	}
	url := config.Domain + uri + "?app_id=" + config.AppId + "&data=" + data
	resp, err := http.Post(url, "application/json;charset=UTF-8", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resultInfo, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, err
	}
	resultCode := &bean.Result{}
	err = json.Unmarshal(resultInfo, &resultCode)
	if err != nil {
		return nil, err
	}
	if !resultCode.IsSuccess() {
		return nil, errors.New(resultCode.Msg)
	}
	fmt.Println(key.Decode(resultCode.Data))
	err = json.Unmarshal([]byte(key.Decode(resultCode.Data)), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
