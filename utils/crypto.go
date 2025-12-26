// Package utils provides utility functions and constants for ChainUp Custody SDK.
package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// CryptoProvider defines the interface for cryptographic operations.
type CryptoProvider interface {
	// EncryptWithPrivateKey encrypts data with the private key.
	EncryptWithPrivateKey(data string) (string, error)

	// DecryptWithPublicKey decrypts data with the public key.
	DecryptWithPublicKey(data string) (string, error)

	// SignWithPrivateKey signs data using RSA-SHA256 with the private key.
	// This is used for MPC withdraw and web3 transaction signatures.
	SignWithPrivateKey(data string) (string, error)

	// VerifyWithPublicKey verifies signature using RSA-SHA256 with the public key.
	VerifyWithPublicKey(data string, signature string) (bool, error)
}

// RSACryptoProvider implements RSA encryption/decryption operations.
type RSACryptoProvider struct {
	privateKey     *rsa.PrivateKey
	publicKey      *rsa.PublicKey
	signPrivateKey *rsa.PrivateKey // Optional: separate key for signing
	charset        string
}

// NewRSACryptoProvider creates a new RSA crypto provider.
// privateKeyPEM: PEM-encoded RSA private key (optional if only decrypting)
// publicKeyPEM: PEM-encoded RSA public key (optional if only encrypting)
// charset: Character encoding for data (defaults to UTF-8)
func NewRSACryptoProvider(privateKeyPEM, publicKeyPEM, charset string) (*RSACryptoProvider, error) {
	provider := &RSACryptoProvider{
		charset: charset,
	}

	if charset == "" {
		provider.charset = DefaultCharset
	}

	// Parse private key if provided
	if privateKeyPEM != "" {
		privateKey, err := ParsePrivateKey(privateKeyPEM)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}
		provider.privateKey = privateKey
	}

	// Parse public key if provided
	if publicKeyPEM != "" {
		publicKey, err := ParsePublicKey(publicKeyPEM)
		if err != nil {
			return nil, fmt.Errorf("failed to parse public key: %w", err)
		}
		provider.publicKey = publicKey
	}

	return provider, nil
}

// NewRSACryptoProviderWithSignKey creates a new RSA crypto provider with a separate signing key.
// privateKeyPEM: PEM-encoded RSA private key for encryption
// publicKeyPEM: PEM-encoded RSA public key for decryption
// signPrivateKeyPEM: PEM-encoded RSA private key for signing (optional, uses privateKey if empty)
// charset: Character encoding for data (defaults to UTF-8)
func NewRSACryptoProviderWithSignKey(privateKeyPEM, publicKeyPEM, signPrivateKeyPEM, charset string) (*RSACryptoProvider, error) {
	provider, err := NewRSACryptoProvider(privateKeyPEM, publicKeyPEM, charset)
	if err != nil {
		return nil, err
	}

	// Parse sign private key if provided
	if signPrivateKeyPEM != "" {
		signKey, err := ParsePrivateKey(signPrivateKeyPEM)
		if err != nil {
			return nil, fmt.Errorf("failed to parse sign private key: %w", err)
		}
		provider.signPrivateKey = signKey
	}

	return provider, nil
}

// NewRSACryptoProviderWithKeys creates a new RSA crypto provider with pre-parsed keys.
// privateKey: Parsed RSA private key (optional if only decrypting)
// publicKey: Parsed RSA public key (optional if only encrypting)
// charset: Character encoding for data (defaults to UTF-8)
func NewRSACryptoProviderWithKeys(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, charset string) (*RSACryptoProvider, error) {
	provider := &RSACryptoProvider{
		privateKey: privateKey,
		publicKey:  publicKey,
		charset:    charset,
	}

	if charset == "" {
		provider.charset = DefaultCharset
	}

	return provider, nil
}
// EncryptWithPrivateKey encrypts data with the private key.
// Uses rsa.SignPKCS1v15 with crypto.Hash(0) to perform raw private key encryption.
// The encrypted data can be decrypted with the corresponding public key.
func (r *RSACryptoProvider) EncryptWithPrivateKey(data string) (string, error) {
	if r.privateKey == nil {
		return "", errors.New("private key not set")
	}

	dataBytes := []byte(data)
	// PKCS1v15 padding requires 11 bytes overhead
	chunkSize := r.privateKey.N.BitLen()/8 - 11
	chunks := splitChunks(dataBytes, chunkSize)

	var encrypted []byte
	for _, chunk := range chunks {
		// Use SignPKCS1v15 with Hash(0) to encrypt without hashing
		encryptedChunk, err := rsa.SignPKCS1v15(nil, r.privateKey, crypto.Hash(0), chunk)
		if err != nil {
			return "", fmt.Errorf("failed to encrypt chunk: %w", err)
		}
		encrypted = append(encrypted, encryptedChunk...)
	}

	// Use URL-safe Base64 encoding (no padding)
	return base64.RawURLEncoding.EncodeToString(encrypted), nil
}

