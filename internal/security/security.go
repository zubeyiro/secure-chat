package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
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
