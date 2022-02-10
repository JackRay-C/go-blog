package v1

import (
	v1 "blog/app/api/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitHeadsRouter(group *gin.RouterGroup) {
	headGroup := group.Group("heads")
	{
		head := v1.NewHead()
		headGroup.GET("", wrapper.Wrapper(head.List))
		headGroup.POST("", wrapper.Wrapper(head.Post))
		headGroup.GET("/:id", wrapper.Wrapper(head.Get))
		headGroup.PUT("/:id", wrapper.Wrapper(head.Put))
		headGroup.DELETE("/:id", wrapper.Wrapper(head.Delete))
	}
}
