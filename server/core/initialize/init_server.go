package initialize

import (
	"blog/app/routes"
	"blog/core/global"
	"blog/core/server"
	"fmt"
	"github.com/gin-gonic/autotls"
	"golang.org/x/crypto/acme/autocert"
	"net"
)

func StartServer() {

	routers := routes.NewRouters(global.Setting.App)

	if global.Setting.Server.EnableTls {

		manager := &autocert.Manager{
			Prompt:      autocert.AcceptTOS,
			Cache:       autocert.DirCache("acme"),
			HostPolicy:  autocert.HostWhitelist("api.renhj.cc", "www.renhj.cc", "renhj.cc"),
			RenewBefore: 10,
		}

		if err := autotls.RunWithManager(routers, manager); err != nil {
			panic(err)
		}

	} else {

		s := server.NewServer(global.Setting.Server, routers)

		fmt.Println()
		fmt.Printf("App %s:%s run at: \n", global.Setting.App.Name, global.Setting.App.Version)
		fmt.Printf("- Local:   http://%s:%d/\n", getLoopback(), global.Setting.Server.Port)
		fmt.Printf("- Network: http://%s:%d/\n", getLocalIPv4(), global.Setting.Server.Port)
		fmt.Println()
		err := s.ListenAndServe()
		if err != nil {
			panic(fmt.Errorf("启动服务器失败: %s\n", err))
		}

	}

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
