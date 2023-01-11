package console

import (
	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitPostRouter(group *gin.RouterGroup) {
	postGroup := group.Group("posts")
	{
		post := console.NewPost()

		postGroup.GET("", wrapper.Wrapper(post.List))
		postGroup.POST("", wrapper.Wrapper(post.Post))
		postGroup.GET("/:id",wrapper.Wrapper(post.Get))
		postGroup.PUT("/:id", wrapper.Wrapper(post.Put))
		postGroup.DELETE("/:id", wrapper.Wrapper(post.Delete))


		postGroup.GET("/:id/tags", wrapper.Wrapper(post.ListTags))
	}
}

