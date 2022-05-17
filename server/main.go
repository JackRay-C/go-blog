package main

import (
	"blog/internal/server"
	"blog/pkg/global"
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// usage:
//	Options:
//		-c --config  	config.yaml file path


func main() {
	var cfg string
	pflag.StringVarP(&cfg, "config", "c", "conf/default.yaml", "config file")
	pflag.Parse()

	_, err := os.Stat(cfg)
	if os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}

	server.New(cfg)
	server.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)

	for true {
		sig := <-ch
		global.Log.Infof("os signal: %v", sig)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := global.Server.Shutdown(ctx); err != nil {
				global.Log.Fatalf("server shutdown: %s", err)
			}
			global.Log.Infof("shutdown server ...")
			signal.Stop(ch)
			os.Exit(0)
			return
		case syscall.SIGUSR2:
			global.Log.Infof("reload")
			server.New(cfg)
			server.Start()
		}
	}
}


