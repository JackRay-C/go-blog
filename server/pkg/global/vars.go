package global

import (
	"blog/internal/config"
	"blog/internal/logger"
	"blog/internal/mail"
	"blog/internal/snowflake"
	"blog/internal/storage"
	"github.com/JackRay-C/go-mapcache"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
)

var (
	Viper     *viper.Viper
	App       *config.App
	Log       logger.Logger
	Cache     mapcache.Cache
	DB        *gorm.DB
	Snowflake *snowflake.Snowflake
	Storage   storage.Storage
	Mail      *mail.Email
	Server    *http.Server
	Routers   *gin.Engine
)


