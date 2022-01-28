package routes

import (
	"blog/app/middleware"
	"blog/app/routes/v1"
	"blog/app/routes/v2"
	"blog/core/setting"
	_ "embed"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouters(setting *setting.App) *gin.Engine {

	if setting.RunMode == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if !setting.LogColorConsole {
		gin.DisableConsoleColor() // 禁用控制台打印日志颜色
	}

	r := gin.New()

	if gin.Mode() == gin.DebugMode {
		pprof.Register(r)
	}

	// 静态文件服务器
	r.StaticFS("/static", http.Dir(setting.StaticPath))

	// 公共中间件
	r.Use(middleware.Cors())           // 跨域设置
	r.Use(middleware.AppInfo())        // 元数据信息
	r.Use(middleware.RequestID())      // 为每个请求生成一个唯一的请求ID
	r.Use(middleware.Logger())         // 请求完成后记录日志
	r.Use(middleware.Recovery())       // 全局异常处理
	r.Use(middleware.Authentication()) // 根据token判断是否登录

	webGroup := r.Group("/api/v1")
	{
		v1.InitAuthRouter(webGroup)
		v1.InitPublicPostRouter(webGroup)
		v1.InitPublicSubjectRouter(webGroup)
		v1.InitPublicUserRouter(webGroup)
		v1.InitPublicTagRouter(webGroup)
		v1.InitPublicFileRouter(webGroup)
		v1.InitPublicDictRouter(webGroup)
		v1.InitPublicCommentRouter(webGroup)
	}

	consoleGroup := r.Group("/api/v1/admin/")
	consoleGroup.Use(middleware.Permission()) // 认证及鉴权中间件
	{
		v1.InitUserRouter(consoleGroup)
		v1.InitPostRouter(consoleGroup)
		v1.InitTagRouter(consoleGroup)
		v1.InitSubjectRouter(consoleGroup)
		v1.InitFileRouter(consoleGroup)
		v1.InitRolesRouter(consoleGroup)
		v1.InitDictRouter(consoleGroup)
		v1.InitCommentRouter(consoleGroup)
		v1.InitPermissionRouter(consoleGroup)
	}

	consoleGroupV2 := r.Group("/api/v2/admin")
	consoleGroupV2.Use(middleware.Permission())
	{
		v2.InitV2PostRouter(consoleGroupV2)
	}
	return r
}
