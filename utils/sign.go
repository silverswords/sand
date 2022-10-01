package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	privateFilePath = "config/apiclient_key.pem"
)

func getPrivateKey() (*rsa.PrivateKey, error) {
	file, err := os.Open(privateFilePath)
	if err != nil {
		return nil, err
	}

	pemBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("private key error")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key err:%s", err.Error())
	}

	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("privateFile is not rsa private key")
	}

	return privateKey, nil
}

func SignSHA256WithRSA(source string) ([]byte, error) {
	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}

	if privateKey == nil {
		return nil, fmt.Errorf("private key should not be nil")
	}

	h := sha256.New()
	_, err = h.Write([]byte(source))
	if err != nil {
		return nil, err
	}

	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}

	return signature, nil
}
