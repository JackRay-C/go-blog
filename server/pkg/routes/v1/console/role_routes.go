package console

import (

	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitRolesRouter(group *gin.RouterGroup) {
	roleGroup := group.Group("roles")
	{
		roles := console.NewRole()
		roleGroup.GET("", wrapper.Wrapper(roles.List))
		roleGroup.POST("", wrapper.Wrapper(roles.Post))
		roleGroup.GET("/:id", wrapper.Wrapper(roles.Get))
		roleGroup.PUT("/:id", wrapper.Wrapper(roles.Put))
		roleGroup.DELETE("/:id", wrapper.Wrapper(roles.Delete))

		rolePermission := console.NewRolePermission()
		roleGroup.GET("/:id/permissions", wrapper.Wrapper(rolePermission.Get))
		roleGroup.PUT("/:id/permissions", wrapper.Wrapper(rolePermission.Put))

	}
}
