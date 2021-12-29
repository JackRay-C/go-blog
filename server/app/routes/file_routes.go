package routes

import (
	v1 "blog/app/api/v1"
	"blog/app/api/web"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(group *gin.RouterGroup) {
	fileGroup := group.Group("files")
	{
		image := v1.NewFile()
		fileGroup.GET("", Wrapper(image.List))
		fileGroup.GET("/:id", Wrapper(image.Get))
		fileGroup.POST("", Wrapper(image.Post))
		fileGroup.DELETE("/:id", Wrapper(image.Delete))
	}
}

func InitPublicFileRouter(group *gin.RouterGroup) {
	fileGroup := group.Group("files")
	{
		image := web.NewFile()

		fileGroup.GET("/:id", Wrapper(image.Get))

	}
}
