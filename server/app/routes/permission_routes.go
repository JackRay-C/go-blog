package routes

import (
	v1 "blog/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitPermissionRouter(group *gin.RouterGroup) {
	permissionGroup := group.Group("permissions")
	{
		permission := v1.NewPermission()
		permissionGroup.GET("", Wrapper(permission.List))
		permissionGroup.POST("", Wrapper(permission.Post))
		permissionGroup.GET("/:id", Wrapper(permission.Get))
		permissionGroup.PUT("/:id", Wrapper(permission.Put))
		permissionGroup.DELETE("/:id", Wrapper(permission.Delete))
	}
}
