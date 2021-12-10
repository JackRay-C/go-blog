package setting

import "time"

type Server struct {
	Port         int           `mapstructure:"port"`
	EnableTls    bool          `mapstructure:"enableTLS"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}
