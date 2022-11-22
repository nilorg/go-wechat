package client

import (
	"errors"

	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

var (
	// AppID 应用Key
	AppID = ""
	// AppSecret 秘密
	AppSecret = ""
)

// Parameter 参数
type Parameter map[string]string

// Execute 执行
func Execute(httpClient *http.Client, url string, param Parameter) (json *simplejson.Json, err error) {
	err = checkConfig()
	if err != nil {
		return
	}
	param["appid"] = AppID
	param["secret"] = AppSecret

	result, err := Get(httpClient, url, param)
	if err != nil {
		return
	}
	json, err = simplejson.NewJson([]byte(result))
	if err != nil {
		return
	}
	if errmsg, ok := json.CheckGet("errcode"); ok {
		bytes, _ := errmsg.Encode()
		err = errors.New(string(bytes))
		return
	}
	return
}

func checkConfig() error {
	if AppID == "" {
		return errors.New("AppID 不能为空")
	}
	if AppSecret == "" {
		return errors.New("AppSecret 不能为空")
	}
	return nil
}
