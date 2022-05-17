package web

import (
	"blog/pkg/api/v1/web"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitWebDictRouter(group *gin.RouterGroup)  {
	commentGroup := group.Group("dicts")
	{
		dict := web.NewDict()
		commentGroup.GET("", wrapper.Wrapper(dict.List))
		commentGroup.GET("/:id",wrapper.Wrapper(dict.Get))
	}
}
