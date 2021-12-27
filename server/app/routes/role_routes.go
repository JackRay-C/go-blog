package routes

import (
	v1 "blog/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRolesRouter(group *gin.RouterGroup) {
	roleGroup := group.Group("roles")
	{
		roles := v1.NewRole()
		roleGroup.GET("", Wrapper(roles.List))
		roleGroup.POST("", Wrapper(roles.Post))
		roleGroup.GET("/:id", Wrapper(roles.Get))
		roleGroup.PUT("/:id", Wrapper(roles.Put))
		roleGroup.DELETE("/:id", Wrapper(roles.Delete))

		roleMenu := v1.NewRoleMenu()
		roleGroup.GET("/:id/menus", Wrapper(roleMenu.Get))
		roleGroup.PUT("/:id/menus", Wrapper(roleMenu.Put))

		rolePermission := v1.NewRolePermission()
		roleGroup.GET("/:id/permission", Wrapper(rolePermission.Get))
		roleGroup.PUT("/:id/permission", Wrapper(rolePermission.Put))

	}
}
