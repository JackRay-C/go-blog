package middleware

import (
	"blog/pkg/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := global.Log
		t := time.Now()
		c.Next()

		context := c.Copy()
		go func() {
			// 打印日志请求
			uri, clientIP, method := context.Request.RequestURI, context.ClientIP(), context.Request.Method
			latency := time.Since(t)
			status := context.Writer.Status()

			message := fmt.Sprintf("%d | %-10s | %15s | %-7s %s ", status, latency, clientIP, method, uri)

			log.Info(message)
		}()
	}
}
