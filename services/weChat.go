package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"

	"github.com/silverswords/sand/utils"
)

type config struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	MchID     string `json:"mch_id"`
	SerialNo  string `json:"serial_no"`
}

type LoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type PhoneResp struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	PhoneInfo struct {
		PhoneNumber     string `json:"phone_number"`
		PurePhoneNumber string `json:"pure_phone_number"`
		CountryCode     string `json:"country_code"`
		Water           struct {
			Timestamp int    `json:"timestamp"`
			AppID     string `json:"app_id"`
		}
	} `json:"phone_info"`
}

type weChat struct {
	client *http.Client
	token  *TokenResp
}

const (
	code2Session = iota
	getAccessToken
	getUnlimited
	getPhoneNumber
	createOrder
	notify
)

var url = []string{
	"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
	"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
	"https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN",
	"https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s",
	"https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi",
	"https://",
}

func CreateWeChatService() WeChat {
	return &weChat{
		client: http.DefaultClient,
		token:  &TokenResp{},
	}
}

func get(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *weChat) request(method string, url string, jsonData []byte) (*http.Response, error) {
	buf := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	ua := fmt.Sprintf("api/v1(%s) Go/%s", runtime.GOOS, runtime.Version())
	req.Header.Set("User-Agent", ua)

	authorization, err := generateAuthorizationHeader(method, req.URL.RequestURI(), buf.String())
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", authorization)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func generateAuthorizationHeader(method, canonicalURL, signBody string) (string, error) {
	config := getConfig()

	nonce, err := utils.GenerateNonce()
	if err != nil {
		return "", err
	}
	timestamp := time.Now().Unix()
	message := fmt.Sprintf("%s\n%s\n%d\n%s\n%s\n", method, canonicalURL, timestamp, nonce, signBody)
	signature, err := utils.SignSHA256WithRSA(message)
	if err != nil {
		return "", err
	}

	authorization := fmt.Sprintf(
		"%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\"",
		"Authorization: WECHATPAY2-SHA256-RSA2048", config.MchID,
		nonce, timestamp, config.SerialNo, signature,
	)
	return authorization, nil
}

func getConfig() *config {
	c := &config{}
	data, err := ioutil.ReadFile("../config/app.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, c); err != nil {
		panic(err)
	}

	return c
}

func (s *weChat) Login(code string) (*LoginResponse, error) {
	config := getConfig()
	url := fmt.Sprintf(url[code2Session], config.AppID, config.AppSecret, code)

	buf, err := get(s.client, url)
	if err != nil {
		return nil, err
	}

	var loginResp LoginResponse
	if err := json.Unmarshal(buf, &loginResp); err != nil {
		return nil, err
	}

	if loginResp.ErrCode != 0 {
		return nil, errors.New(loginResp.ErrMsg)
	}

	return &loginResp, nil
}

func (s *weChat) GetAccessToken() (string, error) {
	if s.token.ExpiresIn < 600 {
		config := getConfig()
		url := fmt.Sprintf(url[getAccessToken], config.AppID, config.AppSecret)

		buf, err := get(s.client, url)
		if err != nil {
			return "", err
		}

		var tokenResp TokenResp
		if err := json.Unmarshal(buf, &tokenResp); err != nil {
			return "", err
		}

		if tokenResp.ErrCode != 0 {
			return "", errors.New(tokenResp.ErrMsg)
		}

		s.token.AccessToken = tokenResp.AccessToken
	}

	return s.token.AccessToken, nil
}

func (s *weChat) GetPhoneNumber(code string) (*PhoneResp, error) {
	access_token, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}

	var (
		postData struct {
			Code string `json:"code"`
		}
		phoneResp PhoneResp
	)
	postData.Code = code

	jsonData, jsonErr := json.Marshal(postData)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := fmt.Sprintf(url[getPhoneNumber], access_token)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf, &phoneResp); err != nil {
		return nil, err
	}

	return &phoneResp, nil
}

func (s *weChat) GetPrepayInfo(description string, outTradeNo string, total int, openID string) (string, string, error) {
	config := getConfig()

	var (
		postData struct {
			AppID       string `json:"appid"`
			MchID       string `json:"mchid"`
			Description string `json:"description"`
			OutTradeNo  string `json:"out_trade_no"`
			NotifyUrl   string `json:"notify_url"`
			Amount      struct {
				Total int `json:"total"`
			} `json:"amount"`
			Payer struct {
				OpenID string `json:"openid"`
			} `json:"payer"`
		}
		prepayResp struct {
			PrepayID string `json:"prepay_id"`
		}
	)

	postData.AppID = config.AppID
	postData.MchID = config.MchID
	postData.Description = description
	postData.OutTradeNo = outTradeNo
	postData.Amount.Total = total
	postData.Payer.OpenID = openID
	postData.NotifyUrl = url[notify]

	jsonData, jsonErr := json.Marshal(postData)
	if jsonErr != nil {
		return "", "", jsonErr
	}
	resp, err := s.request(http.MethodPost, url[createOrder], jsonData)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	if err := json.Unmarshal(buf, &prepayResp); err != nil {
		return "", "", err
	}

	return prepayResp.PrepayID, config.AppID, nil
}
