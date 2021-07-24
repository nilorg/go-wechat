package client

import (
	"encoding/xml"
	"time"

	"github.com/nilorg/go-wechat/v2/pkg/cdata"
)

// https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140454

const (
	// MessageFollowEventSubscribe subscribe(订阅)
	MessageFollowEventSubscribe = "subscribe"
	// MessageFollowEventUnsubscribe unsubscribe(取消订阅)
	MessageFollowEventUnsubscribe = "unsubscribe"
	// MessageMenuEventClick 点击菜单拉取消息时的事件推送
	MessageMenuEventClick = "CLICK"
	// MessageMenuEventView 点击菜单跳转链接时的事件推送
	MessageMenuEventView = "VIEW"
)

// MessageFollowAccept 关注事件
type MessageFollowAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int         `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，subscribe(订阅)、unsubscribe(取消订阅)
}

// MessageFollowAcceptParse ...
func MessageFollowAcceptParse(xmlValue []byte) *MessageFollowAccept {
	msg := &MessageFollowAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageScanQRCodeFollowAccept 二维码事件,用户未关注时，进行关注后的事件推送
type MessageScanQRCodeFollowAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，subscribe
	EventKey     cdata.CDATA `xml:"EventKey"`     // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket       cdata.CDATA `xml:"Ticket"`       // 二维码的ticket，可用来换取二维码图片
}

// MessageScanQRCodeFollowAcceptParse ...
func MessageScanQRCodeFollowAcceptParse(xmlValue []byte) *MessageScanQRCodeFollowAccept {
	msg := &MessageScanQRCodeFollowAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageScanQRCodeFollowedAccept 二维码事件,用户已关注时的事件推送
type MessageScanQRCodeFollowedAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，SCAN
	EventKey     cdata.CDATA `xml:"EventKey"`     // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket       cdata.CDATA `xml:"Ticket"`       // 二维码的ticket，可用来换取二维码图片
}

// MessageScanQRCodeFollowedAcceptParse ...
func MessageScanQRCodeFollowedAcceptParse(xmlValue []byte) *MessageScanQRCodeFollowedAccept {
	msg := &MessageScanQRCodeFollowedAccept{}
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

// MessageMenuPullAccept 点击菜单拉取消息时的事件推送
type MessageMenuPullAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，CLICK
	EventKey     cdata.CDATA `xml:"EventKey"`     // 事件KEY值，与自定义菜单接口中KEY值对应
}

// MessageMenuPullAcceptParse ...
func MessageMenuPullAcceptParse(xmlValue []byte) *MessageMenuPullAccept {
	msg := &MessageMenuPullAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageMenuSkipAccept 点击菜单跳转链接时的事件推送
type MessageMenuSkipAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，event
	Event        cdata.CDATA `xml:"Event"`        // 事件类型，VIEW
	EventKey     cdata.CDATA `xml:"EventKey"`     // 事件KEY值，设置的跳转URL
}

// MessageMenuSkipAcceptParse ...
func MessageMenuSkipAcceptParse(xmlValue []byte) *MessageMenuSkipAccept {
	msg := &MessageMenuSkipAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// accept

// 回复 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140543

// MessageTextAccept 文本消息
type MessageTextAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // text
	Content      cdata.CDATA `xml:"Content"`      // 文本消息内容
	MsgID        int64       `xml:"MsgId"`        // 消息id，64位整型
}

// MessageTextAcceptParse ...
func MessageTextAcceptParse(xmlValue []byte) *MessageTextAccept {
	msg := &MessageTextAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageImageAccept 图片消息
type MessageImageAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // image
	PicURL       cdata.CDATA `xml:"PicUrl"`       // 图片链接（由系统生成）
	MediaID      cdata.CDATA `xml:"MediaId"`      // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	MsgID        int64       `xml:"MsgId"`        // 消息id，64位整型
}

// MessageImageAcceptParse ...
func MessageImageAcceptParse(xmlValue []byte) *MessageImageAccept {
	msg := &MessageImageAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageVoiceAccept 语音消息
type MessageVoiceAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 语音为voice
	MediaID      cdata.CDATA `xml:"MediaId"`      // 语音消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Format       cdata.CDATA `xml:"Format"`       // 语音格式，如amr，speex等
	MsgID        int64       `xml:"MsgId"`        // 消息id，64位整型
}

// MessageVoiceAcceptParse ...
func MessageVoiceAcceptParse(xmlValue []byte) *MessageVoiceAccept {
	msg := &MessageVoiceAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageVideoAccept 视频消息
type MessageVideoAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 视频为video
	MediaID      cdata.CDATA `xml:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaID cdata.CDATA `xml:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	MsgID        int64       `xml:"MsgId"`        // 消息id，64位整型
}

// MessageVideoAcceptParse ...
func MessageVideoAcceptParse(xmlValue []byte) *MessageVideoAccept {
	msg := &MessageVideoAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageShortVideoAccept 小视频消息
type MessageShortVideoAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 小视频为shortvideo
	MediaID      cdata.CDATA `xml:"MediaId"`      // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaID cdata.CDATA `xml:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	MsgID        int64       `xml:"MsgId"`        // 消息id，64位整型
}

