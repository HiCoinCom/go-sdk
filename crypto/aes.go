package crypto


import (
"crypto/aes"
"encoding/base64"
"fmt"
)

type AesSecurity struct {
	key []byte
}

func (this *AesSecurity) SetAesKey(key []byte) error {
	genKey := make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	this.key = genKey
	return nil
}


// =================== ECB ======================
func (this *AesSecurity) EncryptECB(origData []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(this.key)
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func (this *AesSecurity) DecryptECB(encrypted []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(this.key)
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func (this *AesSecurity) EncryptECB2Base64(origData string) (string) {
	en := this.EncryptECB([]byte(origData))
	return base64.StdEncoding.EncodeToString(en)
}

func (this *AesSecurity) DecryptBase642ECB(origData string) (string, error) {
	en, err := base64.StdEncoding.DecodeString(origData)
	if err != nil {
		fmt.Println("base64 decode origin data error:", origData, err)
		return "", err
	}
	return string( this.DecryptECB(en)), nil
}
