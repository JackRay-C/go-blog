package config

import (
	"encoding/json"
	"time"
)

type AppMode int8

const (
	DevelopmentMode AppMode = iota
	ProductionMode
)

func (a AppMode) String() string {
	switch a {
	case DevelopmentMode:
		return "development"
	case ProductionMode:
		return "production"
	default:
		return "unknown"
	}
}

type App struct {
	AppMode         string    `mapstructure:"app-mode"`
	AppName         string    `mapstructure:"app-name"`
	AppVersion      string    `mapstructure:"app-version"`
	AppDatabaseType string    `mapstructure:"app-database-type"`
	AppStorageType  string    `mapstructure:"app-storage-type"`
	AppHomePath     string    `mapstructure:"app-home-path"`
	AppConfigName   string    `mapstructure:"app-config-name"`
	AppConfigType   string    `mapstructure:"app-config-type"`
	AppLogType      string    `mapstructure:"app-log-type"`
	Logs            *Logs     `mapstructure:"logs"`
	Database        *Database `mapstructure:"database"`
	Smtp            *Smtp     `mapstructure:"smtp"`
	Server          *Server   `mapstructure:"server"`
	Storage         *Storage  `mapsturcture:"storage"`
	Redis           *Redis    `mapsturcture:"redis"`
}

func (a *App) String() string {
	marshal, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type Redis struct {
	Addr     string `mapsturcture:"addr"`
	Password string `mapsturcture:"password"`
	Db       int    `mapsturcture:"db"`
	MaxRetry int    `mapsturcture:"max-retry"`
	PoolSize int    `mapsturcture:"pool-size"`
}

type Logs struct {
	Simple *Simple `mapsturcture:"simple"`
	Zap    *Zap    `mapstructure:"zap"`
}

type Simple struct {
	Level        string `mapstructure:"level"`
	LogInConsole bool   `mapstructure:"log-in-console"`
	Directory    string `mapstructure:"directory"`
	FileName     string `mapstructure:"file-name"`
	LogMaxSize   int    `mapstructure:"log-max-size"`
	LogMaxAge    int    `mapstructure:"log-max-age"`
	Format       string `mapstructure:"format"` // text / json
}

type Zap struct {
	Level         string `mapstructure:"level"`           // 日志级别
	Format        string `mapstructure:"format"`          // 格式化
	Prefix        string `mapstructure:"prefix"`          // 前缀
	Directory     string `mapstructure:"directory"`       // 日志文件夹
	LinkName      string `mapstructure:"link-name"`       // 软连接名称
	ShowLine      bool   `mapstructure:"show-line"`       // 是否显示行
	EncodeLevel   string `mapstructure:"encode-level"`    // 格式化
	StacktraceKey string `mapstructure:"stack-trace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console"`  // 输出到控制台
	LogMaxSize    int    `mapstructure:"log-max-size"`    // 日志滚动大小
	LogMaxAge     int    `mapstructure:"log-max-age"`     // 日志存储时间
	LogMaxBackups int    `mapstructure:"log-max-backups"` // 日志备份时间
}

type Database struct {
	*Mysql
	*MongoDB
	*Sqlite3
}

type Snowflake struct {
	WorkID       int64 `mapstructure:"work-id"`
	DataCenterID int64 `mapstructure:"data-center-id"`
}

type Sqlite3 struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}

type Mysql struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DbName       string `mapstructure:"db-name"`
	ParseTime    bool   `mapstructure:"parse-time"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode"`
	Charset      string `mapstructure:"charset"`
}

type MongoDB struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"db-name"`
	LogMode  string `mapstructure:"log-mode"`
}

type Smtp struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	IsSSL       bool   `mapstructure:"isSSL"`
	FromName    string `mapstructure:"from-name"`
	FromAddress string `mapstructure:"from-address"`
}

type Server struct {
	Protocol              string        `mapstructure:"protocol"`
	Addr                  string        `mapstructure:"addr"`
	Domain                string        `mapstructure:"domain"`
	Port                  int           `mapstructure:"port"`
	ReadTimeout           time.Duration `mapstructure:"readTimeout"`
	WriteTimeout          time.Duration `mapstructure:"writeTimeout"`
	StaticRootPath        string        `mapstructure:"static-root-path"`
	CertFile              string        `mapstructure:"cert-file"`
	CertKey               string        `mapstructure:"cert-key"`
	AccessTokenExpire     time.Duration `mapstructure:"access-token-expire"`
	RefreshTokenExpire    time.Duration `mapstructure:"refresh-token-expire"`
	SnowflakeWorkId       int64         `mapstructure:"snow-flake-work-id"`
	SnowflakeDataCenterId int64         `mapstructure:"snow-flake-data-center-id"`
}

type Storage struct {
	AllowUploadMaxSize  Size        `mapstructure:"allow-upload-max-size"`
	AllowUploadFileExts []string    `mapstructure:"allow-upload-file-exts"`
	Local               *Local      `mapstructure:"local"`
	Qiniu               *Qiniu      `mapstructure:"qiniu"`
	TencentOSS          *TencentOSS `mapstructure:"tencent-oss"`
	AliyunOSS           *AliyunOSS  `mapstructure:"aliyun-oss"`
}

type Local struct {
	Path string `mapstructure:"path"`
	Host string `mapstructure:"host"`
}

type Qiniu struct {
	Zone      string `mapstructure:"zone"`       // 存储区域
	Bucket    string `mapstructure:"bucket"`     // 空间名称
	ImgPath   string `mapstructure:"img-path"`   // cdm加速域名
	UseHTTPS  bool   `mapstructure:"use-https"`  // 是否使用https
	AccessKey string `mapstructure:"access-key"` // 七牛密钥
	SecretKey string `mapstructure:"secret-key"` // 七牛密钥
	Domain    string `mapstructure:"domain"`     // 域名
}

type AliyunOSS struct {
	AccessKeyId      string `mapstructure:"access-key-id"`
	AccessKeySecret  string `mapstructure:"access-key-secret"`
	BucketName       string `mapstructure:"bucket-name"`
	BucketUrl        string `mapstructure:"bucket-url"`
	Endpoint         string `mapstructure:"endpoint"`
	BasePath         string `mapstructure:"test-path"`
	HTTPTimeout      int64  `mapstructure:"http-timeout"`
	ReadWriteTimeout int64  `mapstructure:"read-write-timeout"`
	EnableCRC        bool   `mapstructure:"enable-crc"`
	ForbidOverWrite  bool   `mapstructure:"forbid-over-write"`
}

type TencentOSS struct {
	AccessKeyId     string `mapstructure:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret"`
	BucketUrl       string `mapstructure:"bucket-url"`
}
