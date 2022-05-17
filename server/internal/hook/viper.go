package hook

import (
	"blog/internal/config"
	"blog/internal/storage"
	"blog/pkg/global"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"path"
	"reflect"
)


func GetDecodeHookConfig() mapstructure.DecodeHookFunc {
	return mapstructure.ComposeDecodeHookFunc(mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		mapstructure.StringToIPHookFunc(),
		mapstructure.StringToIPNetHookFunc(),
		StorageSizeHookFunc())
}

func StorageSizeHookFunc() mapstructure.DecodeHookFuncType {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {

		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(config.Size(1)) {
			return data, nil
		}

		return storage.ParseSize(data), nil
	}
}

func New(app *config.App) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath(path.Join(app.AppHomePath, "conf"))

	v.SetConfigName(app.AppConfigName)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Fatal error read config file: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		global.Log.Infof("hook config file \"%s\" changed, reload. ", e.Name)
		if err := v.Unmarshal(&global.App, viper.DecodeHook(GetDecodeHookConfig())); err != nil {
			global.Log.Fatalf("config changed, Fatal error unmarshal config file: \"%s\"", err)
		}
	})
	if err := v.Unmarshal(&global.App, viper.DecodeHook(GetDecodeHookConfig())); err != nil {
		return nil, errors.New(fmt.Sprintf("Fatal error unmarshal config file: \"%s\"", err))
	}
	return v, nil
}
