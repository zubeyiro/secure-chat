package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		panic(err)
	}

	return privateKey, &privateKey.PublicKey
}

func Encrypt(message []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	encryptedMessage, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)

	if err != nil {
		return nil, err
	}

	return encryptedMessage, nil
}

func Decrypt(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	decryptedMessage, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, message, nil)

	if err != nil {
		return nil, err
	}

	return decryptedMessage, nil
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
