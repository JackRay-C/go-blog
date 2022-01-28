package v1

import (
	v1 "blog/app/api/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitDictRouter(group *gin.RouterGroup) {
	dictGroup := group.Group("dicts")
	{
		dict := v1.NewDict()

		dictGroup.POST("", wrapper.Wrapper(dict.Post))
		dictGroup.DELETE("/:id", wrapper.Wrapper(dict.Delete))
		dictGroup.PUT("/:id", wrapper.Wrapper(dict.Put))
	}
}

func InitPublicDictRouter(group *gin.RouterGroup) {
	dictGroup := group.Group("dicts")
	{
		dict := v1.NewDict()
		dictGroup.GET("", wrapper.Wrapper(dict.List))
		dictGroup.GET("/:id", wrapper.Wrapper(dict.Get))

	}
}

