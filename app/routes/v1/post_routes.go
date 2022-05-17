package v1

import (
	console "blog/app/api/console/v1"
	web "blog/app/api/web/v1"
	"blog/app/utils/wrapper"

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
	}
}

func InitPublicPostRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("posts")
	{
		post := web.NewPost()
		routerGroup.GET("", wrapper.Wrapper(post.List))
		routerGroup.GET("/:id", wrapper.Wrapper(post.Get))


	}
}
