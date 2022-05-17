package web

import (
	"blog/pkg/api/v1/web"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitWebFileRouter(group *gin.RouterGroup) {
	fileGroup := group.Group("files")
	{
		image := web.NewFile()

		fileGroup.GET("/:id", wrapper.Wrapper(image.Get))

	}
}
