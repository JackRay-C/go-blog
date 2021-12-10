package setting

type Setting struct {
	Server     *Server     `mapstructure:"server"`      // 服务器配置
	Mysql      *Mysql      `mapstructure:"mysql"`    // mysql数据库配置
	App        *App        `mapstructure:"app"`         //  应用配置
	Smtp       *Smtp       `mapstructure:"smtp"`        // 邮件配置
	Jwt        *Jwt        `mapstructure:"jwt"`         // jwt配置
	Zap        *Zap        `mapstructure:"zap"`         // zap日志配置
	Local      *Local      `mapstructure:"local"`       // 本地文件存储
	Qiniu      *Qiniu      `mapstructure:"qiniu"`       // 七牛云文件存储
	AliyunOSS  *AliyunOSS  `mapstructure:"aliyun-oss"` // 阿里云oss存储
	TencentOSS *TencentOSS `mapstructure:"tencent-oss"` // 腾讯oss存储
	Snowflake  *Snowflake  `mapstructure:"snowflake"`   // 雪花ID生成器
}
