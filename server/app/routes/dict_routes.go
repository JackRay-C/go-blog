package routes

import (
	v1 "blog/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitDictRouter(group *gin.RouterGroup) {
	dictGroup := group.Group("dicts")
	{
		dict := v1.NewDict()

		dictGroup.POST("", Wrapper(dict.Post))
		dictGroup.DELETE("/:id", Wrapper(dict.Delete))
		dictGroup.PUT("/:id", Wrapper(dict.Put))
	}
}

func InitPublicDictRouter(group *gin.RouterGroup) {
	dictGroup := group.Group("dicts")
	{
		dict := v1.NewDict()
		dictGroup.GET("", Wrapper(dict.List))
		dictGroup.GET("/:id", Wrapper(dict.Get))

	}
}

