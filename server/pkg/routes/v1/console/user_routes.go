package console

import (
	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("users")
	{
		user := console.NewUser()
		userGroup.GET("", wrapper.Wrapper(user.List))
		userGroup.POST("", wrapper.Wrapper(user.Post))
		userGroup.PUT("/:id", wrapper.Wrapper(user.Put))
		userGroup.DELETE("/:id", wrapper.Wrapper(user.Delete))
		userGroup.GET("/info", wrapper.Wrapper(user.GetUserInfo))
		userGroup.GET("/:id", wrapper.Wrapper(user.Get))

		userRole := console.NewUserRole()
		userGroup.GET("/:id/roles", wrapper.Wrapper(userRole.Get))
		userGroup.PUT("/:id/roles", wrapper.Wrapper(userRole.Put))
	}
}

