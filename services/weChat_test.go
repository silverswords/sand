package services

import (
	"testing"
)

func TestGetToken(t *testing.T) {
	weChat := CreateWeChatService()
	token, err := weChat.GetAccessToken()
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error(token)
	}

	t.Error(token)
}

func TestGetPhoneNumber(t *testing.T) {
	weChat := CreateWeChatService()
	phone, err := weChat.GetPhoneNumber("123")
	if err != nil {
		t.Error(err)
	}

	t.Errorf("Got error: %v", phone)
}
