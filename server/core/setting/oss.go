package setting

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
