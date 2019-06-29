package template

import (
	wechat "github.com/nilorg/go-wechat"
)

// Template 消息
type Template struct {
	client wechat.Clienter
}

// NewTemplate ...
func NewTemplate(c wechat.Clienter) *Template {
	return &Template{
		client: c,
	}
}

// Send 发送模板消息
// https://mp.weixin.qq.com/debug/cgi-bin/readtmpl?t=tmplmsg/faq_tmpl
func (t *Template) Send(data *ReplyInfo) error {
	_, err := wechat.PostJSON("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+t.client.GetAccessToken(), data)
	if err != nil {
		return err
	}
	return nil
}
