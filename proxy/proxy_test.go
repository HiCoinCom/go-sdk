package proxy

import (
	"chainup.com/go-sdk/entity/bean/deposit"
	"chainup.com/go-sdk/entity/common"
	"chainup.com/go-sdk/entity/utils"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestPrivEncrypt(t *testing.T) {
	key := utils.DataCrypto{
		PrivateKey: "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDO0ovpG7plFUL2+kLMu/VDgC7+o06kNIjQ72Jd9FYGM+xXeiasUxZUTzyXk9LPrc7UU4ERo/W1XDvu4F+a8yL5EXkJAHqvjy53ODcBz7p8v3nwSkeyKhe8T6OdKydcrBuuPBKFYGqC2Qfd4358FxZZskMwFaCJv/MF/P/2JlbXLCztktqVX1jbkeAgViaeKgB9Ew3TGrTV6R3GS82hozgqeTjcSQXbKCegHlu6AQJj/NPSi2Xcly7Ci9fqmOUNl6sTB6PoemrkUnYl6ov9VWyA8rLme1i3hXhVQvy2653nAAmETcX/Ttjo8s3zUi5oX9bYKkfoQglkucAIB0aFRgXfAgMBAAECggEAe+GjSjqAmEF2/eaDEozW6Nqjy4RX+4U4GiX4/pv21vpM60zQU1HlQxS/D2/Myvapt4ZY/g7+guY/Q+MDgRV3ckaj/99YnUWfKGv0jtI1OrmaOXLUpQQPeY0eTGrnxduVLPTwhxFsWvnSe3bjcUKG8a0UD340za26zHBxhKgMRg8sGjlvO7QBUu2HRJHB1h2216zBhs6C76xpk3jE3uIR4h90YS2QnlU5altr/UTxZ0jfzVB5E1Auaa/4WNlY7n1BF0DqKV2whNgIqeeVa/BPa29sjNtsEhJbKTZ9MsASEV9FfCgjOUVE4r3cFPWO3T4Oijfdk+mmGazUVDb/Sd5AIQKBgQDS4oCyiqSeZfVSkE6c0BQYtRrxklbN/T56JbWTCQHXTinYtjkDDd6F+nmMA7/5/qFCQ/xClALlSqZ2VhsESkPL2BB2t/woaB3StAVJIAHg8tYKoN5vgg/Jc4j4n3UrAcF5OnTX6j7GwIRCO6/SFzddUYL5I6Hx5joIvzY0fJkxEQKBgQD7EZATi0dnE5GfKdTZ0VK5jfGexRXg0b65QNLMSbbtQogilqZA8707hg+40tg1fYfrpKOWKr+XXjmcthlK2UO9mjR2Z2//fkTyALoc8gOg+Sf/i/x4bhHZ38C+UtARY+U/IfHsnmojSIGLP7ujkmLjjNVPuL3lD8ImapiBq03H7wKBgEfitgQmhp8IGmwejO2fcHpLofee7sAhB18T46VPeLUwC8u3sg98dtIs46n6zUImrkmsax023xXSMJQ+Hc+EkT+3U0Vkyivr8d6VRwf6RSmtHZFZ7PkN2NvO6m31zTbvzkfIyXOgge22Sl9ZgUGYcGL7Gi5bGyUeWcIJjCywHHhBAoGAEMt799K5VAvbEeqacneuMPttzlEgxYlCWOIdN802j2iD8sxhErc7UWAGbTatf+aF0R2SU9lIN7f14qXLy1X9UhGesMz6kPzBX+7shEMsOvhv1IhouyWhzBFLs2+Wu5MTIsYCiPGP6AxRnh8FTkvckZw6KLsS+N+0/RzkVsg+y8ECgYA8grOcKwG0uHH8j0XK5AcrtOoTAxsWRuH1tkzGuOOFRaZXaqIwgRvQ025pTe3lnTCgiKc669DPqCRcwYeZbjC1LlVjPYypBd3uOIX+62fpehxCIV9naeEgqvyCguXzvI94pzohHDbQxjhcdvH3eYODsybiGMeOB2XOqCclT/vxLw==",
		PublicKey:  "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnvBDdX1VBnY3dEn3YQu4wZeaKvV8xRlP2ljEwZAPpN9UYHNS+O/hYOZD3HZoQFEhXtImKu4l7C1pDn+lIPzWDRqQRXNHt8crFeGxi8d6OXB6gYmzXz4KzRuuMBB5HmxuaaiIh4w0dZna71z7nI9Y4Gn/LibvVrxpOLVqj5c7SBKyIgXBrGfaOXPDeOwj2X94Lp4GtKyhyDiEJ2asBHNTjpF9Yzlk6pKFDvrijpgdcjgEygNAMi8/bfJae5MWko/7scgXu3NKE2QTl3RVCcHRVKU45r1IKFZBLMkz97NU9ZcIyaBtUDTJGzV1K90sEOfZfKvU210+hC1UB2NNzqCszQIDAQAB ",
	}
	coinInfo := deposit.SyncDepositListArgs{
		MaxId: "0",
	}
	if !coinInfo.Validate() {
		fmt.Println("validate fail")
	}
	marshal, _ := json.Marshal(coinInfo)

	config := common.WaasConfig{
		AppId:  "fc64c905be168391e0158e9bc9bbfb7a",
		Domain: "",
	}

	b := &deposit.SyncDepositListResult{}
	_, err := Request(key, key.Encode(string(marshal)), &b, string(common.SyncDepositList), config)
	if err != nil {
		return
	}
	if b.Code != "0" {
		fmt.Println(b.Msg)
	}
	fmt.Println(b)
}

func TestT(t *testing.T) {
	fromString, _ := decimal.NewFromString("123.123123213")
	fmt.Println(fromString.Exponent())
}
