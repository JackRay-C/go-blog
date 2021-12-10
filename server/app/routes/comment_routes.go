package routes

import (
	v1 "blog/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(group *gin.RouterGroup) {
	commentGroup := group.Group("comments")
	{
		comment := v1.NewComment()
		commentGroup.GET("", Wrapper(comment.List))
		commentGroup.POST("", Wrapper(comment.Post))
		commentGroup.GET("/:id", Wrapper(comment.Get))
		commentGroup.PUT("/:id", Wrapper(comment.Put))
		commentGroup.PATCH("/:id", Wrapper(comment.Patch))
		commentGroup.DELETE("/:id", Wrapper(comment.Delete))
	}
}
