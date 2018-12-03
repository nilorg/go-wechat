package custom

import (
	wechat "github.com/nilorg/go-wechat"
	"github.com/nilorg/go-wechat/context"
)

// Custom 客服
type Custom struct {
	context *context.Context
}

// NewCustom ...
func NewCustom(c *context.Context) *Custom {
	return &Custom{
		context: c,
	}
}

// send 发消息
func (c *Custom) send(req interface{}) error {
	_, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="+c.context.GetAccessToken(), req)
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
