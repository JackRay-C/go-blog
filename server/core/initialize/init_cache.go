package initialize

import (
	"blog/core/global"
	"github.com/JackRay-C/go-mapcache"
)

func SetupCache()  {
	global.Cache = mapcache.NewCache()
}
