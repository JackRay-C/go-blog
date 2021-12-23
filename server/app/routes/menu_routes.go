package routes

import (
	v1 "blog/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(group *gin.RouterGroup) {
	menuGroup := group.Group("menus")
	{
		menus := v1.NewMenu()
		menuGroup.GET("", Wrapper(menus.List))
		menuGroup.POST("", Wrapper(menus.Post))
		menuGroup.GET("/:id", Wrapper(menus.Get))
		menuGroup.PUT("/:id", Wrapper(menus.Put))
		menuGroup.DELETE("/:id", Wrapper(menus.Delete))
	}
}
