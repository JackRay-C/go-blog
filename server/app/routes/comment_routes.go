package routes

import (
	v1 "blog/app/api/v1"
	"blog/app/api/web"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(group *gin.RouterGroup) {
	commentGroup := group.Group("comments")
	{
		comment := v1.NewComment()
		commentGroup.GET("", Wrapper(comment.List))
		commentGroup.GET("/:id", Wrapper(comment.Get))
		commentGroup.DELETE("/:id", Wrapper(comment.Delete))
	}
}

func InitPublicCommentRouter(group *gin.RouterGroup)  {
	commentGroup := group.Group("comments")
	{
		comment := web.NewComment()
		commentGroup.POST("", Wrapper(comment.Post))
		commentGroup.GET("", Wrapper(comment.List))
		commentGroup.DELETE("/:id", Wrapper(comment.Delete))
	}
}