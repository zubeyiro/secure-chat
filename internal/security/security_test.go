package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneration(t *testing.T) {
	defer shouldntPanic(t)

	GenerateKeyPair()
}

func TestEncryptSuccess(t *testing.T) {
	key := GenerateKeyPair()
	cipher, err := Encrypt([]byte("input"), &key.PublicKey)

	assert.Equal(t, err, nil)
	assert.Greater(t, len(cipher), 0)
}

func TestEncryptFail(t *testing.T) {
	defer shouldPanic(t)
	Encrypt([]byte("input"), nil)
}

func TestDecryptSuccess(t *testing.T) {
	key := GenerateKeyPair()
	cipher, err := Encrypt([]byte("input"), &key.PublicKey)

	assert.Nil(t, err)
	assert.Greater(t, len(cipher), 0)

	clearText, err := Decrypt(cipher, key)

	assert.Nil(t, err)
	assert.Equal(t, clearText, "input")
}

func TestDecryptFail(t *testing.T) {
	defer shouldPanic(t)
	Decrypt("input", nil)
}

func TestPublicKeyBase64ExportSuccess(t *testing.T) {
	key := GenerateKeyPair()
	str := ExportPublicKeyBase64(&key.PublicKey)

	assert.Greater(t, len(str), 0)
}

func TestPublicKeyBase64ExportFail(t *testing.T) {
	defer shouldPanic(t)
	ExportPublicKeyBase64(nil)
}

func TestParsePublicKeyFromBase64Success(t *testing.T) {
	key := GenerateKeyPair()
	str := ExportPublicKeyBase64(&key.PublicKey)
	parsedKey := ParsePublicKeyFromBase64(str)

	assert.Equal(t, key.PublicKey.Equal(parsedKey), true)
}

func TestParsePublicKeyFromBase64Fail(t *testing.T) {
	defer shouldPanic(t)
	ParsePublicKeyFromBase64("")
}

func shouldntPanic(t *testing.T) {
	r := recover()

	assert.Nil(t, r)
}

func shouldPanic(t *testing.T) {
	r := recover()

	assert.NotNil(t, r)
}
