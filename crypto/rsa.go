package crypto

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

var (
	PubkeyNotSet = errors.New("gorsa: pubkey not set")
	PrikeyNotSet = errors.New("gorsa: prikey not set")
)

var (
	ErrDataLen         = errors.New("gorsa: data length error")
	ErrDataToLarge     = errors.New("gorsa: message too long for RSA public key size")
	ErrDataBroken      = errors.New("gorsa: data broken, first byte is not zero")
	ErrKeyPairDismatch = errors.New("gorsa: data is not encrypted by the private key")
)

type Rsa struct {
	pubkey      *rsa.PublicKey
	prikey      *rsa.PrivateKey
	pkcsVersion int
}

var (
	priType = map[int]string{1: "RSA PRIVATE KEY", 8: "PRIVATE KEY"}
	pubType = map[int]string{1: "RSA PUBLIC KEY", 8: "PUBLIC KEY"}
)

func GenRsaKey(pkcsVersion, bits int) (priKey []byte, pubKey []byte, err error) {
	// check version
	if pkcsVersion != 1 && pkcsVersion != 8 {
		return nil, nil, errors.New("pkcsVersion err, must be 1 or 8")
	}
	// generate privateKey
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	var derStream []byte
	if pkcsVersion == 1 {
		derStream = x509.MarshalPKCS1PrivateKey(privateKey)
	} else {
		derStream, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return nil, nil, err
		}
	}
	block := &pem.Block{
		Type:  priType[pkcsVersion],
		Bytes: derStream,
	}
	priKey = pem.EncodeToMemory(block)
	// generate publicKey
	publicKey := &privateKey.PublicKey
	var derPkix []byte
	if pkcsVersion == 1 {
		derPkix = x509.MarshalPKCS1PublicKey(publicKey)
	} else {
		derPkix, err = x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			return nil, nil, err
		}
	}
	block = &pem.Block{
		Type:  pubType[pkcsVersion],
		Bytes: derPkix,
	}
	pubKey = pem.EncodeToMemory(block)
	return priKey, pubKey, nil
}

func (r *Rsa) ToPkcs8() (string, error) {
	if r.pkcsVersion == 8 {
		return r.GetPrivateKey()
	}

	derStream, err := x509.MarshalPKCS8PrivateKey(r.prikey)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	priKey := pem.EncodeToMemory(block)
	return string(priKey), nil
}

func (r *Rsa) SetPublicKey(pub []byte) (err error) {
	if !strings.Contains(string(pub), "-----") {
		pub = []byte("-----BEGIN PUBLIC KEY-----\n" + string(pub) + "\n-----END PUBLIC KEY-----")
	}
	block, _ := pem.Decode(pub)
	if block == nil {
		return errors.New("failed to SetPublicKey, err in pem.Decode")
	}
	var PKCS1Err string
	pubPKCS1, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err == nil {
		r.pkcsVersion = 1
		r.pubkey = pubPKCS1
		return nil
	} else {
		PKCS1Err = err.Error()
	}
	pubPKCS8, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to SetPublicKey, ParsePKCS1PublicKey err: %v, ParsePKIXPublicKey err: %v", PKCS1Err, err))
	}
	r.pkcsVersion = 8
	r.pubkey = pubPKCS8.(*rsa.PublicKey)
	return nil
}

func (r *Rsa) SetPrivateKey(pri []byte) (err error) {
	if !strings.Contains(string(pri), "-----") {
		pri = []byte("-----BEGIN PRIVATE KEY-----\n" + string(pri) + "\n-----END PRIVATE KEY-----")
	}
	block, _ := pem.Decode(pri)
	if block == nil {
		return errors.New("failed to SetPrivateKey, err in pem.Decode")
	}
	var PKCS1Err string
	priPKCS1, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		r.pkcsVersion = 1
		r.prikey = priPKCS1
		return nil
	} else {
		PKCS1Err = err.Error()
	}
	priPKCS8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to SetPrivateKey, ParsePKCS1PrivateKey err: %v, ParsePKCS8PrivateKey err: %v", PKCS1Err, err))
	}

	r.pkcsVersion = 8
	r.prikey = priPKCS8.(*rsa.PrivateKey)
	return nil
}

