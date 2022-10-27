package models

// WechatRequestModel 微信请求模型
type WechatRequestModel struct {
	MsgSignature string `form:"msg_signature"` // 企业微信
	Signature    string `form:"signature"`
	Timestamp    int64  `form:"timestamp" binding:"required"`
	Nonce        int64  `form:"nonce" binding:"required"`
	OpenID       string `form:"openid" binding:"-"`
	Echostr      string `form:"echostr" binding:"required"`
}
