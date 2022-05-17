package web

import (
	"blog/pkg/api/v1/web"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitWebHeadRouter(group *gin.RouterGroup) {
	headGroup := group.Group("heads")
	{

		head := web.NewHead()
		headGroup.GET("/:id", wrapper.Wrapper(head.Get))
		headGroup.GET("", wrapper.Wrapper(head.List))
	}
}
