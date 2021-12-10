package viper

import (
	"blog/core/storage"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func getDecodeHookConfig() mapstructure.DecodeHookFunc {
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
		if t != reflect.TypeOf(storage.Size(1)) {
			return data, nil
		}

		return storage.ParseSize(data), nil
	}
}
