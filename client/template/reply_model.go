package template

import "encoding/json"

// DataItem ...
type DataItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// ReplyInfo 消息模板回复
type ReplyInfo struct {
	ToUser   string               `json:"touser"`
	ID       string               `json:"template_id"`
	URL      string               `json:"url"`
	TopColor string               `json:"topcolor"`
	Data     map[string]*DataItem `json:"data"`
}

// JSON ...
func (r *ReplyInfo) JSON() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}

// NewReplyInfo ...
func NewReplyInfo(tmplID string) *ReplyInfo {
	return &ReplyInfo{
		ID: tmplID,
	}
}
