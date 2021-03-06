package console

import (

	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitHeadsRouter(group *gin.RouterGroup) {
	headGroup := group.Group("heads")
	{
		head := console.NewHead()
		headGroup.GET("", wrapper.Wrapper(head.List))
		headGroup.POST("", wrapper.Wrapper(head.Post))
		headGroup.GET("/:id", wrapper.Wrapper(head.Get))
		headGroup.PUT("/:id", wrapper.Wrapper(head.Put))
		headGroup.DELETE("/:id", wrapper.Wrapper(head.Delete))
	}
}
