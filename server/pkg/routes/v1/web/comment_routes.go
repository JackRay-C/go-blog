package web

import (
	"blog/pkg/api/v1/web"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitWebCommentRouter(group *gin.RouterGroup)  {
	commentGroup := group.Group("comments")
	{
		comment := web.NewComment()
		commentGroup.GET("", wrapper.Wrapper(comment.List))
		commentGroup.GET("/:id",wrapper.Wrapper(comment.Get))
	}
}