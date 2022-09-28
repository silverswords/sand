package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type config struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
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

type weChat struct {
	client *http.Client
	token  *TokenResp
}

const (
	code2Session = iota
	getAccessToken
	getUnlimited
)

var url = []string{
	"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
	"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
	"https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN",
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
