package console

import (
	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(group *gin.RouterGroup) {
	commentGroup := group.Group("comments")
	{
		comment := console.NewComment()
		commentGroup.GET("", wrapper.Wrapper(comment.List))
		//commentGroup.POST("", wrapper.Wrapper(comment.Post))
		commentGroup.GET("/:id", wrapper.Wrapper(comment.Get))
		commentGroup.DELETE("/:id", wrapper.Wrapper(comment.Delete))
	}
}