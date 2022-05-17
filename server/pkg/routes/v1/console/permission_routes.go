package console

import (

	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitPermissionRouter(group *gin.RouterGroup) {
	permissionGroup := group.Group("permissions")
	{
		permission := console.NewPermission()
		permissionGroup.GET("", wrapper.Wrapper(permission.List))
		permissionGroup.POST("", wrapper.Wrapper(permission.Post))
		permissionGroup.GET("/:id", wrapper.Wrapper(permission.Get))
		permissionGroup.PUT("/:id", wrapper.Wrapper(permission.Put))
		permissionGroup.DELETE("/:id", wrapper.Wrapper(permission.Delete))
	}
}
