package middleware

import (
	"github.com/gin-gonic/gin"
)

// Header 头处理
func Header() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Server", "Wechat-Proxy")
		ctx.Next()
	}
}
