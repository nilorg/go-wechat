package qrcode

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/nilorg/go-wechat/v2/client"
)

// Qrcode 二维码
type Qrcode struct {
	client client.Clienter
}

// NewQrcode ...
func NewQrcode(c client.Clienter) *Qrcode {
	return &Qrcode{
		client: c,
	}
}

// CreateTemp 生成临时带参数的二维码
func (q *Qrcode) CreateTemp(req *TempQrcodeRequest) (*TempQrcodeReply, error) {
	result, err := client.PostJSON(fmt.Sprintf("%s/cgi-bin/qrcode/create?access_token=%s", q.client.GetBaseURL(), q.client.GetAccessToken()), req)
	if err != nil {
		return nil, err
	}
	reply := new(TempQrcodeReply)
	json.Unmarshal(result, reply)
	return reply, nil
}

// Download 下载二维码
func (q *Qrcode) Download(ticket string) ([]byte, error) {
	ticket = url.QueryEscape(ticket)
	result, err := client.Get("https://mp.weixin.qq.com/cgi-bin/showqrcode", map[string]string{
		"ticket": ticket,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateLimit 生成一个永久带参数的二维码
func (q *Qrcode) CreateLimit(req *LimitQrcodeRequest) (*LimitQrcodeReply, error) {
	result, err := client.PostJSON(fmt.Sprintf("%s/cgi-bin/qrcode/create?access_token=%s", q.client.GetBaseURL(), q.client.GetAccessToken()), req)
	if err != nil {
		return nil, err
	}
	reply := new(LimitQrcodeReply)
	json.Unmarshal(result, reply)
	return reply, nil
}
