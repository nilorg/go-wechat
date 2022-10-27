package client

import (
	"encoding/json"
	"fmt"
)

// DraftAddRequest 新建草稿
type DraftAddRequest struct {
	Articles []*DraftArticle `json:"articles"`
}

type DraftArticle struct {
	Title              string `json:"title"`                 // 标题
	Author             string `json:"author"`                // 作者
	Digest             string `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。
	Content            string `json:"content"`               // 图文消息的具体内容，支持 HTML 标签，必须少于2万字符，小于1M，且此处会去除 JS ,涉及图片 url 必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片 url 将被过滤。
	ContentSourceURL   string `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL
	ThumbMediaID       string `json:"thumb_media_id"`        // 图文消息的封面图片素材id（必须是永久MediaID）
	NeedOpenComment    int    `json:"need_open_comment"`     // 是否打开评论，0不打开(默认)，1打开
	OnlyFansCanComment int    `json:"only_fans_can_comment"` // 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论
}

// DraftAddReply 新建草稿回复
type DraftAddReply struct {
	MediaID string `json:"media_id"` // 上传后的获取标志，长度不固定，但不会超过 128 字符
}

// DraftAdd 新建草稿
func (c *Client) DraftAdd(req *DraftAddRequest) (*DraftAddReply, error) {
	url := fmt.Sprintf("%s/cgi-bin/draft/add", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, req)
	if err != nil {
		return nil, err
	}
	reply := new(DraftAddReply)
	err = json.Unmarshal(result, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// DraftGetReply 获取草稿回复
type DraftGetReply struct {
	NewsItem []*DraftNewsItem `json:"news_item"` // 多图文消息应有多段 news_item 结构
}

type DraftNewsItem struct {
	Title              string `json:"title"`                 // 标题
	Author             string `json:"author"`                // 作者
	Digest             string `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。
	Content            string `json:"content"`               // 图文消息的具体内容，支持 HTML 标签，必须少于2万字符，小于1M，且此处会去除 JS ,涉及图片 url 必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片 url 将被过滤。
	ContentSourceURL   string `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL
	ThumbMediaID       string `json:"thumb_media_id"`        // 图文消息的封面图片素材id（必须是永久MediaID）
	NeedOpenComment    int    `json:"need_open_comment"`     // 是否打开评论，0不打开(默认)，1打开
	OnlyFansCanComment int    `json:"only_fans_can_comment"` // 是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论
	URL                string `json:"url"`                   // 草稿的临时链接
}

// DraftAdd 新建草稿
func (c *Client) DraftGet(mediaID string) (*DraftGetReply, error) {
	url := fmt.Sprintf("%s/cgi-bin/draft/add", c.opts.BaseURL)
	if !c.opts.Proxy {
		url += fmt.Sprintf("?access_token=%s", c.opts.Token.GetAccessToken())
	}
	result, err := PostJSON(url, map[string]interface{}{
		"media_id": mediaID,
	})
	if err != nil {
		return nil, err
	}
	reply := new(DraftGetReply)
	err = json.Unmarshal(result, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
