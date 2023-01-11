package web

import (
	"blog/pkg/api/v1/web"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitWebTagRouter(group *gin.RouterGroup)  {
	tagGroup := group.Group("tags")
	{
		tag := web.NewTag()
		tagGroup.GET("", wrapper.Wrapper(tag.List))
		tagGroup.GET("/:id",wrapper.Wrapper(tag.Get))
		tagGroup.GET("/:id/posts", wrapper.Wrapper(tag.Posts))
	}
}