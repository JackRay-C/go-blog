package viper

import (
	"blog/core/global"
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)



func Viper(path string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType(ViperDefaultConfigType)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed: ", e.Name)
		if err := v.Unmarshal(&global.Setting, viper.DecodeHook(getDecodeHookConfig())); err != nil {
			panic(fmt.Errorf("Fatal error unmarshal config file: %s \n", err))
		}
	})
	if err := v.Unmarshal(&global.Setting, viper.DecodeHook(getDecodeHookConfig())); err != nil {
		panic(fmt.Errorf("Fatal error unmarshal config file: %s \n", err))
	}
	return v
}

func ViperEmbed(b []byte) *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	err := v.ReadConfig(bytes.NewBuffer(b))
	if err != nil {
		panic(fmt.Errorf("Fatal error read ember config file: %s \n", err))
	}
	if err := v.Unmarshal(&global.Setting, viper.DecodeHook(getDecodeHookConfig())); err != nil {
		panic(fmt.Errorf("Fatal error unmarshal embed config file: %s \n", err))
	}
	return v
}
