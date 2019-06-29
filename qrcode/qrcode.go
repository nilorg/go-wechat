package qrcode

import (
	"encoding/json"
	"net/url"

	wechat "github.com/nilorg/go-wechat"
)

// Qrcode 二维码
type Qrcode struct {
	client wechat.Clienter
}

// NewQrcode ...
func NewQrcode(c wechat.Clienter) *Qrcode {
	return &Qrcode{
		client: c,
	}
}

// CreateTemp 生成临时带参数的二维码
func (q *Qrcode) CreateTemp(req *TempQrcodeRequest) (*TempQrcodeReply, error) {
	result, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="+q.client.GetAccessToken(), req)
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
	result, err := wechat.Get("https://mp.weixin.qq.com/cgi-bin/showqrcode", map[string]string{
		"ticket": ticket,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateLimit 生成一个永久带参数的二维码
func (q *Qrcode) CreateLimit(req *LimitQrcodeRequest) (*LimitQrcodeReply, error) {
	result, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="+q.client.GetAccessToken(), req)
	if err != nil {
		return nil, err
	}
	reply := new(LimitQrcodeReply)
	json.Unmarshal(result, reply)
	return reply, nil
}
