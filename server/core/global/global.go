package global

import (
	"blog/app/mail"
	"blog/core/logger"
	"blog/core/setting"
	"blog/core/snowflake"
	"blog/core/storage"
	"github.com/JackRay-C/go-mapcache"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Viper     *viper.Viper
	Logger    logger.Logger
	DB        *gorm.DB
	Setting   *setting.Setting
	Snowflake *snowflake.Snowflake
	Storage   storage.Storage
	Email     *mail.Email
	Cache     mapcache.Cache
)
