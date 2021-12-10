package routes

import (
	"blog/app/api/web"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	authRouter := Router.Group("auth")
	{
		auth := web.NewAuth()
		authRouter.POST("/login", Wrapper(auth.Login)) // 登录
		authRouter.GET("/captcha", Wrapper(auth.Captcha)) // 验证码
		authRouter.POST("/register", Wrapper(auth.Register)) // 注册
	}

	return authRouter
}
