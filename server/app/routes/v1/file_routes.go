package v1

import (
	v1 "blog/app/api/console/v1"
	web "blog/app/api/web/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(group *gin.RouterGroup) {
	fileGroup := group.Group("files")
	{
		image := v1.NewFile()
		fileGroup.GET("", wrapper.Wrapper(image.List))
		fileGroup.GET("/:id", wrapper.Wrapper(image.Get))
		fileGroup.POST("", wrapper.Wrapper(image.Post))
		fileGroup.DELETE("/:id",wrapper.Wrapper(image.Delete))
	}
}

func InitPublicFileRouter(group *gin.RouterGroup) {
	fileGroup := group.Group("files")
	{
		image := web.NewFile()

		fileGroup.GET("/:id", wrapper.Wrapper(image.Get))

	}
}
