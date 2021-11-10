package client

import (
	"encoding/json"
	"fmt"
)

// TemplateDataItem ...
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// TemplateReplyInfo 消息模板回复
type TemplateReplyInfo struct {
	ToUser   string                       `json:"touser"`
	ID       string                       `json:"template_id"`
	URL      string                       `json:"url"`
	TopColor string                       `json:"topcolor"`
	Data     map[string]*TemplateDataItem `json:"data"`
}

// JSON ...
func (r *TemplateReplyInfo) JSON() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

// NewTemplateReplyInfo ...
func NewTemplateReplyInfo(tmplID string) *TemplateReplyInfo {
	return &TemplateReplyInfo{
		ID: tmplID,
	}
}

// Send 发送模板消息
// https://mp.weixin.qq.com/debug/cgi-bin/readtmpl?t=tmplmsg/faq_tmpl
func (c *Client) TemplateSend(data *TemplateReplyInfo) error {
	url := fmt.Sprintf("%s/cgi-bin/message/template/send", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	_, err := PostJSON(url, data)
	if err != nil {
		return err
	}
	return nil
}
