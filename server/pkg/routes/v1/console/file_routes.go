package console

import (
	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(group *gin.RouterGroup) {
	fileGroup := group.Group("files")
	{
		image := console.NewFile()
		fileGroup.GET("", wrapper.Wrapper(image.List))
		fileGroup.GET("/:id", wrapper.Wrapper(image.Get))
		fileGroup.POST("", wrapper.Wrapper(image.Post))
		fileGroup.DELETE("/:id",wrapper.Wrapper(image.Delete))
	}
}

