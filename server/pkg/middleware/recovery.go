package middleware

import (

	"blog/pkg/model/vo"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc  {
	return func(c *gin.Context) {
		defer func() {
			if err:= recover(); err != nil {
				log.Printf("panic %s", err)
				debug.PrintStack()
				_ = c.Error(vo.InternalServerError.SetMsg("%s", err))
				c.Abort()
			}
		}()
		c.Next()
	}
}