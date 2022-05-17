package web

import (
	"blog/pkg/api/v1/web"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitWebSubjectRouter(group *gin.RouterGroup)  {
	subjectGroup := group.Group("subjects")
	{
		subject := web.NewSubject()
		subjectGroup.GET("", wrapper.Wrapper(subject.List))
		subjectGroup.GET("/:id",wrapper.Wrapper(subject.Get))
	}
}