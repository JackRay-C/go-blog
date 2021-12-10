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
		roleGroup.PATCH("/:id", Wrapper(roles.Patch))
		roleGroup.DELETE("/:id", Wrapper(roles.Delete))


		roleGroup.GET("/:id/menus", Wrapper(roles.GetMenus))
		roleGroup.POST("/:id/menus", Wrapper(roles.PostMenus))
		roleGroup.PUT("/:id/menus", Wrapper(roles.PutMenus))

	}
}
