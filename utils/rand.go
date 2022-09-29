package utils

import (
	"crypto/rand"
)

func RandString(len int) (string, error) {
	bytes := make([]byte, len)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
