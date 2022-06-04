package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

func GenerateKeyPair() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		panic(err)
	}

	return privateKey
}

func Encrypt(message []byte, publicKey *rsa.PublicKey) (string, error) {
	encryptedMessage, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedMessage), nil
}

func Decrypt(message string, privateKey *rsa.PrivateKey) (string, error) {
	bytes, _ := base64.StdEncoding.DecodeString(message)

	decryptedMessage, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, bytes, nil)

	if err != nil {
		return "", err
	}

	return string(decryptedMessage), nil
}

func ExportPublicKeyBase64(key *rsa.PublicKey) string {
	return base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(key))
}

func ParsePublicKeyFromBase64(key string) *rsa.PublicKey {
	b, err := base64.StdEncoding.DecodeString(key)

	if err != nil {
		panic(err)
	}

	publicKey, err := x509.ParsePKCS1PublicKey(b)

	if err != nil {
		panic(err)
	}

	return publicKey
}
