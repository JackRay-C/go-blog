package console

import (

	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	authRouter := Router.Group("auth")
	{
		auth := console.NewAuth()
		authRouter.POST("/login", wrapper.Wrapper(auth.Login))          // 登录
		authRouter.GET("/captcha", wrapper.Wrapper(auth.Captcha))       // 验证码
		authRouter.POST("/register", wrapper.Wrapper(auth.Register))    // 注册
		authRouter.POST("/refresh", wrapper.Wrapper(auth.RefreshToken)) // 刷新token
	}

	return authRouter
}
