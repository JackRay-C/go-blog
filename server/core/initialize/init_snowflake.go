package initialize

import (
	"blog/core/global"
	"blog/core/snowflake"
	"fmt"
	"log"
)

/**
初始化全局ID生成器
*/
func SetupSnowflake() {
	log.Println("初始化雪花ID生成器...")
	var err error
	global.Snowflake, err = snowflake.New(global.Setting.Snowflake)
	if err != nil {
		panic(fmt.Sprintf("初始化雪花ID生成器错误：%s\n", err))
	}
	log.Println("初始化雪花ID生成器完成.")
}
