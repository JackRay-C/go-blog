package main

import (
	"blog/core/global"
	"blog/core/initialize"
	"blog/core/viper"
	"embed"
	_ "embed"
	"flag"

	"log"

	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download




// @title 博客系统
// @version 1.0
// @description Go blog
func main() {
	log.Println("启动服务器...")
	startTime := time.Now()

	setupConfig()
	initialize.SetupLogger()
	initialize.SetupDatabase()
	initialize.SetupSnowflake()
	initialize.SetupSMTP()
	initialize.SetupStorage()

	endTime := time.Since(startTime)
	log.Printf("启动耗时： %s", endTime)
	initialize.StartServer()
}


//go:embed configs
var configs embed.FS

func setupConfig() {
	log.Println("初始化配置文件...")
	var configFile string
	flag.StringVar(&configFile, "c", "", "config file path.")
	flag.Parse()

	if configFile == "" {
		file, err := configs.ReadFile("configs/config.yaml")
		if err != nil {
			panic(err)
		}
		global.Viper = viper.ViperEmbed(file)
	} else {
		global.Viper = viper.Viper(configFile)
	}
	log.Println("初始化配置文件完成")
}
