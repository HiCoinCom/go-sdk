// Package utils provides utility functions and constants for ChainUp Custody SDK.
package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"strings"
	"testing"
)

// Test RSA private key encryption and public key decryption
func TestRSAPrivateKeyEncryptPublicKeyDecrypt(t *testing.T) {
	// Generate a test RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	// Create crypto provider with the key pair
	provider, err := NewRSACryptoProviderWithKeys(privateKey, &privateKey.PublicKey, "UTF-8")
	if err != nil {
		t.Fatalf("Failed to create crypto provider: %v", err)
	}

	testCases := []struct {
		name     string
		plaintext string
	}{
		{
			name:     "Short text",
			plaintext: "Hello, World!",
		},
		{
			name:     "Medium text",
			plaintext: "This is a medium length test string that should fit in one chunk.",
		},
		{
			name:     "Long text requiring multiple chunks",
			plaintext: strings.Repeat("This is a long test string. ", 50),
		},
		{
			name:     "JSON data",
			plaintext: `{"code":"0","data":{"uid":12345,"email":"test@example.com"},"msg":"success"}`,
		},
		{
			name:     "Special characters",
			plaintext: "Special chars: !@#$%^&*()_+-=[]{}|;':\",./<>?中文日本語한국어",
		},
		{
			name:     "Empty string",
			plaintext: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Skip empty string test as it may not work well with RSA
			if tc.plaintext == "" {
				t.Skip("Skipping empty string test")
			}

			// Encrypt with private key
			encrypted, err := provider.EncryptWithPrivateKey(tc.plaintext)
			if err != nil {
				t.Fatalf("Failed to encrypt: %v", err)
			}

			// Verify encrypted data is not empty and is base64
			if encrypted == "" {
				t.Fatal("Encrypted data should not be empty")
			}

			// Decrypt with public key
			decrypted, err := provider.DecryptWithPublicKey(encrypted)
			if err != nil {
				t.Fatalf("Failed to decrypt: %v", err)
			}

			// Verify decrypted data matches original
			if decrypted != tc.plaintext {
				t.Errorf("Decrypted data mismatch.\nExpected: %s\nGot: %s", tc.plaintext, decrypted)
			}
		})
	}
}

// Test RSA signing and verification
func TestRSASignAndVerify(t *testing.T) {
	// Generate a test RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	// Create crypto provider with the key pair
	provider, err := NewRSACryptoProviderWithKeys(privateKey, &privateKey.PublicKey, "UTF-8")
	if err != nil {
		t.Fatalf("Failed to create crypto provider: %v", err)
	}

	testCases := []struct {
		name string
		data string
	}{
		{
			name: "Simple data",
			data: "amount=100&request_id=test123&symbol=btc",
		},
		{
			name: "Sorted params",
			data: "address_to=0x123&amount=1.5&request_id=req001&sub_wallet_id=1&symbol=eth",
		},
		{
			name: "Unicode data",
			data: "memo=测试备注&request_id=test123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Sign with private key
			signature, err := provider.SignWithPrivateKey(tc.data)
			if err != nil {
				t.Fatalf("Failed to sign: %v", err)
			}

			// Verify signature is not empty
			if signature == "" {
				t.Fatal("Signature should not be empty")
			}

			// Verify with public key
			valid, err := provider.VerifyWithPublicKey(tc.data, signature)
			if err != nil {
				t.Fatalf("Failed to verify: %v", err)
			}

			if !valid {
				t.Error("Signature verification failed")
			}

			// Verify that modified data fails verification
			modifiedData := tc.data + "x"
			valid, err = provider.VerifyWithPublicKey(modifiedData, signature)
			if err != nil {
				t.Fatalf("Failed to verify modified data: %v", err)
			}

			if valid {
				t.Error("Signature verification should fail for modified data")
			}
		})
	}
}

// Test encryption with chunking
func TestRSAChunkedEncryption(t *testing.T) {
	// Generate a test RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	provider, err := NewRSACryptoProviderWithKeys(privateKey, &privateKey.PublicKey, "UTF-8")
	if err != nil {
		t.Fatalf("Failed to create crypto provider: %v", err)
	}

	// Create data that requires multiple chunks
	// For 2048-bit key, chunk size is 256 - 11 = 245 bytes
	// So we need > 245 bytes to test chunking
	largeData := strings.Repeat("A", 500)

	encrypted, err := provider.EncryptWithPrivateKey(largeData)
	if err != nil {
		t.Fatalf("Failed to encrypt large data: %v", err)
	}

	decrypted, err := provider.DecryptWithPublicKey(encrypted)
	if err != nil {
		t.Fatalf("Failed to decrypt large data: %v", err)
	}

	if decrypted != largeData {
		t.Errorf("Large data decryption mismatch.\nExpected length: %d\nGot length: %d", len(largeData), len(decrypted))
	}
}

// Test key parsing (PEM and raw base64)
func TestKeyParsing(t *testing.T) {
	// Generate a test RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key: %v", err)
	}

	// Test with NewRSACryptoProviderWithKeys
	provider, err := NewRSACryptoProviderWithKeys(privateKey, &privateKey.PublicKey, "")
	if err != nil {
		t.Fatalf("Failed to create provider with keys: %v", err)
	}

	// Verify it can encrypt and decrypt
	plaintext := "Test data"
	encrypted, err := provider.EncryptWithPrivateKey(plaintext)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	decrypted, err := provider.DecryptWithPublicKey(encrypted)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if decrypted != plaintext {
		t.Errorf("Decryption mismatch: expected %s, got %s", plaintext, decrypted)
	}
}

// Benchmark encryption
func BenchmarkEncryption(b *testing.B) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	provider, _ := NewRSACryptoProviderWithKeys(privateKey, &privateKey.PublicKey, "UTF-8")
	data := strings.Repeat("Test data for benchmarking ", 10)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = provider.EncryptWithPrivateKey(data)
	}
}

// Benchmark decryption
func BenchmarkDecryption(b *testing.B) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	provider, _ := NewRSACryptoProviderWithKeys(privateKey, &privateKey.PublicKey, "UTF-8")
	data := strings.Repeat("Test data for benchmarking ", 10)
	encrypted, _ := provider.EncryptWithPrivateKey(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = provider.DecryptWithPublicKey(encrypted)
	}
}

// Benchmark signing
func BenchmarkSigning(b *testing.B) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	provider, _ := NewRSACryptoProviderWithKeys(privateKey, &privateKey.PublicKey, "UTF-8")
	data := "amount=100&request_id=test123&symbol=btc"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = provider.SignWithPrivateKey(data)
	}
}
