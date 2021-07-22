package message

import (
	"encoding/xml"

	"github.com/nilorg/go-wechat/v2/pkg/cdata"
)

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140454

const (
	// FollowEventSubscribe subscribe(订阅)
	FollowEventSubscribe = "subscribe"
	// FollowEventUnsubscribe unsubscribe(取消订阅)
	FollowEventUnsubscribe = "unsubscribe"
	// MenuEventClick 点击菜单拉取消息时的事件推送
	MenuEventClick = "CLICK"
	// MenuEventView 点击菜单跳转链接时的事件推送
	MenuEventView = "VIEW"
)

// FollowAccept 关注事件
type FollowAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int         `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，subscribe(订阅)、unsubscribe(取消订阅)
}

// FollowAcceptParse ...
func FollowAcceptParse(xmlValue []byte) *FollowAccept {
	msg := &FollowAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// ScanQRCodeFollowAccept 二维码事件,用户未关注时，进行关注后的事件推送
type ScanQRCodeFollowAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，subscribe
	EventKey     cdata.CDATA `xml:"EventKey"`     // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket       cdata.CDATA `xml:"Ticket"`       // 二维码的ticket，可用来换取二维码图片
}

// ScanQRCodeFollowAcceptParse ...
func ScanQRCodeFollowAcceptParse(xmlValue []byte) *ScanQRCodeFollowAccept {
	msg := &ScanQRCodeFollowAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// ScanQRCodeFollowedAccept 二维码事件,用户已关注时的事件推送
type ScanQRCodeFollowedAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，SCAN
	EventKey     cdata.CDATA `xml:"EventKey"`     // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket       cdata.CDATA `xml:"Ticket"`       // 二维码的ticket，可用来换取二维码图片
}

// ScanQRCodeFollowedAcceptParse ...
func ScanQRCodeFollowedAcceptParse(xmlValue []byte) *ScanQRCodeFollowedAccept {
	msg := &ScanQRCodeFollowedAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// // LocationAccept 上报地理位置事件
// // 用户同意上报地理位置后，每次进入公众号会话时，都会在进入时上报地理位置，或在进入会话后每5秒上报一次地理位置，公众号可以在公众平台网站中修改以上设置。
// // 上报地理位置时，微信会将上报地理位置事件推送到开发者填写的URL。
// type LocationAccept struct {
// 	XMLName      xml.Name `xml:"xml"`
// 	ToUserName   string   `xml:"ToUserName,cdata"`   // 开发者微信号
// 	FromUserName string   `xml:"FromUserName,cdata"` // 发送方帐号（一个OpenID）
// 	CreateTime   int64    `xml:"CreateTime"`         // 消息创建时间 （整型）
// 	MsgType      string   `xml:"MsgType,cdata"`      // 消息类型，event
// 	Event        string   `xml:"Event,cdata"`        // 事件类型，LOCATION
// 	Latitude     string   `xml:"Latitude,cdata"`     // 地理位置纬度
// 	Longitude    string   `xml:"Longitude,cdata"`    // 地理位置经度
// 	Precision    string   `xml:"Precision,cdata"`    // 地理位置精度
// }

// MenuPullAccept 点击菜单拉取消息时的事件推送
type MenuPullAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，CLICK
	EventKey     cdata.CDATA `xml:"EventKey"`     // 事件KEY值，与自定义菜单接口中KEY值对应
}

// MenuPullAcceptParse ...
func MenuPullAcceptParse(xmlValue []byte) *MenuPullAccept {
	msg := &MenuPullAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MenuSkipAccept 点击菜单跳转链接时的事件推送
type MenuSkipAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，VIEW
	EventKey     cdata.CDATA `xml:"EventKey"`     // 事件KEY值，设置的跳转URL
}

// MenuSkipAcceptParse ...
func MenuSkipAcceptParse(xmlValue []byte) *MenuSkipAccept {
	msg := &MenuSkipAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}
