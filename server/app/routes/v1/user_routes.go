package v1

import (
	v1 "blog/app/api/console/v1"
	v12 "blog/app/api/web/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("users")
	{
		user := v1.NewUser()
		userGroup.GET("", wrapper.Wrapper(user.List))
		userGroup.POST("", wrapper.Wrapper(user.Post))
		userGroup.PUT("/:id", wrapper.Wrapper(user.Put))
		userGroup.DELETE("/:id", wrapper.Wrapper(user.Delete))

		userRole := v1.NewUserRole()
		userGroup.GET("/:id/roles", wrapper.Wrapper(userRole.Get))
		userGroup.PUT("/:id/roles", wrapper.Wrapper(userRole.Put))
	}
}

func InitPublicUserRouter(group *gin.RouterGroup)  {
	userGroup := group.Group("users")
	{
		user := v12.NewUser()
		userGroup.GET("/:id", wrapper.Wrapper(user.Get))
		userGroup.GET("/info", wrapper.Wrapper(user.GetUserInfo))

	}
}