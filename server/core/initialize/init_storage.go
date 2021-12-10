package initialize

import (
	"blog/core/global"
	aliyun_oss "blog/core/storage/aliyun-oss"
	"blog/core/storage/local"
	"blog/core/storage/qiniu"
	tencent_oss "blog/core/storage/tencent-oss"
	"fmt"
	"log"
	"strings"
)

/**
初始化文件存储
*/
func SetupStorage() {
	log.Println("初始化存储... ")
	switch {
	case strings.ToUpper(global.Setting.App.StorageType) == "LOCAL":
		global.Storage = local.NewLocalStorage(global.Setting.App, global.Setting.Local)
	case strings.ToUpper(global.Setting.App.StorageType) == "QINIU":
		global.Storage = &qiniu.Qiniu{
			App:   global.Setting.App,
			Qiniu: global.Setting.Qiniu,
		}
	case strings.ToUpper(global.Setting.App.StorageType) == "ALIYUN-OSS":
		fmt.Println(global.Setting.App)
		fmt.Println(global.Setting.AliyunOSS)
		global.Storage = aliyun_oss.NewAliyunOss(global.Setting.App, global.Setting.AliyunOSS)
	case strings.ToUpper(global.Setting.App.StorageType) == "TENCENT-OSS":
		global.Storage = &tencent_oss.TencentOss{
			App:        global.Setting.App,
			TencentOSS: global.Setting.TencentOSS,
		}
	default:
		global.Storage = local.NewLocalStorage(global.Setting.App, global.Setting.Local)
	}
	log.Printf("初始化%s存储完成.\n", global.Setting.App.StorageType)
}
