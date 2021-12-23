package routes

import (
	v1 "blog/app/api/v1"
	"blog/app/api/web"
	"github.com/gin-gonic/gin"
)

func InitTagRouter(group *gin.RouterGroup) {
	tagGroup := group.Group("tags")
	{
		tag := v1.NewTag()

		tagGroup.GET("", Wrapper(tag.List))
		tagGroup.GET("/:id", Wrapper(tag.Get))
		tagGroup.POST("", Wrapper(tag.Post))
		tagGroup.DELETE("/:id", Wrapper(tag.Delete))
		tagGroup.PUT("/:id", Wrapper(tag.Put))
	}
}

func InitPublicTagRouter(group *gin.RouterGroup) {
	tagGroup := group.Group("tags")
	{
		tag := web.NewTag()

		tagGroup.GET("", Wrapper(tag.List))
		tagGroup.GET("/:id", Wrapper(tag.Get))
	}
}
