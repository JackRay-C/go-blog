package v1

import (
	v1 "blog/app/api/console/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitRepositoryRouter(group *gin.RouterGroup) {
	repositoryGroup := group.Group("repository")
	{
		repository := v1.NewRepository()
		repositoryGroup.GET("", wrapper.Wrapper(repository.List))
		//headGroup.POST("", wrapper.Wrapper(head.Post))
		//headGroup.GET("/:id", wrapper.Wrapper(head.Get))
		//headGroup.PUT("/:id", wrapper.Wrapper(head.Put))
		//headGroup.DELETE("/:id", wrapper.Wrapper(head.Delete))
	}
}