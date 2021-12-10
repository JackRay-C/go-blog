package middleware

import (
	"blog/core/global"
	"github.com/gin-gonic/gin"
)

// 为每一个请求生成一个唯一ID
func RequestID() gin.HandlerFunc  {
	return func(c *gin.Context) {
		requestId, _ := global.Snowflake.NextID()
		c.Set("requestId", requestId)
		c.Next()
	}
}
