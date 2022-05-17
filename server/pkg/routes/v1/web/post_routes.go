package web

import (
	"blog/pkg/api/v1/web"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitWebPostsRoutes(group *gin.RouterGroup)  {

	postGroup := group.Group("posts")
	{
		post := web.NewPost()

		postGroup.GET("", wrapper.Wrapper(post.List))
		postGroup.GET("/:id", wrapper.Wrapper(post.Get))
	}

}