package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
)

func Sign(data string, privateKeyBytes []byte) ([]byte, error) {
	h := sha256.New()
	h.Write([]byte(data))

	hashed := h.Sum(nil)

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}

	return signature, nil
}
