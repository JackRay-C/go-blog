package server

import (
	"blog/internal/cache"
	"blog/internal/database"
	"blog/internal/hook"
	"blog/internal/logger"
	"blog/internal/mail"
	"blog/internal/snowflake"
	"blog/internal/storage"
	"blog/pkg/global"
	"blog/pkg/initialize"
	"blog/pkg/routes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/autotls"
	"github.com/spf13/viper"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	Https  = "https"
	Http   = "http"
	Socket = "socket"
)


func New(file string) {

	unmarshalConfig(file)

	var err error

	global.Log, err = logger.New(global.App)
	if err != nil {
		log.Fatal(err)
	}
	global.Log.Infof("初始化数据库...")
	if global.DB, err = database.New(global.App); err != nil {
		global.Log.Fatalf("初始化数据库失败: %s", err)
	}

	initialize.InitTable()

	global.Log.Infof("初始化雪花ID...")
	if global.Snowflake, err = snowflake.New(global.App); err != nil {
		global.Log.Fatalf("初始化雪花生成器失败：: %s", err)
	}

	global.Log.Infof("初始化SMTP配置...")
	global.Mail = mail.New(global.App)

	global.Log.Infof("初始化存储...")
	if global.Storage, err = storage.New(global.App); err != nil {
		global.Log.Fatalf("初始化存储失败： %s", err)
	}

	global.Log.Infof("初始化路由...")
	global.Routers = routes.NewRouters(global.App)

	global.Log.Infof("初始化缓存Redis...")

	if global.Cache, err = cache.NewRedis(global.App); err != nil {
		global.Log.Fatalf("初始化Redis失败: %s", err)
	}

	global.Log.Infof("初始化Server...")
	global.Server = &http.Server{
		Addr:           fmt.Sprintf(":%d", global.App.Server.Port),
		Handler:        global.Routers,
		ReadTimeout:    global.App.Server.ReadTimeout * time.Second,
		WriteTimeout:   global.App.Server.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func unmarshalConfig(cfg string)  {
	global.Viper = viper.New()

	global.Viper.SetConfigFile(cfg)

	if err := global.Viper.ReadInConfig(); err != nil {
		log.Printf("viper read config file failed: %s", err)
		os.Exit(1)
	}
	global.Viper.WatchConfig()
	global.Viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file changed: %s\n", e.Name)
		if err := global.Viper.Unmarshal(&global.App, viper.DecodeHook(hook.GetDecodeHookConfig())); err != nil {
			log.Fatalf("unmarshal config file failed: %s", err)
		}
	})
	if err := global.Viper.Unmarshal(&global.App, viper.DecodeHook(hook.GetDecodeHookConfig())); err != nil {
		log.Fatalf("unmarshal config file failed: %s", err)
	}
}

func Start() {
	go func() {
		if global.App.Server.Protocol == Https {
			manager := &autocert.Manager{
				Prompt:      autocert.AcceptTOS,
				Cache:       autocert.DirCache("acme"),
				HostPolicy:  autocert.HostWhitelist("api.renhj.cc", "www.renhj.cc", "renhj.cc"),
				RenewBefore: 10,
			}
			global.Log.Infof("Application %s:%s run at: ", global.App.AppName, global.App.AppVersion)
			global.Log.Infof("Network: https://%s:%d/", getLocalIPv4(), 443)
			if err := autotls.RunWithManager(global.Routers, manager); err != nil {
				global.Log.Fatal(err)
			}
			manager.Listener()
		}

		global.Log.Infof("Application %s: %s run at: ", global.App.AppName, global.App.AppVersion)
		global.Log.Infof("Local:   http://%s:%d/", getLoopback(), global.App.Server.Port)
		global.Log.Infof("Network: http://%s:%d/", getLocalIPv4(), global.App.Server.Port)
		if err := global.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Log.Fatal(err)
		}
	}()
}


func getLocalIPv4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getLoopback() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
