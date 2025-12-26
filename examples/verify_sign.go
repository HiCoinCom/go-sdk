//go:build ignore

package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
)

func main() {
	// RSA 私钥 (从 mpc_example.go)
	privateKeyStr := ``

	// 解析私钥获取公钥
	keyBytes, _ := base64.StdEncoding.DecodeString(privateKeyStr)
	privateKey, _ := x509.ParsePKCS8PrivateKey(keyBytes)
	rsaKey := privateKey.(*rsa.PrivateKey)
	pubKey := &rsaKey.PublicKey

	// 要验证的签名
	signature := "PmkfOQBP4rjLybnBi4prQa/pRhdVNDL1o65LqQZOJe44RlDY5WZ3ai3RPu7tA+tmwQmklVQLsiwLVVR6q+V9e2feUJKUrD+qTsXnVF+XPc/2misHxiH1gPJ2SHJBD3M4kac0bHgo27GmwribX6IKIeXWiAK1JDmUjsZG5Ta8KUwE7we0WcLGHt1KDFmjSzAXnbgb3LJJEjeTfQxO/4sN26A+on2JoBEV8d9Aby+mh+WMypT28V2sGJfogFbNiTd/Xdx1RZlzI168hfMH0JqZ3AHMmRYqj7Hr4s7uRiDD87bA9Cld+P/8xnMYnWbFEpi2ItHDM6em5Zwbf1jJ3qE9qA=="

	// 签名只包含这些字段 (根据 withdraw.go)
	params := map[string]string{
		"request_id":    "12345678906",
		"sub_wallet_id": "1000537",
		"symbol":        "Sepolia",
		"address_to":    "0xdcb0D867403adE76e75a4A6bBcE9D53C9d05B981",
		"amount":        "0.001",
	}

	// 按 key 排序并拼接
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, params[k]))
	}
	signData := strings.ToLower(strings.Join(parts, "&"))
	fmt.Println("Sign data:", signData)

	// 解码签名
	signBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		signBytes, err = base64.URLEncoding.DecodeString(signature)
		if err != nil {
			fmt.Println("Failed to decode signature:", err)
			return
		}
	}

	// SHA256 哈希
	hash := sha256.New()
	hash.Write([]byte(signData))

	// 验证签名
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash.Sum(nil), signBytes)
	if err != nil {
		fmt.Println("Signature verification FAILED:", err)
	} else {
		fmt.Println("Signature verification SUCCESS!")
	}
}
