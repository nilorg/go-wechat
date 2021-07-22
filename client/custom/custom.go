package custom

import (
	"fmt"

	"github.com/nilorg/go-wechat/v2/client"
)

// Custom 客服
type Custom struct {
	client client.Clienter
}

// NewCustom ...
func NewCustom(c client.Clienter) *Custom {
	return &Custom{
		client: c,
	}
}

// send 发消息
func (c *Custom) send(req interface{}) error {
	_, err := client.PostJSON(fmt.Sprintf("%s/cgi-bin/message/custom/send?access_token=%s", c.client.GetBaseURL(), c.client.GetAccessToken()), req)
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

// SendMenu 发送菜单消息
func (c *Custom) SendMenu(req *MenuRequest) error {
	return c.send(req)
}
