package setting

import (
	"blog/core/storage"
	"time"
)

type App struct {
	Name               string        `mapstructure:"name"`
	Version            string        `mapstructure:"version"`
	RunMode            string        `mapstructure:"run-mode"`
	StaticPath         string        `mapstructure:"static-path"`
	StorageType        string        `mapstructure:"storage-type"`
	UploadMaxSize      storage.Size  `mapstructure:"upload-max-size"`
	UploadAllowExts    []string      `mapstructure:"upload-allow-exts"`
	LogColorConsole    bool          `mapstructure:"log-color-console"`
	DBType             string        `mapstructure:"db-type"`
	AccessTokenExpire  time.Duration `mapstructure:"access-token-expire"`
	RefreshTokenExpire time.Duration `mapstructure:"refresh-token-expire"`
}
