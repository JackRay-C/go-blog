package v1

import (
	v1 "blog/app/api/v1"
	"blog/app/api/web"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitTagRouter(group *gin.RouterGroup) {
	tagGroup := group.Group("tags")
	{
		tag := v1.NewTag()

		tagGroup.GET("", wrapper.Wrapper(tag.List))
		tagGroup.GET("/:id", wrapper.Wrapper(tag.Get))
		tagGroup.POST("", wrapper.Wrapper(tag.Post))
		tagGroup.DELETE("/:id", wrapper.Wrapper(tag.Delete))
		tagGroup.PUT("/:id", wrapper.Wrapper(tag.Put))
	}
}

func InitPublicTagRouter(group *gin.RouterGroup) {
	tagGroup := group.Group("tags")
	{
		tag := web.NewTag()

		tagGroup.GET("", wrapper.Wrapper(tag.List))
		tagGroup.GET("/:id", wrapper.Wrapper(tag.Get))
	}
}
