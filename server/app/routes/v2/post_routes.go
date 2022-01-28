package v2

import (
	v2 "blog/app/api/v2"
	"blog/app/api/web"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)


func InitV2PostRouter(group *gin.RouterGroup)  {
	postGroup := group.Group("posts")
	{
		post := v2.NewPost()
		postGroup.POST("",wrapper.Wrapper(post.Initialize))
		postGroup.GET("/id", wrapper.Wrapper(post.Pull))
		postGroup.PUT("/id", wrapper.Wrapper(post.Commit))
		postGroup.POST("/:id/publish", wrapper.Wrapper(post.Publish))
	}
}

func InitPublicPostRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("posts")
	{
		post := web.NewPost()
		routerGroup.GET("", wrapper.Wrapper(post.List))
		routerGroup.GET("/:id", wrapper.Wrapper(post.Get))
		routerGroup.POST("/:id/like", wrapper.Wrapper(post.Like))

	}
}
