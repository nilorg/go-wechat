package client

import (
	"encoding/json"
	"fmt"
)

type QiyeKfSyncMessageRequest struct {
	Cursor      string `json:"cursor"`       // 上一次调用时返回的next_cursor，第一次拉取可以不填。若不填，从3天内最早的消息开始返回。
	Token       string `json:"token"`        // 回调事件返回的token字段，10分钟内有效；可不填，如果不填接口有严格的频率限制。
	Limit       int    `json:"limit"`        // 期望请求的数据量，默认值和最大值都为1000。注意：可能会出现返回条数少于limit的情况，需结合返回的has_more字段判断是否继续请求。
	VoiceFormat int    `json:"voice_format"` // 语音消息类型，0-Amr 1-Silk，默认0。可通过该参数控制返回的语音格式，开发者可按需选择自己程序支持的一种格式
}

type QiyeKfSyncMessageResponse struct {
	NextCursor string                   `json:"next_cursor"` // 下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。
	HasMore    int                      `json:"has_more"`    // 是否还有更多数据。0-否；1-是。不能通过判断msg_list是否空来停止拉取，可能会出现has_more为1，而msg_list为空的情况
	MsgList    []*QiyeKfSyncMessageItem `json:"msg_list"`    // 消息列表
}

type QiyeKfSyncMessageItem struct {
	MsgID          string                         `json:"msgid"`                   // 消息ID
	OpenKfID       string                         `json:"open_kfid"`               // 客服帐号ID（msgtype为event，该字段不返回）
	ExternalUserID string                         `json:"external_userid"`         // 客户UserID（msgtype为event，该字段不返回）
	SendTime       int                            `json:"send_time"`               // 消息发送时间
	Origin         int                            `json:"origin"`                  // 消息来源。3-微信客户发送的消息 4-系统推送的事件消息 5-接待人员在企业微信客户端发送的消息
	ServicerUserID string                         `json:"servicer_userid"`         // 从企业微信给客户发消息的接待人员userid（即仅origin为5才返回；msgtype为event，该字段不返回）
	MsgType        string                         `json:"msgtype"`                 // 对不同的msgtype，有相应的结构描述，下面进一步说明
	Text           *QiyeKfSyncMessageText         `json:"text,omitempty"`          // 文本消息
	Image          *QiyeKfSyncMessageImage        `json:"image,omitempty"`         // 图片消息
	Voice          *QiyeKfSyncMessageVoice        `json:"voice,omitempty"`         // 语音消息
	File           *QiyeKfSyncMessageFile         `json:"file,omitempty"`          // 文件消息
	Location       *QiyeKfSyncMessageLocation     `json:"location,omitempty"`      // 位置消息
	Link           *QiyeKfSyncMessageLink         `json:"link,omitempty"`          // 链接消息
	BusinessCard   *QiyeKfSyncMessageBusinessCard `json:"business_card,omitempty"` // 名片消息
	Miniprogram    *QiyeKfSyncMessageMiniprogram  `json:"miniprogram,omitempty"`   // 小程序消息
	Msgmenu        *QiyeKfSyncMessageMenu         `json:"msgmenu,omitempty"`       // 消息菜单
	Event          *QiyeKfSyncMessageEvent        `json:"event,omitempty"`         // 事件消息
}

type QiyeKfSyncMessageText struct {
	Content string `json:"content"` // 文本内容
	MenuID  string `json:"menu_id"` // 客户点击菜单消息，触发的回复消息中附带的菜单ID
}

type QiyeKfSyncMessageImage struct {
	MediaID string `json:"media_id"` // 图片文件id
}

