package routes

import (
	v1 "blog/app/api/v1"
	v2 "blog/app/api/v2"
	"blog/app/api/web"

	"github.com/gin-gonic/gin"
)

func InitPostRouter(group *gin.RouterGroup) {
	postGroup := group.Group("posts")
	{
		post := v1.NewPost()

		postGroup.GET("", Wrapper(post.List))
		postGroup.POST("", Wrapper(post.Post))
		postGroup.GET("/:id", Wrapper(post.Get))
		postGroup.PUT("/:id", Wrapper(post.Put))
		postGroup.DELETE("/:id", Wrapper(post.Delete))
	}
}

func InitV2PostRouter(group *gin.RouterGroup)  {
	postGroup := group.Group("posts")
	{
		post := v2.NewPost()
		postGroup.POST("",Wrapper(post.Initialize))
		postGroup.GET("/id", Wrapper(post.Pull))
		postGroup.PUT("/id", Wrapper(post.Commit))
		postGroup.POST("/:id/publish", Wrapper(post.Publish))
	}
}

func InitPublicPostRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("posts")
	{
		post := web.NewPost()
		routerGroup.GET("", Wrapper(post.List))
		routerGroup.GET("/:id", Wrapper(post.Get))
		routerGroup.POST("/:id/like", Wrapper(post.Like))

	}
}
