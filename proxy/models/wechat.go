package models

// WechatRequestModel 微信请求模型
type WechatRequestModel struct {
	Signature string `form:"signature" binding:"required"`
	Timestamp int64  `form:"timestamp" binding:"required"`
	Nonce     int64  `form:"nonce" binding:"required"`
	OpenID    string `form:"openid" binding:"required"`
	Echostr   string `form:"echostr" binding:"required"`
}
