package v1

import (
	v1 "blog/app/api/console/v1"
	v12 "blog/app/api/web/v1"

	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitSubjectRouter(group *gin.RouterGroup) {
	subjectGroup := group.Group("subjects")
	{
		subject := v1.NewSubject()

		subjectGroup.GET("", wrapper.Wrapper(subject.List))
		subjectGroup.GET("/:id", wrapper.Wrapper(subject.Get))
		subjectGroup.POST("", wrapper.Wrapper(subject.Post))
		subjectGroup.PUT("/:id", wrapper.Wrapper(subject.Put))
		subjectGroup.DELETE("/:id", wrapper.Wrapper(subject.Delete))
	}
}

func InitPublicSubjectRouter(group *gin.RouterGroup) {
	subjectGroup := group.Group("subjects")
	{
		subject := v12.NewSubject()
		subjectGroup.GET("", wrapper.Wrapper(subject.List))
		subjectGroup.GET("/:id", wrapper.Wrapper(subject.Get))

		//subjectGroup.GET("/:id/posts", Wrapper(subject.GetPosts))
	}
}
