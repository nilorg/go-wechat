package middleware

import (
	"github.com/gin-gonic/gin"
)

// Header 头处理
func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Server", "Wechat-Gateway")
		c.Next()
	}
}
