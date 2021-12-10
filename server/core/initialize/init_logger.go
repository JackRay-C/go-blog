package initialize

import (
	"blog/core/global"
	"blog/core/logger"
	"log"
)

/**
初始化日志
*/
func SetupLogger() {
	log.Println("初始化日志组件...")
	if global.Setting.Zap == nil {
		global.Logger = logger.NewSimpleLogger()
	}
	global.Logger = logger.NewZapLogger(global.Setting.Zap)

	log.Printf("初始化日志组件完成\n")
}
