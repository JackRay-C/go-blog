package routes

import (

	"blog/internal/config"
	"blog/pkg/middleware"
	"blog/pkg/routes/v1/console"
	"blog/pkg/routes/v1/web"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

const (
	PRODUCTION = "production"
	DEVELOPMENT = "development"
)

func NewRouters(setting *config.App) *gin.Engine {
	if setting.AppMode == PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	} else if setting.AppMode == DEVELOPMENT {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	if gin.Mode() == gin.DebugMode {
		pprof.Register(r)
	}

	// 静态文件服务器
	r.StaticFS("/static", http.Dir(path.Join(setting.AppHomePath, setting.Server.StaticRootPath)))
	//r.Static("/", path.Join(setting.AppHomePath, setting.Server.StaticRootPath))



	// 中间件
	r.Use(middleware.Cors())
	r.Use(middleware.AppInfo())
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())

	consoleGroup := r.Group("/api/v1/console")
	consoleGroup.Use(middleware.Permission())
	{
		console.InitCommentRouter(consoleGroup)
		console.InitDictRouter(consoleGroup)
		console.InitDraftRouter(consoleGroup)
		console.InitPostRouter(consoleGroup)
		console.InitSubjectRouter(consoleGroup)
		console.InitFileRouter(consoleGroup)
		console.InitPermissionRouter(consoleGroup)
		console.InitUserRouter(consoleGroup)
		console.InitTagRouter(consoleGroup)
		console.InitRolesRouter(consoleGroup)
	}

	webGroup := r.Group("/api/v1/")
	{
		web.InitAuthRouter(webGroup)
		web.InitWebCommentRouter(webGroup)
		web.InitWebDictRouter(webGroup)
		web.InitWebFileRouter(webGroup)
		web.InitWebPostsRoutes(webGroup)
		web.InitWebSubjectRouter(webGroup)
		web.InitWebTagRouter(webGroup)
		web.InitWebUserRouter(webGroup)

	}

	return r
}

