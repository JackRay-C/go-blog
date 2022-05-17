package v2

import (
	v1 "blog/app/api/console/v1"
	v12 "blog/app/api/web/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)


func InitV2PostRouter(group *gin.RouterGroup)  {
	postGroup := group.Group("posts")
	{
		post := v1.NewPost()
		postGroup.GET("", wrapper.Wrapper(post.Post))
		//postGroup.GET("", wrapper.Wrapper(post.Repositories))
		//postGroup.POST("",wrapper.Wrapper(post.Initialize))
		//postGroup.GET("/id", wrapper.Wrapper(post.Pull))
		//postGroup.PUT("/id", wrapper.Wrapper(post.Commit))
		//postGroup.POST("/:id/publish", wrapper.Wrapper(post.Publish))
	}
}

func InitPublicPostRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("posts")
	{
		post := v12.NewPost()
		routerGroup.GET("", wrapper.Wrapper(post.List))
		routerGroup.GET("/:id", wrapper.Wrapper(post.Get))
		//routerGroup.POST("/:id/like", wrapper.Wrapper(post.Like))

	}
}