// splitChunks splits a byte slice into chunks of specified size
func splitChunks(buf []byte, lim int) [][]byte {
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

// DecryptWithPublicKey decrypts data with the public key.
// This matches the Java SDK's public key decryption for response verification.
// Supports both standard and URL-safe Base64 encoding.
func (r *RSACryptoProvider) DecryptWithPublicKey(data string) (string, error) {
	if r.publicKey == nil {
		return "", errors.New("public key not set")
	}

	// Try URL-safe Base64 first (what the server uses), then standard Base64
	data = convertURLSafeBase64ToStandard(data)
	encrypted, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// Try with padding
		data = addBase64Padding(data)
		encrypted, err = base64.StdEncoding.DecodeString(data)
		if err != nil {
			return "", fmt.Errorf("failed to decode base64: %w", err)
		}
	}

	keySize := r.publicKey.Size()

	var decrypted []byte
	for i := 0; i < len(encrypted); i += keySize {
		end := i + keySize
		if end > len(encrypted) {
			end = len(encrypted)
		}

		chunk := encrypted[i:end]
		// RSA public key decryption (decrypt data encrypted with private key)
		decryptedChunk, err := rsaPublicKeyDecrypt(r.publicKey, chunk)
		if err != nil {
			return "", fmt.Errorf("failed to decrypt chunk: %w", err)
		}
		decrypted = append(decrypted, decryptedChunk...)
	}

	return string(decrypted), nil
}

// convertURLSafeBase64ToStandard converts URL-safe Base64 to standard Base64
func convertURLSafeBase64ToStandard(data string) string {
	// Replace URL-safe characters with standard ones
	data = strings.ReplaceAll(data, "-", "+")
	data = strings.ReplaceAll(data, "_", "/")
	return data
}

// addBase64Padding adds padding to Base64 string if needed
func addBase64Padding(data string) string {
	switch len(data) % 4 {
	case 2:
		data += "=="
	case 3:
		data += "="
	}
	return data
}

// rsaPublicKeyDecrypt decrypts data using RSA public key (for data encrypted with private key)
func rsaPublicKeyDecrypt(pub *rsa.PublicKey, ciphertext []byte) ([]byte, error) {
	// RSA public key decryption: m = c^e mod n
	c := new(big.Int).SetBytes(ciphertext)
	e := big.NewInt(int64(pub.E))
	m := new(big.Int).Exp(c, e, pub.N)

	// Get the decrypted bytes
	decrypted := m.Bytes()

	// Handle PKCS1 v1.5 padding
	if len(decrypted) > 0 && decrypted[0] == 0x01 {
		// Find the 0x00 separator
		for i := 1; i < len(decrypted); i++ {
			if decrypted[i] == 0x00 {
				return decrypted[i+1:], nil
			}
		}
	}

	// If no padding found, try to find 0x00 separator anyway
	for i := 0; i < len(decrypted); i++ {
		if decrypted[i] == 0x00 && i > 0 {
			return decrypted[i+1:], nil
		}
	}

	return decrypted, nil
}

// SetSignPrivateKey sets a separate private key for signing.
// If signPrivateKey is set, SignWithPrivateKey will use it instead of privateKey.
func (r *RSACryptoProvider) SetSignPrivateKey(key *rsa.PrivateKey) {
	r.signPrivateKey = key
}

