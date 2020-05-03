package custom

import (
	wechat "github.com/nilorg/go-wechat"
)

// Custom 客服
type Custom struct {
	client wechat.Clienter
}

// NewCustom ...
func NewCustom(c wechat.Clienter) *Custom {
	return &Custom{
		client: c,
	}
}

// send 发消息
func (c *Custom) send(req interface{}) error {
	_, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="+c.client.GetAccessToken(), req)
	return err
}

// SendText 发送文本消息
func (c *Custom) SendText(req *TextRequest) error {
	return c.send(req)
}

// SendImage 发送图片消息
func (c *Custom) SendImage(req *ImageRequest) error {
	return c.send(req)
}

// SendNews 发送图文消息（点击跳转到外链）
// 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
func (c *Custom) SendNews(req *NewsRequest) error {
	return c.send(req)
}