func (r *Rsa) GetPublicKey() (string, error) {
	publicKey := &r.prikey.PublicKey
	var derPkix []byte
	var err error
	if r.pkcsVersion == 1 {
		derPkix = x509.MarshalPKCS1PublicKey(publicKey)
	} else {
		derPkix, err = x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			return "", err
		}
	}
	block := &pem.Block{
		Type:  pubType[r.pkcsVersion],
		Bytes: derPkix,
	}
	pubKey := pem.EncodeToMemory(block)
	return string(pubKey), nil
}

func (r *Rsa) GetPrivateKey() (string, error) {
	var derStream []byte
	var err error

	if r.pkcsVersion == 1 {
		derStream = x509.MarshalPKCS1PrivateKey(r.prikey)
	} else {
		derStream, err = x509.MarshalPKCS8PrivateKey(r.prikey)
		if err != nil {
			return "", err
		}
	}

	block := &pem.Block{
		Type:  priType[r.pkcsVersion],
		Bytes: derStream,
	}
	priKey := pem.EncodeToMemory(block)
	return string(priKey), nil
}

func (r *Rsa) rsaPriKeyEncrypt(src []byte) ([]byte, error) {
	if r.prikey == nil {
		return nil, PrikeyNotSet
	}
	partLen := r.prikey.N.BitLen()/8 - 11
	chunks := r.splitChunks(src, partLen)
	buffer := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		decrypted, err := rsa.SignPKCS1v15(nil, r.prikey, crypto.Hash(0), chunk)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to RsaPriKeyEncrypt, err in priKeyEncrypt, err: %v", err))
		}
		buffer.Write(decrypted)
	}
	return buffer.Bytes(), nil
}

func (r *Rsa) rsaPubKeyDecrypt(ciphertext []byte) ([]byte, error) {
	if r.pubkey == nil {
		return nil, PubkeyNotSet
	}
	partLen := r.pubkey.N.BitLen() / 8
	chunks := r.splitChunks(ciphertext, partLen)
	buffer := bytes.NewBuffer(nil)
	for _, chunk := range chunks {
		encrypted, err := r.pubKeyDecrypt(r.pubkey, chunk)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to RsaPubKeyDecrypt, err in pubKeyDecrypt, err: %v", err))
		}
		buffer.Write(encrypted)
	}
	return buffer.Bytes(), nil
}

func (r *Rsa) PriEncryptRetBase64Coding(data []byte, base64Coding *base64.Encoding) (string, error) {
	encryptd, err := r.rsaPriKeyEncrypt(data)
	if err != nil {
		return "", err
	}
	return base64Coding.EncodeToString(encryptd), nil
}

func (r *Rsa) PubDecryptRetBase64Coding(data string, base64Coding *base64.Encoding) ([]byte, error) {
	baseDecode, err := base64Coding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	decryptd, err := r.rsaPubKeyDecrypt(baseDecode)
	if err != nil {
		return nil, err
	}
	return decryptd, nil
}

func (r *Rsa) splitChunks(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}

// modify from utils/rsa
func (r *Rsa) pubKeyDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if k != len(data) {
		return nil, ErrDataLen
	}
	m := new(big.Int).SetBytes(data)
	if m.Cmp(pub.N) > 0 {
		return nil, ErrDataToLarge
	}
	m.Exp(m, big.NewInt(int64(pub.E)), pub.N)
	d := leftPad(m.Bytes(), k)
	if d[0] != 0 {
		return nil, ErrDataBroken
	}
	if d[1] != 1 {
		return nil, ErrKeyPairDismatch
	}
	var i = 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return nil, nil
	}
	return d[i:], nil
}

// copy from utils/rsa
// leftPad returns a new slice of length size. The contents of input are right
// aligned in the new slice.
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}
