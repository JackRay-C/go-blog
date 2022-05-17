package console

import (
	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitSubjectRouter(group *gin.RouterGroup) {
	subjectGroup := group.Group("subjects")
	{
		subject := console.NewSubject()

		subjectGroup.GET("", wrapper.Wrapper(subject.List))
		subjectGroup.GET("/:id", wrapper.Wrapper(subject.Get))
		subjectGroup.POST("", wrapper.Wrapper(subject.Post))
		subjectGroup.PUT("/:id", wrapper.Wrapper(subject.Put))
		subjectGroup.DELETE("/:id", wrapper.Wrapper(subject.Delete))
	}
}

