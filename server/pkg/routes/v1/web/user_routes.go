package web

import (
	"blog/pkg/api/v1/web"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitWebUserRouter(group *gin.RouterGroup)  {
	userGroup := group.Group("users")
	{
		user := web.NewUser()
		userGroup.GET("/:id",wrapper.Wrapper(user.Get))
	}
}