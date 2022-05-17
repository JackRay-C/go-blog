package middleware

import (
	"blog/pkg/global"
	"github.com/gin-gonic/gin"
)

func AppInfo() gin.HandlerFunc  {
	return func(c *gin.Context) {
		c.Set("app_name", global.App.AppName)
		c.Set("app_version", global.App.AppVersion)
		c.Next()
	}
}
