package console

import (

	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitDictRouter(group *gin.RouterGroup) {
	dictGroup := group.Group("dicts")
	{
		dict := console.NewDict()
		dictGroup.GET("", wrapper.Wrapper(dict.List))
		dictGroup.GET("/:id", wrapper.Wrapper(dict.Get))
		dictGroup.POST("", wrapper.Wrapper(dict.Post))
		dictGroup.DELETE("/:id", wrapper.Wrapper(dict.Delete))
		dictGroup.PUT("/:id", wrapper.Wrapper(dict.Put))
	}
}