// MessageShortVideoAcceptParse ...
func MessageShortVideoAcceptParse(xmlValue []byte) *MessageShortVideoAccept {
	msg := &MessageShortVideoAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageLocationAccept 地理位置消息
type MessageLocationAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // location
	LocationX    cdata.CDATA `xml:"Location_X"`   //	地理位置维度
	LocationY    cdata.CDATA `xml:"Location_Y"`   //	地理位置经度
	Scale        cdata.CDATA `xml:"Scale"`        //	地图缩放大小
	Label        cdata.CDATA `xml:"Label"`        // 地理位置信息
	MsgID        int64       `xml:"MsgId"`        // 消息id，64位整型
}

// MessageLocationAcceptParse ...
func MessageLocationAcceptParse(xmlValue []byte) *MessageLocationAccept {
	msg := &MessageLocationAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// MessageLinkAccept 链接消息
type MessageLinkAccept struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 开发者微信号
	FromUserName cdata.CDATA `xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // 消息类型，link
	Title        cdata.CDATA `xml:"Title"`        //	消息标题
	Description  cdata.CDATA `xml:"Description"`  //	消息描述
	URL          cdata.CDATA `xml:"Url"`          //	消息链接
	MsgID        int64       `xml:"MsgId"`        // 消息id，64位整型
}

// MessageLinkAcceptParse ...
func MessageLinkAcceptParse(xmlValue []byte) *MessageLinkAccept {
	msg := &MessageLinkAccept{}
	if err := xml.Unmarshal(xmlValue, msg); err != nil {
		return nil
	}
	return msg
}

// reply

// 回复 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140543

// MessageTextReply 文本消息
type MessageTextReply struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`   // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA `xml:"FromUserName"` // 开发者微信号
	CreateTime   int64       `xml:"CreateTime"`   // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`      // text
	Content      cdata.CDATA `xml:"Content"`      // 回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
}

// XML ...
func (t *MessageTextReply) XML() string {
	bytes, _ := xml.Marshal(t)
	return string(bytes)
}

// NewMessageTextReply ...
func NewMessageTextReply() *MessageTextReply {
	return &MessageTextReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "text",
	}
}

// MessageImageReply 图片消息
type MessageImageReply struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`    // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA `xml:"FromUserName"`  // 开发者微信号
	CreateTime   int64       `xml:"CreateTime"`    // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`       // image
	MediaID      cdata.CDATA `xml:"Image>MediaId"` // 通过素材管理中的接口上传多媒体文件，得到的id。
}

// XML ...
func (i *MessageImageReply) XML() string {
	bytes, _ := xml.Marshal(i)
	return string(bytes)
}

// NewMessageImageReply ...
func NewMessageImageReply() *MessageImageReply {
	return &MessageImageReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "image",
	}
}

// MessageVoiceReply 语音消息
type MessageVoiceReply struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cdata.CDATA `xml:"ToUserName"`    // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA `xml:"FromUserName"`  // 开发者微信号
	CreateTime   int64       `xml:"CreateTime"`    // 消息创建时间 （整型）
	MsgType      cdata.CDATA `xml:"MsgType"`       // 语音，voice
	MediaID      cdata.CDATA `xml:"Voice>MediaId"` // 通过素材管理中的接口上传多媒体文件，得到的id
}

// XML ...
func (v *MessageVoiceReply) XML() string {
	bytes, _ := xml.Marshal(v)
	return string(bytes)
}

// NewMessageVoiceReply ...
func NewMessageVoiceReply() *MessageVoiceReply {
	return &MessageVoiceReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "voice",
	}
}

// MessageVideoReply 视频消息
type MessageVideoReply struct {
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
func (v *MessageVideoReply) XML() string {
	bytes, _ := xml.Marshal(v)
	return string(bytes)
}

// NewMessageVideoReply ...
func NewMessageVideoReply() *MessageVideoReply {
	return &MessageVideoReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "video",
	}
}

// MessageMusicReply 音乐消息
type MessageMusicReply struct {
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
func (m *MessageMusicReply) XML() string {
	bytes, _ := xml.Marshal(m)
	return string(bytes)
}

// NewMessageMusicReply ...
func NewMessageMusicReply() *MessageMusicReply {
	return &MessageMusicReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "music",
	}
}

// MessageNewsReply 回复图文消息
type MessageNewsReply struct {
	XMLName      xml.Name              `xml:"xml"`
	ToUserName   cdata.CDATA           `xml:"ToUserName"`    // 接收方帐号（收到的OpenID）
	FromUserName cdata.CDATA           `xml:"FromUserName"`  // 开发者微信号
	CreateTime   int64                 `xml:"CreateTime"`    // 消息创建时间 （整型）
	MsgType      cdata.CDATA           `xml:"MsgType"`       // news
	ArticleCount cdata.CDATA           `xml:"ArticleCount"`  // 图文消息个数，限制为8条以内
	Articles     []*MessageNewsArticle `xml:"Articles>item"` // 多条图文消息信息，默认第一个item为大图,注意，如果图文数超过8，则将会无响应
}

// MessageNewsArticle ...
type MessageNewsArticle struct {
	Title       cdata.CDATA `xml:"Title"`       // 图文消息标题
	Description cdata.CDATA `xml:"Description"` // 图文消息描述
	PicURL      cdata.CDATA `xml:"PicUrl"`      // 图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
	URL         cdata.CDATA `xml:"Url"`         // 点击图文消息跳转链接
}

// XML ...
func (n *MessageNewsReply) XML() string {
	bytes, _ := xml.Marshal(n)
	return string(bytes)
}

// NewMessageNewsReply ...
func NewMessageNewsReply() *MessageNewsReply {
	return &MessageNewsReply{
		CreateTime: time.Now().Unix(),
		MsgType:    "news",
	}
}
