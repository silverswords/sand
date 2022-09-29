package services

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/silverswords/sand/utils"
)

type sign struct {
}

func CreateSignService() Sign {
	return &sign{}
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
	nonceStr, err := utils.RandString(20)
	if err != nil {
		return nil, err
	}

	timeStamp := time.Now().Unix()
	_package := "prepay_id=" + prepayID
	message := fmt.Sprintf("%s\n%d\n%s\n%s\n", appID, timeStamp, nonceStr, _package)
	paySign, err := utils.Sign(message, []byte{})
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
