package services

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/silverswords/sand/utils"
)

const (
	privateFilePath = "../config/apiclient_key.pem"
)

type sign struct {
	privateKey *rsa.PrivateKey
}

func CreateSignService() (Sign, error) {
	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}

	sign := &sign{privateKey: privateKey}
	return sign, nil
}

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

type PayInfo struct {
	AppID     string `json:"appid"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}

func (s *sign) GetSignedInfo(prepayID string, appID string) (*PayInfo, error) {
	var (
		payInfo *PayInfo
	)
	nonceStr, err := utils.GenerateNonce()
	if err != nil {
		return nil, err
	}

	timeStamp := time.Now().Unix()
	_package := "prepay_id=" + prepayID
	message := fmt.Sprintf("%s\n%d\n%s\n%s\n", appID, timeStamp, nonceStr, _package)
	paySign, err := utils.SignSHA256WithRSA(message, s.privateKey)
	if err != nil {
		return nil, err
	}

	payInfo.AppID = appID
	payInfo.TimeStamp = strconv.FormatInt(timeStamp, 10)
	payInfo.NonceStr = nonceStr
	payInfo.Package = "prepay_id=" + prepayID
	payInfo.SignType = "RSA"
	payInfo.PaySign = base64.StdEncoding.EncodeToString(paySign)

	return payInfo, nil
}
