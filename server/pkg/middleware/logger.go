package middleware

import (
	"blog/pkg/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//log, _ := logger.NewZapLogger(&config.App{AppLogType: logger.SimpleLog, Logs: &config.Logs{Simple: &config.Simple{
		//	Level:        "info",
		//	LogInConsole: false,
		//	Directory:    path.Join(global.App.AppHomePath, "logs"),
		//	FileName:     "access_" + time.Now().Format("2006-01-02"),
		//	LogMaxSize:   100000,
		//	LogMaxAge:    30,
		//	Format:       "json",
		//}}})
		log :=  global.Log
		t := time.Now()
		requestId, _ := c.Get(global.RequestIDKey)
		c.Next()

		context := c.Copy()
		go func() {
			// 打印日志请求
			uri, clientIP, method := context.Request.RequestURI, context.ClientIP(), context.Request.Method
			latency := time.Since(t)
			status := context.Writer.Status()

			message := fmt.Sprintf("%d | %-10s | %15s | %-7s | %d | %s ", status, latency, clientIP, method, requestId, uri)

			log.Info(message)
		}()
	}
}
