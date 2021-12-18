package routes

import (
	v1 "blog/app/api/v1"
	"blog/app/api/web"
	"github.com/gin-gonic/gin"
)

func InitSubjectRouter(group *gin.RouterGroup) {
	subjectGroup := group.Group("subjects")
	{
		subject := v1.NewSubject()

		subjectGroup.GET("", Wrapper(subject.List))
		subjectGroup.GET("/:id", Wrapper(subject.Get))
		subjectGroup.POST("", Wrapper(subject.Post))
		subjectGroup.PUT("/:id", Wrapper(subject.Put))
		subjectGroup.DELETE("/:id", Wrapper(subject.Delete))
	}
}

func InitPublicSubjectRouter(group *gin.RouterGroup) {
	subjectGroup := group.Group("subjects")
	{
		subject := web.NewSubject()
		subjectGroup.GET("", Wrapper(subject.List))
		subjectGroup.GET("/:id", Wrapper(subject.Get))

		//subjectGroup.GET("/:id/posts", Wrapper(subject.GetPosts))
	}
}
