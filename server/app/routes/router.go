package routes

import (
	"blog/app/middleware"
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
	r.Use(middleware.Cors())
	r.Use(middleware.AppInfo())   // 元数据信息
	r.Use(middleware.RequestID()) // 为每个请求生成一个唯一的请求ID
 	r.Use(middleware.Logger())    // 请求完成后记录日志
	r.Use(middleware.Recovery())  // 全局异常处理
	r.Use(middleware.Authentication()) // 根据token判断是否登录


	publicGroup := r.Group("/api/v1")
	{
		InitAuthRouter(publicGroup)
		InitPublicPostRouter(publicGroup)
		InitPublicSubjectRouter(publicGroup)
		InitPublicUserRouter(publicGroup)
		InitPublicTagRouter(publicGroup)
		InitPublicFileRouter(publicGroup)
		InitPublicDictRouter(publicGroup)
	}
	privateGroup := r.Group("/api/v1/admin/")
	privateGroup.Use(middleware.Permission()) // 获取用户角色及权限列表
	{
		InitUserRouter(privateGroup)
		InitPostRouter(privateGroup)
		InitTagRouter(privateGroup)
		InitSubjectRouter(privateGroup)
		InitFileRouter(privateGroup)
		InitRolesRouter(privateGroup)
		InitMenuRouter(privateGroup)
		InitDictRouter(privateGroup)
		InitCommentRouter(privateGroup)
	}

	return r
}
