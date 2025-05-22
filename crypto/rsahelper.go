package crypto

import (
	"encoding/base64"
	"errors"
	"strings"
)

func PrivEncrypt(data string, priv string) (string, error) {
	if data == "" {
		return "", nil
	}

	key := new(Rsa)
	err := key.SetPrivateKey([]byte(priv))
	if err != nil {
		return "", err
	}

	var ret string
	ret, err = key.PriEncryptRetBase64Coding([]byte(data), base64.RawURLEncoding)
	if err != nil {
		return "", err
	}
	return ret, nil
}

func PubDecrypt(data string, pub string) (string, error) {
	if data == "" {
		return "", nil
	}

	key := new(Rsa)
	err := key.SetPublicKey([]byte(pub))
	if err != nil {
		return "", err
	}

	var ret []byte
	if strings.HasSuffix(data, "=") {
		ret, err = key.PubDecryptRetBase64Coding(data, base64.URLEncoding)
	} else {
		ret, err = key.PubDecryptRetBase64Coding(data, base64.RawURLEncoding)
	}
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

/**
 * 去除公私钥首尾与换行
 */
func FormatKey(key string) string {
	key = strings.TrimSpace(key)
	//fmt.Println([]byte(key))
	if strings.HasPrefix(key, "---") && strings.HasSuffix(key, "---") {
		str := strings.Split(key, "\n")
		return strings.Join(str[1:len(str)-1], "")
	}
	return key
}

/**
 * 返回去除首尾与换行的公私钥
 */
func NewRsaKey() (string, string, error) {
	priv, pub, err := GenRsaKey(8, 2048)
	if err != nil {
		return "", "", err
	}

	if priv == nil || len(priv) == 0 {
		return "", "", errors.New("priv data is nil")
	}

	if pub == nil || len(pub) == 0 {
		return "", "", errors.New("pub data is nil")
	}

	return FormatKey(string(priv)), FormatKey(string(pub)), nil
}

func GetPubFromPriv(priv string) (string, error) {
	key := new(Rsa)
	err := key.SetPrivateKey([]byte(priv))
	if err != nil {
		return "", err
	}

	var pub string
	pub, err = key.GetPublicKey()
	if err != nil {
		return "", err
	}
	return FormatKey(pub), nil
}
