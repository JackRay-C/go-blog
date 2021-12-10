package routes

import (
	v1 "blog/app/api/v1"
	"blog/app/api/web"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("users")
	{
		user := v1.NewUser()
		userGroup.GET("", Wrapper(user.List))
		userGroup.POST("", Wrapper(user.Post))
		userGroup.PATCH("/:id", Wrapper(user.Patch))
		userGroup.DELETE("/:id", Wrapper(user.Delete))


		userGroup.GET("/:id/roles", Wrapper(user.ListRole))
		userGroup.POST("/:id/roles", Wrapper(user.PostRole))
		userGroup.PUT("/:id/roles", Wrapper(user.PutRole))


		userGroup.GET("/:id/menus", Wrapper(user.ListMenus))

		userGroup.GET("/:id/files", Wrapper(user.ListFiles))

	}
}

func InitPublicUserRouter(group *gin.RouterGroup)  {
	userGroup := group.Group("users")
	{
		user := web.NewUser()
		userGroup.GET("/:id", Wrapper(user.Get))

		userGroup.GET("/:id/posts", Wrapper(user.ListPosts))

		userGroup.GET("/:id/subjects", Wrapper(user.ListSubjects))
	}
}