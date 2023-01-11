package console

import (

	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitTagRouter(group *gin.RouterGroup) {
	tagGroup := group.Group("tags")
	{
		tag := console.NewTag()

		tagGroup.GET("", wrapper.Wrapper(tag.List))
		tagGroup.GET("/:id", wrapper.Wrapper(tag.Get))
		tagGroup.POST("", wrapper.Wrapper(tag.Post))
		tagGroup.DELETE("/:id", wrapper.Wrapper(tag.Delete))
		tagGroup.PUT("/:id", wrapper.Wrapper(tag.Put))


		tagGroup.GET("/:id/posts", wrapper.Wrapper(tag.ListPosts))
	}
}

