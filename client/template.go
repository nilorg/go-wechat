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

type TemplateMiniProgram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

// TemplateSendRequest 消息模板
type TemplateSendRequest struct {
	ToUser      string                       `json:"touser"`
	ID          string                       `json:"template_id"`
	URL         string                       `json:"url"`
	Data        map[string]*TemplateDataItem `json:"data"`
	MiniProgram *TemplateMiniProgram         `json:"miniprogram"`
	ClientMsgID string                       `json:"client_msg_id"`
}

// JSON ...
func (r *TemplateSendRequest) JSON() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

// NewTemplateReplyInfo ...
func NewTemplateReplyInfo(tmplID string) *TemplateSendRequest {
	return &TemplateSendRequest{
		ID: tmplID,
	}
}

// Send 发送模板消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#5
func (c *Client) TemplateSend(req *TemplateSendRequest) error {
	url := fmt.Sprintf("%s/cgi-bin/message/template/send", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	_, err := PostJSON(c.opts.HttpClient, url, req)
	if err != nil {
		return err
	}
	return nil
}
