package wechat

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	err := NewError([]byte(`{ "access_token":"ACCESS_TOKEN",
		"expires_in":7200,
		"refresh_token":"REFRESH_TOKEN",
		"openid":"OPENID",
		"scope":"SCOPE" }`))
	fmt.Println(err == nil)
}
