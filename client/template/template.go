package template

import (
	"fmt"

	"github.com/nilorg/go-wechat/v2/client"
)

// Template 消息
type Template struct {
	client client.Clienter
}

// NewTemplate ...
func NewTemplate(c client.Clienter) *Template {
	return &Template{
		client: c,
	}
}

// Send 发送模板消息
// https://mp.weixin.qq.com/debug/cgi-bin/readtmpl?t=tmplmsg/faq_tmpl
func (t *Template) Send(data *ReplyInfo) error {
	_, err := client.PostJSON(fmt.Sprintf("%s/cgi-bin/message/template/send?access_token=%s", t.client.GetBaseURL(), t.client.GetAccessToken()), data)
	if err != nil {
		return err
	}
	return nil
}
