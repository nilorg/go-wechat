package gin

import (
	ggin "github.com/gin-gonic/gin"
	wechat "github.com/nilorg/go-wechat"
)

// WechatContext include gin context and local context
type WechatContext struct {
	*ggin.Context
	WxClient wechat.Clienter
}

// WechatContextFunction wechat gin.Contenxt and local model.Context
func WechatContextFunction(ctlFunc func(ctx *WechatContext), client wechat.Clienter) ggin.HandlerFunc {
	return func(ctx *ggin.Context) {

		wechatContext := &WechatContext{
			Context:  ctx,
			WxClient: client,
		}

		ctlFunc(wechatContext)
	}
}

// ReplyXML 回复xml
func (w *WechatContext) ReplyXML(xml wechat.XMLer) {
	w.String(200, xml.XML())
}
