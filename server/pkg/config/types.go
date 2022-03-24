package config

import (
	"blog/core/storage"
	"time"
)

type Snowflake struct {
	WorkID       int64 `mapstructure:"work-id"`
	DataCenterID int64 `mapstructure:"data-center-id"`
}

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

type Smtp struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	IsSSL    bool   `mapstructure:"isSSL"`
	From     string `mapstructure:"from"`
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

type Local struct {
	Path string `mapstructure:"path"`
}

type Qiniu struct {
	Zone      string `mapstructure:"zone"`       // 存储区域
	Bucket    string `mapstructure:"bucket"`     // 空间名称
	ImgPath   string `mapstructure:"img-path"`   // cdm加速域名
	UseHTTPS  bool   `mapstructure:"use-https"`  // 是否使用https
	AccessKey string `mapstructure:"access-key"` // 七牛密钥
	SecretKey string `mapstructure:"secret-key"` // 七牛密钥
}

type AliyunOSS struct {
	AccessKeyId      string `mapstructure:"access-key-id"`
	AccessKeySecret  string `mapstructure:"access-key-secret"`
	BucketName       string `mapstructure:"bucket-name"`
	BucketUrl        string `mapstructure:"bucket-url"`
	Endpoint         string `mapstructure:"endpoint"`
	BasePath         string `mapstructure:"base-path"`
	HTTPTimeout      int64  `mapstructure:"http-timeout"`
	ReadWriteTimeout int64  `mapstructure:"read-write-timeout"`
	EnableCRC        bool   `mapstructure:"enable-crc"`
	ForbidOverWrite  bool   `mapstructure:"forbid-over-write"`
}

type TencentOSS struct {
	AccessKeyId string `mapstructure:"access-key-id"`
	BucketUrl   string `mapstructure:"bucket-url"`
}

type Server struct {
	Port         int           `mapstructure:"port"`
	EnableTls    bool          `mapstructure:"enableTLS"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}
