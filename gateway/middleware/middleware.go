package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nilorg/go-wechat/v2/gateway/models"
)

// Header 头处理
func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Server", "Wechat-Gateway")
		c.Next()
	}
}

// CheckWechatUser 检查微信用户是否存在
func CheckWechatUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var model models.WechatRequestModel
		if err := ctx.ShouldBindWith(&model, binding.Query); err == nil {
			ctx.Set("current", model)
		}
		ctx.Next()
	}
}
