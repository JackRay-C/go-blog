package v1

import (
	v1 "blog/app/api/console/v1"
	web "blog/app/api/web/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(group *gin.RouterGroup) {
	commentGroup := group.Group("comments")
	{
		comment := v1.NewComment()
		commentGroup.GET("", wrapper.Wrapper(comment.List))
		commentGroup.POST("", wrapper.Wrapper(comment.Post))
		commentGroup.GET("/:id", wrapper.Wrapper(comment.Get))
		commentGroup.DELETE("/:id", wrapper.Wrapper(comment.Delete))
	}
}

func InitPublicCommentRouter(group *gin.RouterGroup)  {
	commentGroup := group.Group("comments")
	{
		comment := web.NewComment()
		commentGroup.GET("", wrapper.Wrapper(comment.List))
		commentGroup.GET("/:id",wrapper.Wrapper(comment.Get))
	}
}