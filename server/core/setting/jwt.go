package setting

import (
	"time"
)

type Jwt struct {
	Secret   string        `mapstructure:"secret"`
	Expire   time.Duration `mapstructure:"expire"`
	Issuer   string        `mapstructure:"issuer"`
}
