package ginplus

import (
	"github.com/gin-gonic/gin"
	"github.com/nilorg/go-wechat/v2/client"
	"github.com/nilorg/go-wechat/v2/pkg/cdata"
)

// WechatContext include gin context and local context
type WechatContext struct {
	*gin.Context
	WxClient client.Clienter
}

// WechatContextFunction wechat gin.Contenxt and local model.Context
func WechatContextFunction(ctlFunc func(ctx *WechatContext), client client.Clienter) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		wechatContext := &WechatContext{
			Context:  ctx,
			WxClient: client,
		}

		ctlFunc(wechatContext)
	}
}

// ReplyXML 回复xml
func (w *WechatContext) ReplyXML(xml cdata.XMLer) {
	w.String(200, xml.XML())
}
