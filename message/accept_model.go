package message

import (
	"encoding/xml"

	wechat "github.com/nilorg/go-wechat"
)

// 回复 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140543

// TextAccept 文本消息
type TextAccept struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   wechat.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName wechat.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64        `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      wechat.CDATA `xml:"MsgType"`      // text
	Content      wechat.CDATA `xml:"Content"`      // 文本消息内容
	MsgID        int64        `xml:"MsgId"`        // 消息id，64位整型
}

// TextAcceptParse ...
func TextAcceptParse(xmlValue []byte) *TextAccept {
	msg := &TextAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// ImageAccept 图片消息
type ImageAccept struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   wechat.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName wechat.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64        `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      wechat.CDATA `xml:"MsgType"`      // image
	PicURL       wechat.CDATA `xml:"PicUrl"`       // 图片链接（由系统生成）
	MediaID      wechat.CDATA `xml:"MediaId"`      // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	MsgID        int64        `xml:"MsgId"`        // 消息id，64位整型
}

// ImageAcceptParse ...
func ImageAcceptParse(xmlValue []byte) *ImageAccept {
	msg := &ImageAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// VoiceAccept 语音消息
type VoiceAccept struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   wechat.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName wechat.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64        `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      wechat.CDATA `xml:"MsgType"`      // 语音为voice
	MediaID      wechat.CDATA `xml:"MediaId"`      // 语音消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Format       wechat.CDATA `xml:"Format"`       // 语音格式，如amr，speex等
	MsgID        int64        `xml:"MsgId"`        // 消息id，64位整型
}

// VoiceAcceptParse ...
func VoiceAcceptParse(xmlValue []byte) *VoiceAccept {
	msg := &VoiceAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// VideoAccept 视频消息
type VideoAccept struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   wechat.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName wechat.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64        `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      wechat.CDATA `xml:"MsgType"`      // 视频为video
	MediaID      wechat.CDATA `xml:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaID wechat.CDATA `xml:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	MsgID        int64        `xml:"MsgId"`        // 消息id，64位整型
}

// VideoAcceptParse ...
func VideoAcceptParse(xmlValue []byte) *VideoAccept {
	msg := &VideoAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// ShortVideoAccept 小视频消息
type ShortVideoAccept struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   wechat.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName wechat.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64        `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      wechat.CDATA `xml:"MsgType"`      // 小视频为shortvideo
	MediaID      wechat.CDATA `xml:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaID wechat.CDATA `xml:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	MsgID        int64        `xml:"MsgId"`        // 消息id，64位整型
}

// ShortVideoAcceptParse ...
func ShortVideoAcceptParse(xmlValue []byte) *ShortVideoAccept {
	msg := &ShortVideoAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// LocationAccept 地理位置消息
type LocationAccept struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   wechat.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName wechat.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64        `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      wechat.CDATA `xml:"MsgType"`      // location
	LocationX    wechat.CDATA `xml:"Location_X"`   //	地理位置维度
	LocationY    wechat.CDATA `xml:"Location_Y"`   //	地理位置经度
	Scale        wechat.CDATA `xml:"Scale"`        //	地图缩放大小
	Label        wechat.CDATA `xml:"Label"`        // 地理位置信息
	MsgID        int64        `xml:"MsgId"`        // 消息id，64位整型
}

// LocationAcceptParse ...
func LocationAcceptParse(xmlValue []byte) *LocationAccept {
	msg := &LocationAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// LinkAccept 链接消息
type LinkAccept struct {
	XMLName      xml.Name     `xml:"xml"`
	ToUserName   wechat.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName wechat.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64        `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      wechat.CDATA `xml:"MsgType"`      // 消息类型，link
	Title        wechat.CDATA `xml:"Title"`        //	消息标题
	Description  wechat.CDATA `xml:"Description"`  //	消息描述
	URL          wechat.CDATA `xml:"Url"`          //	消息链接
	MsgID        int64        `xml:"MsgId"`        // 消息id，64位整型
}

// LinkAcceptParse ...
func LinkAcceptParse(xmlValue []byte) *LinkAccept {
	msg := &LinkAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}