type QiyeKfSyncMessageVoice struct {
	MediaID string `json:"media_id"` // 语音文件id
}
type QiyeKfSyncMessageFile struct {
	MediaID string `json:"media_id"` // 文件id
}
type QiyeKfSyncMessageLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}
type QiyeKfSyncMessageLink struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
	PicURL string `json:"pic_url"`
}
type QiyeKfSyncMessageBusinessCard struct {
	UserID string `json:"userid"`
}
type QiyeKfSyncMessageMiniprogram struct {
	Title        string `json:"title"`
	AppID        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

type QiyeKfSyncMessageMenuItemClick struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
type QiyeKfSyncMessageMenuItemView struct {
	URL     string `json:"url"`
	Content string `json:"content"`
}
type QiyeKfSyncMessageMenuItemMiniprogram struct {
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Content  string `json:"content"`
}
type QiyeKfSyncMessageMenuItem struct {
	Type        string                                `json:"type"`
	Click       *QiyeKfSyncMessageMenuItemClick       `json:"click,omitempty"`
	View        *QiyeKfSyncMessageMenuItemView        `json:"view,omitempty"`
	Miniprogram *QiyeKfSyncMessageMenuItemMiniprogram `json:"miniprogram,omitempty"`
}
type QiyeKfSyncMessageMenu struct {
	HeadContent string                       `json:"head_content"`
	List        []*QiyeKfSyncMessageMenuItem `json:"list"`
	TailContent string                       `json:"tail_content"`
}

type QiyeKfSyncMessageEventWechatChannels struct {
	Nickname string `json:"nickname"`
}
type QiyeKfSyncMessageEvent struct {
	EventType         string                                `json:"event_type"`
	OpenKfID          string                                `json:"open_kfid"`
	ExternalUserID    string                                `json:"external_userid"`
	Scene             string                                `json:"scene,omitempty"`
	SceneParam        string                                `json:"scene_param,omitempty"`
	WelcomeCode       string                                `json:"welcome_code,omitempty"`
	WechatChannels    *QiyeKfSyncMessageEventWechatChannels `json:"wechat_channels,omitempty"`
	FailMsgID         string                                `json:"fail_msgid,omitempty"`
	FailType          int                                   `json:"fail_type,omitempty"`
	ServicerUserID    string                                `json:"servicer_userid,omitempty"`
	Status            int                                   `json:"status,omitempty"`
	ChangeType        int                                   `json:"change_type,omitempty"`
	OldServicerUserID string                                `json:"old_servicer_userid,omitempty"`
	NewServicerUserID string                                `json:"new_servicer_userid,omitempty"`
	MsgCode           string                                `json:"msg_code,omitempty"`
	RecallMsgID       string                                `json:"recall_msgid,omitempty"`
}

// KfSyncMessage 获取客服消息
func (c *QiyeClient) KfSyncMessage(req *QiyeKfSyncMessageRequest) (*QiyeKfSyncMessageResponse, error) {
	url := fmt.Sprintf("%s/cgi-bin/kf/sync_msg", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	resp := new(QiyeKfSyncMessageResponse)
	err = json.Unmarshal(result, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type QiyeKfSendMessageRequest struct {
	ToUser       string                         `json:"touser"`
	OpenKfID     string                         `json:"open_kfid"`
	MsgID        string                         `json:"msgid"`
	MsgType      string                         `json:"msgtype"`
	Text         *QiyeKfSyncMessageText         `json:"text,omitempty"`          // 文本消息
	Image        *QiyeKfSyncMessageImage        `json:"image,omitempty"`         // 图片消息
	Voice        *QiyeKfSyncMessageVoice        `json:"voice,omitempty"`         // 语音消息
	File         *QiyeKfSyncMessageFile         `json:"file,omitempty"`          // 文件消息
	Location     *QiyeKfSyncMessageLocation     `json:"location,omitempty"`      // 位置消息
	Link         *QiyeKfSyncMessageLink         `json:"link,omitempty"`          // 链接消息
	BusinessCard *QiyeKfSyncMessageBusinessCard `json:"business_card,omitempty"` // 名片消息
	Miniprogram  *QiyeKfSyncMessageMiniprogram  `json:"miniprogram,omitempty"`   // 小程序消息
	Msgmenu      *QiyeKfSyncMessageMenu         `json:"msgmenu,omitempty"`       // 消息菜单
}

type QiyeKfSendMessageResponse struct {
	MsgID string `json:"msgid"`
}

// KfSendMessage 发送客服消息
func (c *QiyeClient) KfSendMessage(req *QiyeKfSendMessageRequest) (*QiyeKfSendMessageResponse, error) {
	url := fmt.Sprintf("%s/cgi-bin/kf/send_msg", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	resp := new(QiyeKfSendMessageResponse)
	err = json.Unmarshal(result, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type QiyeKfSendMsgOnEventRequest struct {
	Code    string                 `json:"code"`
	MsgID   string                 `json:"msgid"`
	MsgType string                 `json:"msgtype"`
	Text    *QiyeKfSyncMessageText `json:"text,omitempty"`    // 文本消息
	Msgmenu *QiyeKfSyncMessageMenu `json:"msgmenu,omitempty"` // 消息菜单
}

type QiyeKfSendMsgOnEventResponse struct {
	MsgID string `json:"msgid"`
}

// KfSendMessage 发送客服消息
func (c *QiyeClient) KfSendMsgOnEvent(req *QiyeKfSendMsgOnEventRequest) (*QiyeKfSendMsgOnEventResponse, error) {
	url := fmt.Sprintf("%s/cgi-bin/kf/send_msg_on_event", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	resp := new(QiyeKfSendMsgOnEventResponse)
	err = json.Unmarshal(result, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
