package web

import (
	"blog/pkg/api/v1/console"
	"blog/pkg/utils/wrapper"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	authRouter := Router.Group("auth")
	{
		auth := console.NewAuth()
		authRouter.POST("/login", wrapper.Wrapper(auth.Login))            // 登录
		authRouter.POST("/register", wrapper.Wrapper(auth.Register))      // 注册
		authRouter.GET("/refresh", wrapper.Wrapper(auth.RefreshToken))    // 刷新token
		authRouter.GET("/info", wrapper.Wrapper(auth.Info))               // 获取当前用户信息
		authRouter.GET("/roles", wrapper.Wrapper(auth.Roles))             // 获取当前用户角色信息
		authRouter.GET("/permissions", wrapper.Wrapper(auth.Permissions)) // 获取所有权限
		authRouter.GET("/captcha", wrapper.Wrapper(auth.Captcha))         // 验证码
	}

	return authRouter
}
