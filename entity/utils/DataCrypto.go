package utils

import (
	"chainup.com/go-sdk/crypto"
	"fmt"
	"strings"
)

type DataCrypto struct {
	PrivateKey string
	PublicKey  string
}

func NewDataCrypto() *DataCrypto {
	return &DataCrypto{}
}

func NewDataCryptoWithKeys(priv, pub string) *DataCrypto {
	return &DataCrypto{
		PrivateKey: priv,
		PublicKey:  pub,
	}
}

func (d *DataCrypto) Decode(cipher string) string {
	ret, err := crypto.PubDecrypt(strings.Replace(cipher, " ", "", -1), d.PublicKey)
	if err != nil {
		fmt.Print("公钥解密数据失败: ", err)
		return ""
	}
	return ret
}

func (d *DataCrypto) Encode(raw string) string {
	result, err := crypto.PrivEncrypt(raw, d.PrivateKey)
	if err != nil {
		fmt.Print("私钥加密数据失败: ", err)
		return ""
	}
	return result
}
