package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type weChat struct {
	client *http.Client
}

type config struct {
	AppID     string
	AppSecret string
}

type LoginResponse struct {
	OpenID     string
	SessionKey string
	UnionID    string
	ErrCode    int
	ErrMsg     string
}

func CreateWeChatService() WeChat {
	return &weChat{client: http.DefaultClient}
}

func getConfig() *config {
	c := &config{}
	data, err := ioutil.ReadFile("./config/app.json")
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
	url := fmt.Sprintf(`"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s
	&grant_type=authorization_code"`, config.AppID, config.AppSecret, code)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loginResp *LoginResponse
	if err := json.Unmarshal(buf, loginResp); err != nil {
		return nil, err
	}

	if loginResp.ErrCode != 0 {
		return loginResp, errors.New(loginResp.ErrMsg)
	}

	return loginResp, nil
}
