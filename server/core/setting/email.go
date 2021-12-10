package setting

type Smtp struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	IsSSL    bool   `mapstructure:"isSSL"`
	From     string `mapstructure:"from"`
}
