package v1

import (
	v1 "blog/app/api/console/v1"
	"blog/app/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitRolesRouter(group *gin.RouterGroup) {
	roleGroup := group.Group("roles")
	{
		roles := v1.NewRole()
		roleGroup.GET("", wrapper.Wrapper(roles.List))
		roleGroup.POST("", wrapper.Wrapper(roles.Post))
		roleGroup.GET("/:id", wrapper.Wrapper(roles.Get))
		roleGroup.PUT("/:id", wrapper.Wrapper(roles.Put))
		roleGroup.DELETE("/:id", wrapper.Wrapper(roles.Delete))

		rolePermission := v1.NewRolePermission()
		roleGroup.GET("/:id/permissions", wrapper.Wrapper(rolePermission.Get))
		roleGroup.PUT("/:id/permissions", wrapper.Wrapper(rolePermission.Put))

	}
}
