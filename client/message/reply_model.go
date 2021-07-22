package message

import (
	"encoding/xml"
	"time"

	"github.com/nilorg/go-wechat/v2/pkg/cdata"
)

// 回复 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140543

// TextReply 文本消息
type TextReply struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA `xml:"FromUserName"` // 开发者微信号
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // text
	Content      cdata.CDATA `xml:"Content"`      // 回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
}

// XML ...
func (t *TextReply) XML() string {
	bytes, _ := xml.Marshal(t)
	return string(bytes)
}

// NewTextReply ...
func NewTextReply() *TextReply {
	return &TextReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "text",
	}
}

// ImageReply 图片消息
type ImageReply struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`    // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA `xml:"FromUserName"`  // 开发者微信号
	CreateTime   int64       `xml:"CreateTime"`    // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`       // image
	MediaID      cdata.CDATA `xml:"Image>MediaId"` // 通过素材管理中的接口上传多媒体文件，得到的id。
}

// XML ...
func (i *ImageReply) XML() string {
	bytes, _ := xml.Marshal(i)
	return string(bytes)
}

// NewImageReply ...
func NewImageReply() *ImageReply {
	return &ImageReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "image",
	}
}

// VoiceReply 语音消息
type VoiceReply struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`    // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA `xml:"FromUserName"`  // 开发者微信号
	CreateTime   int64       `xml:"CreateTime"`    // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`       // 语音，voice
	MediaID      cdata.CDATA `xml:"Voice>MediaId"` // 通过素材管理中的接口上传多媒体文件，得到的id
}

// XML ...
func (v *VoiceReply) XML() string {
	bytes, _ := xml.Marshal(v)
	return string(bytes)
}

// NewVoiceReply ...
func NewVoiceReply() *VoiceReply {
	return &VoiceReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "voice",
	}
}

// VideoReply 视频消息
type VideoReply struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`    // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA `xml:"FromUserName"`  // 开发者微信号
	CreateTime   int64       `xml:"CreateTime"`    // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`       // video
	MediaID      cdata.CDATA `xml:"Video>MediaId"` // 通过素材管理中的接口上传多媒体文件，得到的id
	Title        cdata.CDATA `xml:"Title"`         // 视频消息的标题
	Description  cdata.CDATA `xml:"Description"`   // 视频消息的描述
}

// XML ...
func (v *VideoReply) XML() string {
	bytes, _ := xml.Marshal(v)
	return string(bytes)
}

// NewVideoReply ...
func NewVideoReply() *VideoReply {
	return &VideoReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "video",
	}
}

// MusicReply 音乐消息
type MusicReply struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA `xml:"FromUserName"` // 开发者微信号
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // music
	Title        cdata.CDATA `xml:"Title"`        // 音乐标题
	Description  cdata.CDATA `xml:"Description"`  // 音乐描述
	MusicURL     cdata.CDATA `xml:"MusicURL"`     // 音乐链接
	HQMusicURL   cdata.CDATA `xml:"HQMusicUrl"`   // 高质量音乐链接，WIFI环境优先使用该链接播放音乐
	ThumbMediaID cdata.CDATA `xml:"ThumbMediaId"` // 缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id
}

// XML ...
func (m *MusicReply) XML() string {
	bytes, _ := xml.Marshal(m)
	return string(bytes)
}

// NewMusicReply ...
func NewMusicReply() *MusicReply {
	return &MusicReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "music",
	}
}

// NewsReply 回复图文消息
type NewsReply struct {
	XMLName      xml.Name       `xml:"xml"`
	ToUserName   cdata.CDATA    `xml:"ToUserName"`    // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA    `xml:"FromUserName"`  // 开发者微信号
	CreateTime   int64          `xml:"CreateTime"`    // 消息创建时间 （整型）
	MsgType      cdata.CDATA    `xml:"MsgType"`       // news
	ArticleCount cdata.CDATA    `xml:"ArticleCount"`  // 图文消息个数，限制为8条以内
	Articles     []*NewsArticle `xml:"Articles>item"` // 多条图文消息信息，默认第一个item为大图,注意，如果图文数超过8，则将会无响应
}

// NewsArticle ...
type NewsArticle struct {
	Title       cdata.CDATA `xml:"Title"`       // 图文消息标题
	Description cdata.CDATA `xml:"Description"` // 图文消息描述
	PicURL      cdata.CDATA `xml:"PicUrl"`      // 图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
	URL         cdata.CDATA `xml:"Url"`         // 点击图文消息跳转链接
}

// XML ...
func (n *NewsReply) XML() string {
	bytes, _ := xml.Marshal(n)
	return string(bytes)
}

// NewNewsReply ...
func NewNewsReply() *NewsReply {
	return &NewsReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "news",
	}
}
