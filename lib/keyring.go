package lib

import "github.com/zalando/go-keyring"

func SetToken(token string) error {
	return keyring.Set(AppName, UserName, token)
}

func ClearToken() error {
	return keyring.Delete(AppName, UserName)
}
