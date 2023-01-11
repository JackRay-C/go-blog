package console

import (
	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitCurrentRoutes(group *gin.RouterGroup)  {
	routerGroup := group.Group("/current")
	{
		user := console.NewUser()
		routerGroup.GET("/info", wrapper.Wrapper(user.Get))
		routerGroup.GET("/roles", wrapper.Wrapper(user.Get))
		routerGroup.GET("/permissions", wrapper.Wrapper(user.Get))
		routerGroup.GET("/posts", wrapper.Wrapper(user.Post))
	}

}