// GetSigningKey returns the key used for signing.
// Returns signPrivateKey if set, otherwise returns privateKey.
func (r *RSACryptoProvider) GetSigningKey() *rsa.PrivateKey {
	if r.signPrivateKey != nil {
		return r.signPrivateKey
	}
	return r.privateKey
}

// SignWithPrivateKey signs data using RSA-SHA256.
// This performs: SHA256(data) -> RSA sign -> Base64 encode.
// Used for MPC withdraw and web3 transaction signatures.
// If signPrivateKey is set, it will be used for signing; otherwise privateKey is used.
func (r *RSACryptoProvider) SignWithPrivateKey(data string) (string, error) {
	fmt.Println("==========> ", data)
	// Use signPrivateKey if set, otherwise use privateKey
	signingKey := r.GetSigningKey()
	if signingKey == nil {
		return "", errors.New("no signing key available (neither signPrivateKey nor privateKey is set)")
	}

	// Step 1: SHA256 hash the data
	hash := sha256.New()
	hash.Write([]byte(data))

	// Step 2: RSA sign with PKCS1v15
	signature, err := rsa.SignPKCS1v15(rand.Reader, signingKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return "", fmt.Errorf("failed to sign: %w", err)
	}

	// Step 3: Base64 encode
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifyWithPublicKey verifies a signature using RSA-SHA256.
// This performs: SHA256(data) -> RSA verify with Base64-decoded signature.
func (r *RSACryptoProvider) VerifyWithPublicKey(data string, signature string) (bool, error) {
	if r.publicKey == nil {
		return false, errors.New("public key not set")
	}

	// Decode signature from Base64 (try standard first, then URL-safe)
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		sigBytes, err = base64.URLEncoding.DecodeString(signature)
		if err != nil {
			return false, fmt.Errorf("failed to decode signature: %w", err)
		}
	}

	// Step 1: SHA256 hash the data
	hash := sha256.New()
	hash.Write([]byte(data))

	// Step 2: RSA verify with PKCS1v15
	err = rsa.VerifyPKCS1v15(r.publicKey, crypto.SHA256, hash.Sum(nil), sigBytes)
	if err != nil {
		return false, nil // Verification failed, not an error
	}

	return true, nil
}

// ParsePrivateKey parses an RSA private key.
// Supports both PEM-encoded format and raw base64-encoded format.
// Supports both PKCS1 and PKCS8 formats.
func ParsePrivateKey(keyStr string) (*rsa.PrivateKey, error) {
	var keyBytes []byte

	// Try PEM decode first
	block, _ := pem.Decode([]byte(keyStr))
	if block != nil {
		keyBytes = block.Bytes
	} else {
		// Try raw base64 decode
		var err error
		keyBytes, err = base64.StdEncoding.DecodeString(keyStr)
		if err != nil {
			return nil, errors.New("failed to decode key: not valid PEM or base64")
		}
	}

	// Try PKCS1 format first
	privateKey, err := x509.ParsePKCS1PrivateKey(keyBytes)
	if err == nil {
		return privateKey, nil
	}

	// Try PKCS8 format
	key, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}

	return rsaKey, nil
}

// ParsePublicKey parses an RSA public key.
// Supports both PEM-encoded format and raw base64-encoded format.
// Supports both PKIX and PKCS1 formats.
func ParsePublicKey(keyStr string) (*rsa.PublicKey, error) {
	var keyBytes []byte

	// Try PEM decode first
	block, _ := pem.Decode([]byte(keyStr))
	if block != nil {
		keyBytes = block.Bytes
	} else {
		// Try raw base64 decode
		var err error
		keyBytes, err = base64.StdEncoding.DecodeString(keyStr)
		if err != nil {
			return nil, errors.New("failed to decode key: not valid PEM or base64")
		}
	}

	// Try PKIX format first
	pub, err := x509.ParsePKIXPublicKey(keyBytes)
	if err == nil {
		rsaPub, ok := pub.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("not an RSA public key")
		}
		return rsaPub, nil
	}

	// Try PKCS1 format
	return x509.ParsePKCS1PublicKey(keyBytes)
}
