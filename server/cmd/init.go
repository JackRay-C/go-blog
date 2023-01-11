package cmd

import (
	"blog/internal/config"
	"blog/internal/storage"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const (
	// DefaultConfigName Default config
	DefaultConfigName                = "default"
	DefaultConfigType                = "yaml"
	DefaultConfigPath                = "conf"
	DefaultStaticPath                = "static"
	DefaultBinaryPath                = "bin"
	DefaultAppMode                   = "production"
	DefaultAppName                   = "go-blog"
	DefaultHomePath                  = "."
	DefaultLogType                   = "simple"
	DefaultLogFile                   = "debug.log"
	DefaultLogPath                   = "logs"
	DefaultLogLevel                  = "info"
	DefaultLogFormat                 = "console"
	DefaultLogInConsole              = true
	DefaultLogPrefix                 = ""
	DefaultLogEncodeLevel            = "LowercaseColorLevelEncoder"
	DefaultLogStackTraceKey          = ""
	DefaultLogMaxSize                = 1000
	DefaultLogMaxAge                 = 30
	DefaultLogMaxBackups             = 5
	DefaultStorageType               = "local"
	DefaultStorageUploadMaxSize      = 2 * storage.GB
	DefaultLocalPath                 = "uploads"
	DefaultLocalHost                 = "http://localhost:8000"
	DefaultAliyunOSSEndpoint         = ""
	DefaultAliyunOSSAccessKeyId      = ""
	DefaultAliyunOSSAccessKeySecret  = ""
	DefaultAliyunOSSBucketName       = ""
	DefaultAliyunOSSBucketUrl        = ""
	DefaultAliyunOSSHttpTimeout      = 10
	DefaultAliyunOSSReadWriteTimeout = 200
	DefaultAliyunOSSEnableCrc        = false
	DefaultAliyunOSSForbidOverWrite  = true
	DefaultQiniuZone                 = ""
	DefaultQiniuBucket               = ""
	DefaultQiniuImgPath              = ""
	DefaultQiniuUseHttps             = true
	DefaultQiniuAccessKey            = ""
	DefaultQiniuSecretKey            = ""
	DefaultQiniuDomain               = ""
	DefaultTencentOSSBucket          = ""
	DefaultTencentOSSRegion          = ""
	DefaultTencentOSSSecretID        = ""
	DefaultTencentOSSSecretKey       = ""
	DefaultTencentOSSBaseUrl         = ""
	DefaultTencentOSSBasePrefix      = ""
	DefaultDatabaseType              = "mysql"
	DefaultMysqlUsername             = "root"
	DefaultMysqlPassword             = ""
	DefaultMysqlHost                 = "localhost"
	DefaultMysqlPort                 = 3306
	DefaultMysqlDBName               = "go_blog"
	DefaultMysqlParseTime            = true
	DefaultMysqlCharset              = "utf8"
	DefaultMysqlLogMode              = "info"
	DefaultMysqlMaxIdleConns         = 10
	DefaultMysqlMaxOpenConns         = 10
	DefaultMongodbHost               = "localhost"
	DefaultMongodbUsername           = ""
	DefaultMongodbPassword           = ""
	DefaultSqlite3Username           = ""
	DefaultSqlite3Password           = ""
	DefaultSqlite3Host               = "localhost"
	DefaultSmtpHost                  = ""
	DefaultSmtpPort                  = 465
	DefaultSmtpUsername              = ""
	DefaultSmtpPassword              = ""
	DefaultSmtpIsSSL                 = true
	DefaultSmtpFromName              = "go-blog"
	DefaultSmtpFromAddress           = "admin@go-blog.localhost"
	DefaultServerProtocol            = "http"
	DefaultServerPort                = 8000
	DefaultServerAddr                = "localhost"
	DefaultServerDomain              = "www.renhj.cc"
	DefaultServerReadTimeout         = 60 * time.Second
	DefaultServerWriteTimeout        = 60 * time.Second
	DefaultServerStaticRootPath      = "public"
	DefaultServerCertFile            = "cert.crt"
	DefaultServerCertKey             = "cert.key"
	DefaultServerAccessTokenExpire   = 10 * 60 * time.Second
	DefaultServerRefreshTokenExpire  = 30 * 24 * time.Hour
	DefaultServerWorkId              = 1
	DefaultServerDataCenterId        = 1
	DefaultRedisAddr                 = "localhost:6379"
	DefaultRedisPassword             = ""
	DefaultRedisDb                   = 0
	DefaultRedisMaxRetry             = 3
	DefaultRedisPoolSize             = 10
)

var (
	v                            = viper.New()
	appConfig                    = &config.App{AppName: DefaultAppName, AppVersion: BuildVersion}
	serverConfig                 = &config.Server{}
	mysqlConfig                  = &config.Mysql{}
	mongodbConfig                = &config.MongoDB{}
	sqlite3Config                = &config.Sqlite3{}
	simpleLoggerConfig           = &config.Simple{}
	zapLoggerConfig              = &config.Zap{}
	smtpConfig                   = &config.Smtp{}
	redisConfig                  = &config.Redis{}
	supportLogTpe                = stringArray{"simple", "zap"}
	supportLogFormat             = stringArray{"console", "json"}
	supportAppMode               = stringArray{"production", "development"}
	supportDBType                = stringArray{"mysql", "sqlite3", "mongodb"}
	supportServerProtocol        = stringArray{"http", "https", "http2", "socket"}
	supportLogLevel              = stringArray{"info", "debug", "fatal", "warn", "error"}
	supportStorageType           = stringArray{"local", "aliyun-oss", "tencent-oss", "qiniu"}
	supportConfigType            = stringArray{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "dotenv", "env", "ini"}
	DefaultStorageUploadFileExts = []string{".jpg", "jpeg", ".png", ".bmp", ".psd", ".icon", ".gif", ".mp3", ".wav", ".mid", ".flac','.mp4','.mov", ".avi", ".flv", ".m4v", ".rmvb", ".zip", ".rar", ".tar", ".gz", ".ba2", ".doc", ".xlsx", ".pdf", ".ppt", ".pptx", ".xls", ".docx", ".md", ".html", ".js", ".css", ".java", ".class", ".py", ".go", ".sh", ".log", ".yaml", ".yml','.ini','.json", ".txt", ".xmind"}
	configStorage                = &config.Storage{AllowUploadMaxSize: DefaultStorageUploadMaxSize, AllowUploadFileExts: DefaultStorageUploadFileExts, Local: &config.Local{}, TencentOSS: &config.TencentOSS{}, AliyunOSS: &config.AliyunOSS{}, Qiniu: &config.Qiniu{}}
	initCommand                  = &cobra.Command{
		Use:   "init [flags]",
		Short: "Init go blog directory",
		Run: func(cmd *cobra.Command, args []string) {
			checkSupport()
			writeConfigs()
		},
	}
)

const (
	// Storage type
	Local      = "local"
	AliyunOSS  = "aliyun-oss"
	TencentOSS = "tencent-oss"
	Qiniu      = "qiniu"
	// Mysql Database Type
	MysqlDBType   = "mysql"
	MongoDBType   = "mongodb"
	Sqlite3DBType = "sqlite3"
	// Logger type
	SimpleLoggerType = "simple"
	ZapLoggerType    = "zap"
)

func init() {
	initFlags()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//cobra.OnInitialize(checkSupport, writeConfigs, initConfig)
}

func initFlags() {
	initCommand.Flags().StringVar(&appConfig.AppMode, "app-mode", DefaultAppMode, "server start mode")
	initCommand.Flags().StringVarP(&appConfig.AppHomePath, "home", "d", DefaultHomePath, "init directory")
	initCommand.Flags().StringVar(&appConfig.AppDatabaseType, "database", DefaultDatabaseType, "database type")
	initCommand.Flags().StringVar(&appConfig.AppStorageType, "storage", DefaultStorageType, "storage type")
	//initCommand.Flags().StringVar(&configFileName, "config-file-name", DefaultConfigName, "If noe-empty, config file name is this")
	//initCommand.Flags().StringVar(&configFileType, "config-file-type", DefaultConfigType, "hook config file type")
	initCommand.Flags().StringVar(&appConfig.AppLogType, "logs", DefaultLogType, "Logger type, possible: simple, zap")

	// ###################################### logger bind flags ##################################
	initCommand.Flags().StringVar(&simpleLoggerConfig.Directory, "logs.simple.directory", DefaultLogPath, "Simple logger directory")
	initCommand.Flags().BoolVar(&simpleLoggerConfig.LogInConsole, "logs.simple.log-in-console", DefaultLogInConsole, "If true, simple logger in file also console")
	initCommand.Flags().StringVar(&simpleLoggerConfig.Level, "logs.simple.level", DefaultLogLevel, "Log level")
	initCommand.Flags().StringVar(&simpleLoggerConfig.FileName, "logs.simple.log-file", DefaultLogFile, "Log file name")
	initCommand.Flags().IntVar(&simpleLoggerConfig.LogMaxSize, "logs.simple.log-max-size", DefaultLogMaxSize, "Log file max size")
	initCommand.Flags().IntVar(&simpleLoggerConfig.LogMaxAge, "logs.simple.log-max-age", DefaultLogMaxAge, "Log file max age")
	initCommand.Flags().StringVar(&simpleLoggerConfig.Format, "logs.simple.format", DefaultLogFormat, "Log file format")
	initCommand.Flags().StringVar(&zapLoggerConfig.Directory, "logs.zap.directory", DefaultLogPath, "Zap logger directory")
	initCommand.Flags().StringVar(&zapLoggerConfig.Level, "logs.zap.level", DefaultLogLevel, "Zap logger level")
	initCommand.Flags().StringVar(&zapLoggerConfig.Format, "logs.zap.format", DefaultLogFormat, "Zap logger format, possible: console, json")
	initCommand.Flags().StringVar(&zapLoggerConfig.Prefix, "logs.zap.prefix", DefaultLogPrefix, "Zap logger prefix")
	initCommand.Flags().StringVar(&zapLoggerConfig.LinkName, "logs.zap.filename", DefaultLogFile, "Zap logger filename")
	initCommand.Flags().StringVar(&zapLoggerConfig.EncodeLevel, "logs.zap.encode-level", DefaultLogEncodeLevel, "Zap logger encode level")
	initCommand.Flags().StringVar(&zapLoggerConfig.StacktraceKey, "logs.zap.stack-trace-key", DefaultLogStackTraceKey, "Zap logger stack trace key")
	initCommand.Flags().BoolVar(&zapLoggerConfig.LogInConsole, "logs.zap.log-in-console", DefaultLogInConsole, "If true, zap logger in file also console")
	initCommand.Flags().IntVar(&zapLoggerConfig.LogMaxSize, "logs.zap.log-max-size", DefaultLogMaxSize, "Zap log rotate when log file size")
	initCommand.Flags().IntVar(&zapLoggerConfig.LogMaxAge, "logs.zap.log-max-age", DefaultLogMaxAge, "Zap log rotate  age")
	initCommand.Flags().IntVar(&zapLoggerConfig.LogMaxBackups, "logs.zap.log-max-backups", DefaultLogMaxBackups, "Zap log rotate backups day")
	// ###################################### logger bind flags ##################################

	// ###################################### storage bind flags #################################
	initCommand.Flags().Var(&configStorage.AllowUploadMaxSize, "storage.allow-upload-max-size", "storage allow upload max size")
	initCommand.Flags().StringVar(&configStorage.Local.Path, "storage.local.path", DefaultLocalPath, "local upload path")
	initCommand.Flags().StringVar(&configStorage.Local.Host, "storage.local.host", DefaultLocalHost, "local upload host")
	initCommand.Flags().StringVar(&configStorage.Qiniu.Zone, "storage.qiniu.zone", DefaultQiniuZone, "qiniu oss zone")
	initCommand.Flags().StringVar(&configStorage.Qiniu.Bucket, "storage.qiniu.bucket", DefaultQiniuBucket, "qiniu oss bucket")
	initCommand.Flags().StringVar(&configStorage.Qiniu.ImgPath, "storage.qiniu.img-path", DefaultQiniuImgPath, "qiniu oss img path")
	initCommand.Flags().BoolVar(&configStorage.Qiniu.UseHTTPS, "storage.qiniu.use-https", DefaultQiniuUseHttps, "qiniu oss use https")
	initCommand.Flags().StringVar(&configStorage.Qiniu.AccessKey, "storage.qiniu.access-key", DefaultQiniuAccessKey, "qiniu oss access key")
	initCommand.Flags().StringVar(&configStorage.Qiniu.SecretKey, "storage.qiniu.secret-key", DefaultQiniuSecretKey, "qiniu oss secret key")
	initCommand.Flags().StringVar(&configStorage.Qiniu.Domain, "storage.qiniu.domain", DefaultQiniuDomain, "qiniu oss domain")
	initCommand.Flags().StringVar(&configStorage.AliyunOSS.Endpoint, "storage.aliyun-oss.endpoint", DefaultAliyunOSSEndpoint, "aliyun oss endpoint")
	initCommand.Flags().StringVar(&configStorage.AliyunOSS.AccessKeyId, "storage.aliyun-oss.access-key-id", DefaultAliyunOSSAccessKeyId, "aliyun oss access key id")
	initCommand.Flags().StringVar(&configStorage.AliyunOSS.AccessKeySecret, "storage.aliyun-oss.access-key-secret", DefaultAliyunOSSAccessKeySecret, "aliyun oss access key secret")
	initCommand.Flags().StringVar(&configStorage.AliyunOSS.BucketName, "storage.aliyun-oss.bucket-name", DefaultAliyunOSSBucketName, "aliyun oss bucket name")
	initCommand.Flags().StringVar(&configStorage.AliyunOSS.BucketUrl, "storage.aliyun-oss.bucket-url", DefaultAliyunOSSBucketUrl, "aliyun oss bucket url")
	initCommand.Flags().Int64Var(&configStorage.AliyunOSS.HTTPTimeout, "storage.aliyun-oss.http-timeout", DefaultAliyunOSSHttpTimeout, "aliyun oss http timeout")
	initCommand.Flags().Int64Var(&configStorage.AliyunOSS.ReadWriteTimeout, "storage.aliyun-oss.read-write-timeout", DefaultAliyunOSSReadWriteTimeout, "aliyun oss read write timeout")
	initCommand.Flags().BoolVar(&configStorage.AliyunOSS.EnableCRC, "storage.aliyun-oss.enable-crc", DefaultAliyunOSSEnableCrc, "aliyun oss enable crc")
	initCommand.Flags().BoolVar(&configStorage.AliyunOSS.ForbidOverWrite, "storage.aliyun-oss.forbid-over-write", DefaultAliyunOSSForbidOverWrite, "aliyun oss forbid over write")
	initCommand.Flags().StringVar(&configStorage.TencentOSS.AccessKeyId, "storage.tencent-oss.access-key-id", DefaultTencentOSSSecretID, "tencent oss access key id")
	initCommand.Flags().StringVar(&configStorage.TencentOSS.AccessKeySecret, "storage.tencent-oss.access-key-secret", DefaultTencentOSSSecretKey, "tencent oss access key secret")
	// ###################################### storage bind flags #################################

	// ###################################### server  bind flags #################################
	initCommand.Flags().StringVar(&serverConfig.Protocol, "server.protocol", DefaultServerProtocol, "server protocol")
	initCommand.Flags().StringVar(&serverConfig.Addr, "server.addr", DefaultServerAddr, "server start host")
	initCommand.Flags().IntVar(&serverConfig.Port, "server.port", DefaultServerPort, "server start port")
	initCommand.Flags().StringVar(&serverConfig.Domain, "server.domain", DefaultServerDomain, "server domain")
	initCommand.Flags().DurationVar(&serverConfig.ReadTimeout, "server.read-timeout", DefaultServerReadTimeout, "server read timeout")
	initCommand.Flags().DurationVar(&serverConfig.WriteTimeout, "server.write-timeout", DefaultServerWriteTimeout, "server write timeout")
	initCommand.Flags().StringVar(&serverConfig.StaticRootPath, "server.static-root-path", DefaultServerStaticRootPath, "server static root path")
	initCommand.Flags().StringVar(&serverConfig.CertFile, "server.cert-file", DefaultServerCertFile, "server https cert file")
	initCommand.Flags().StringVar(&serverConfig.CertKey, "server.cert-key", DefaultServerCertKey, "server https cert key")
	initCommand.Flags().DurationVar(&serverConfig.AccessTokenExpire, "server.access-token-expire", DefaultServerAccessTokenExpire, "access token expire duration")
	initCommand.Flags().DurationVar(&serverConfig.RefreshTokenExpire, "server.refresh-token-expire", DefaultServerRefreshTokenExpire, "refresh token expire duration")
	initCommand.Flags().Int64Var(&serverConfig.SnowflakeWorkId, "server.snowflake-work-id", DefaultServerWorkId, "snowflake work id")
	initCommand.Flags().Int64Var(&serverConfig.SnowflakeDataCenterId, "server.snowflake-data-center-id", DefaultServerDataCenterId, "snowflake data center id")
	// ###################################### server  bind flags ##################################

	// ###################################### database bind flags #################################
	initCommand.Flags().StringVar(&mysqlConfig.Host, "database.mysql.host", DefaultMysqlHost, "mysql server host")
	initCommand.Flags().StringVar(&mysqlConfig.Username, "database.mysql.username", DefaultMysqlUsername, "mysql server username")
	initCommand.Flags().StringVar(&mysqlConfig.Password, "database.mysql.password", DefaultMysqlPassword, "mysql server password")
	initCommand.Flags().IntVar(&mysqlConfig.Port, "database.mysql.port", DefaultMysqlPort, "mysql server port")
	initCommand.Flags().StringVar(&mysqlConfig.DbName, "database.mysql.db-name", DefaultMysqlDBName, "server database name")
	initCommand.Flags().StringVar(&mysqlConfig.Charset, "database.mysql.charset", DefaultMysqlCharset, "db connect charset")
	initCommand.Flags().StringVar(&mysqlConfig.LogMode, "database.mysql.log-mode", DefaultMysqlLogMode, "db connect pool log mode")
	initCommand.Flags().BoolVar(&mysqlConfig.ParseTime, "database.mysql.parse-time", DefaultMysqlParseTime, "mysql parse time")
	initCommand.Flags().IntVar(&mysqlConfig.MaxIdleConns, "database.mysql.max-idle-conns", DefaultMysqlMaxIdleConns, "db connect pool idle conn")
	initCommand.Flags().IntVar(&mysqlConfig.MaxOpenConns, "database.mysql.max-open-conns", DefaultMysqlMaxOpenConns, "db connect pool open conn")
	initCommand.Flags().StringVar(&mongodbConfig.Host, "database.mongodb.host", DefaultMongodbHost, "mongodb host")
	initCommand.Flags().StringVar(&mongodbConfig.Username, "database.mongodb.username", DefaultMongodbUsername, "mongodb username")
	initCommand.Flags().StringVar(&mongodbConfig.Password, "database.mongodb.password", DefaultMongodbPassword, "mongodb password")
	initCommand.Flags().StringVar(&sqlite3Config.Password, "database.sqlite3.password", DefaultSqlite3Password, "sqlite3 password")
	initCommand.Flags().StringVar(&sqlite3Config.Username, "database.sqlite3.username", DefaultSqlite3Username, "sqlite3 username")
	initCommand.Flags().StringVar(&sqlite3Config.Host, "database.sqlite3.host", DefaultSqlite3Host, "sqlite3 host")
	// ###################################### database  bind flags #################################

	// ###################################### smtp  bind flags #####################################
	initCommand.Flags().StringVar(&smtpConfig.Username, "smtp.username", DefaultSmtpUsername, "smtp username")
	initCommand.Flags().StringVar(&smtpConfig.Password, "smtp.password", DefaultSmtpPassword, "smtp username")
	initCommand.Flags().StringVar(&smtpConfig.Host, "smtp.host", DefaultSmtpHost, "smtp host")
	initCommand.Flags().BoolVar(&smtpConfig.IsSSL, "smtp.isSSL", DefaultSmtpIsSSL, "smtp is SSL")
	initCommand.Flags().StringVar(&smtpConfig.FromName, "smtp.from-name", DefaultSmtpFromName, "smtp from name")
	initCommand.Flags().StringVar(&smtpConfig.FromAddress, "smtp.from-address", DefaultSmtpFromAddress, "smtp from address")
	// ###################################### smtp  bind flags ###################################

	initCommand.Flags().StringVar(&redisConfig.Addr, "redis.addr", DefaultRedisAddr, "redis address")
	initCommand.Flags().StringVar(&redisConfig.Password, "redis.password", DefaultRedisPassword, "redis password")
	initCommand.Flags().IntVar(&redisConfig.Db, "redis.db", DefaultRedisDb, "redis database")
	initCommand.Flags().IntVar(&redisConfig.MaxRetry, "redis.max-retry", DefaultRedisMaxRetry, "redis connect max retry number")
	initCommand.Flags().IntVar(&redisConfig.PoolSize, "redis.pool-size", DefaultRedisPoolSize, "redis conn pool size")
}

func checkSupport() {
	if !supportLogTpe.Contain(appConfig.AppLogType) {
		log.Fatalf("un support log type: %s, possible: %s", appConfig.AppLogType, supportLogTpe)
	}

	if !supportLogLevel.Contain(simpleLoggerConfig.Level) {
		log.Fatalf("un support log level: %s, possible: %s", simpleLoggerConfig.Level, supportLogLevel)
	}

	if !supportLogLevel.Contain(zapLoggerConfig.Level) {
		log.Fatalf("un support log level: %s, possible: %s", zapLoggerConfig.Level, supportLogLevel)
	}

	if !supportLogFormat.Contain(zapLoggerConfig.Format) {
		log.Fatalf("un support log format: %s, possible: %s", zapLoggerConfig.Format, supportLogFormat)
	}

	if !supportDBType.Contain(appConfig.AppDatabaseType) {
		log.Fatalf("un support database type: %s, possible: %s", appConfig.AppDatabaseType, supportDBType)
	}

	if !supportServerProtocol.Contain(serverConfig.Protocol) {
		log.Fatalf("un support protocol: %s, possible: %s", serverConfig.Protocol, supportServerProtocol)
	}

	if !supportStorageType.Contain(appConfig.AppStorageType) {
		log.Fatalf("un support storage type: %s, possible: %s", appConfig.AppStorageType, supportStorageType)
	}

	if !supportAppMode.Contain(appConfig.AppMode) {
		log.Fatalf("un support app mode: %s, possible: %s", appConfig.AppMode, supportAppMode)
	}

}

func writeConfigs() {
	// 1、create home path
	createDir(appConfig.AppHomePath)

	v.Set("app-home-path", appConfig.AppHomePath)
	v.Set("app-mode", appConfig.AppMode)
	v.Set("app-name", appConfig.AppName)
	v.Set("app-version", appConfig.AppVersion)
	v.Set("app-storage-type", appConfig.AppStorageType)
	v.Set("app-log-type", appConfig.AppLogType)

	v.Set("smtp.username", smtpConfig.Username)
	v.Set("smtp.password", smtpConfig.Password)
	v.Set("smtp.host", smtpConfig.Host)
	v.Set("smtp.from-address", smtpConfig.FromAddress)
	v.Set("smtp.from-name", smtpConfig.FromName)
	v.Set("smtp.isSSL", smtpConfig.IsSSL)

	if appConfig.AppLogType == SimpleLoggerType {
		v.Set("logs.simple.directory", path.Join(appConfig.AppHomePath, simpleLoggerConfig.Directory))
		v.Set("logs.simple.log-in-console", simpleLoggerConfig.LogInConsole)
		v.Set("logs.simple.level", simpleLoggerConfig.Level)
		v.Set("logs.simple.file-name", simpleLoggerConfig.FileName)
		v.Set("logs.simple.log-max-size", simpleLoggerConfig.LogMaxSize)
		v.Set("logs.simple.log-max-age", simpleLoggerConfig.LogMaxAge)
		v.Set("logs.simple.format", simpleLoggerConfig.Format)
	}
	if appConfig.AppLogType == ZapLoggerType {
		v.Set("logs.zap.directory", path.Join(appConfig.AppHomePath, zapLoggerConfig.Directory))
		v.Set("logs.zap.log-in-console", zapLoggerConfig.LogInConsole)
		v.Set("logs.zap.level", zapLoggerConfig.Level)
		v.Set("logs.zap.format", zapLoggerConfig.Format)
		v.Set("logs.zap.prefix", zapLoggerConfig.Prefix)
		v.Set("logs.zap.encode-level", zapLoggerConfig.EncodeLevel)
		v.Set("logs.zap.stack-trace-key", zapLoggerConfig.StacktraceKey)
		v.Set("logs.zap.log-max-size", zapLoggerConfig.LogMaxAge)
		v.Set("logs.zap.log-max-age", zapLoggerConfig.LogMaxAge)
		v.Set("logs.zap.log-max-backups", zapLoggerConfig.LogMaxBackups)
	}

	if appConfig.AppDatabaseType == MysqlDBType {
		v.Set("database.mysql.username", mysqlConfig.Username)
		v.Set("database.mysql.password", mysqlConfig.Password)
		v.Set("database.mysql.host", mysqlConfig.Host)
		v.Set("database.mysql.port", mysqlConfig.Port)
		v.Set("database.mysql.db-name", mysqlConfig.DbName)
		v.Set("database.mysql.charset", mysqlConfig.Charset)
		v.Set("database.mysql.log-mode", mysqlConfig.LogMode)
		v.Set("database.mysql.parse-time", mysqlConfig.ParseTime)
		v.Set("database.mysql.max-idle-conns", mysqlConfig.MaxIdleConns)
		v.Set("database.mysql.max-open-conns", mysqlConfig.MaxOpenConns)
	} else if appConfig.AppDatabaseType == MongoDBType {
		v.Set("database.mongodb.username", mongodbConfig.Username)
		v.Set("database.mongodb.password", mongodbConfig.Password)
		v.Set("database.mongodb.host", mongodbConfig.Host)
	} else if appConfig.AppDatabaseType == Sqlite3DBType {
		v.Set("database.sqlite3.username", sqlite3Config.Username)
		v.Set("database.sqlite3.password", sqlite3Config.Password)
		v.Set("database.sqlite3.host", sqlite3Config.Host)
	}
	v.Set("server.protocol", serverConfig.Protocol)
	v.Set("server.addr", serverConfig.Addr)
	v.Set("server.port", serverConfig.Port)
	v.Set("server.domain", serverConfig.Domain)
	v.Set("server.readTimeout", serverConfig.ReadTimeout)
	v.Set("server.writeTimeout", serverConfig.WriteTimeout)
	v.Set("server.static-root-path", path.Join(appConfig.AppHomePath, serverConfig.StaticRootPath))
	v.Set("server.cert-file", path.Join(appConfig.AppHomePath, serverConfig.CertFile))
	v.Set("server.cert-key", path.Join(appConfig.AppHomePath, serverConfig.CertKey))
	v.Set("server.access-token-expire", serverConfig.AccessTokenExpire)
	v.Set("server.refresh-token-expire", serverConfig.RefreshTokenExpire)
	v.Set("server.snow-flake-work-id", serverConfig.SnowflakeWorkId)
	v.Set("server.snow-flake-data-center-id", serverConfig.SnowflakeDataCenterId)

	v.Set("storage.allow-upload-max-size", configStorage.AllowUploadMaxSize)
	v.Set("storage.allow-upload-file-exts", configStorage.AllowUploadFileExts)

	switch appConfig.AppStorageType {
	case Local:
		v.Set("storage.local.path", path.Join(appConfig.AppHomePath, configStorage.Local.Path))
		v.Set("storage.local.host", configStorage.Local.Host)
	case AliyunOSS:
		v.Set("storage.aliyun-oss.endpoint", configStorage.AliyunOSS.Endpoint)
		v.Set("storage.aliyun-oss.access-key-id", configStorage.AliyunOSS.AccessKeyId)
		v.Set("storage.aliyun-oss.access-key-secret", configStorage.AliyunOSS.AccessKeySecret)
		v.Set("storage.aliyun-oss.bucket-name", configStorage.AliyunOSS.BucketName)
		v.Set("storage.aliyun-oss.bucket-url", configStorage.AliyunOSS.BucketUrl)
		v.Set("storage.aliyun-oss.http-timeout", configStorage.AliyunOSS.HTTPTimeout)
		v.Set("storage.aliyun-oss.read-write-timeout", configStorage.AliyunOSS.ReadWriteTimeout)
		v.Set("storage.aliyun-oss.enable-crc", configStorage.AliyunOSS.EnableCRC)
		v.Set("storage.aliyun-oss.forbid-over-write", configStorage.AliyunOSS.ForbidOverWrite)
	case TencentOSS:
		v.Set("storage.tencent-oss.access-key-id", configStorage.TencentOSS.AccessKeyId)
		v.Set("storage.tencent-oss.access-key-secret", configStorage.TencentOSS.AccessKeySecret)
	case Qiniu:
		v.Set("storage.qiniu.zone", configStorage.Qiniu.Zone)
		v.Set("storage.qiniu.bucket", configStorage.Qiniu.Bucket)
		v.Set("storage.qiniu.img-path", configStorage.Qiniu.ImgPath)
		v.Set("storage.qiniu.use-https", configStorage.Qiniu.UseHTTPS)
		v.Set("storage.qiniu.access-key", configStorage.Qiniu.AccessKey)
		v.Set("storage.qiniu.secret-key", configStorage.Qiniu.SecretKey)
		v.Set("storage.qiniu.domain", configStorage.Qiniu.Domain)
	}

	v.Set("redis.addr", redisConfig.Addr)
	v.Set("redis.password", redisConfig.Password)
	v.Set("redis.db", redisConfig.Db)
	v.Set("redis.max-retry", redisConfig.MaxRetry)
	v.Set("redis.pool-size", redisConfig.PoolSize)

	writeViperDefaultConfigurationToInitPath()
}

func createDir(dir string) {
	if path.IsAbs(dir) {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				log.Fatalf("create directory %s failed: %s", dir, err)
			}
		}
	} else {
		// 不是绝对路径，获取绝对路径
		abs, err := filepath.Abs(dir)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := os.Stat(abs); os.IsNotExist(err) {
			if err := os.MkdirAll(abs, os.ModePerm); err != nil {
				log.Fatalf("create directory %s failed: %s", abs, err)
			}
		}
	}
}

func writeViperDefaultConfigurationToInitPath() {
	// 判断home目录及配置文件是否存在
	f := path.Join(appConfig.AppHomePath, cfgFile)

	dir := path.Dir(f)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		createDir(dir)
	}
	// 配置文件是否存在，不存在则写入，存在覆盖
	if _, err := os.Stat(f); os.IsNotExist(err) {
		if err := v.WriteConfigAs(f); err != nil {
			log.Fatal(err)
		}
		log.Printf("initialized default config to %s.", f)
	} else {
		overwrite := ""
		fmt.Printf(`config file %s has already exists. overwrite it? [Y/N] then [enter] (press 'c' to cancel): `, f)
		_, _ = fmt.Fscanln(os.Stdin, &overwrite)
		if strings.ToLower(overwrite) == "y" {
			if err := v.WriteConfigAs(f); err != nil {
				log.Fatal(err)
			}
			log.Printf("initialized default config to %s.", f)
		} else {
			os.Exit(1)
		}
	}

	logsDir := ""
	switch appConfig.AppLogType {
	case SimpleLoggerType:
		logsDir = path.Join(appConfig.AppHomePath, simpleLoggerConfig.Directory)
	case ZapLoggerType:
		logsDir = path.Join(appConfig.AppHomePath, zapLoggerConfig.Directory)
	}
	createDir(logsDir)
	log.Printf("initialized default log directory %s", logsDir)

	if appConfig.AppStorageType == Local {
		localStoragePath := path.Join(appConfig.AppHomePath, configStorage.Local.Path)
		createDir(localStoragePath)
		log.Printf("initialized default local storage directory %s success. ", localStoragePath)
	}

	log.Println("initialized success. ")
}

//func scanSetupConfiguration() {
//	db := ""
//	fmt.Print(`请选择您使用的数据库?
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//1: MySQL (Default)
//2: MongoDB
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//Select the appropriate number [1-2] then [enter] (press 'c' to cancel): `)
//	_, _ = fmt.Fscanln(os.Stdin, &db)
//
//	if len(db) == 0 || db == "" {
//		dbType = DBTypeName(1)
//	} else {
//		if len(db) == 1 {
//			a, _ := strconv.ParseInt(db, 10, 32)
//			if a == 1 {
//				dbType = DBTypeName(1)
//			} else if a == 2 {
//				dbType = DBTypeName(2)
//			} else {
//				l.Fatalf("un support database type: %s", a)
//				os.Exit(1)
//			}
//		}
//	}
//
//	for i := 0; i < 6; i++ {
//		fmt.Print("\033[1A")
//		fmt.Print("\033[2K")
//	}
//	fmt.Println("您选择的数据库是：" + dbType)
//
//	st := ""
//	fmt.Print(`请选择您使用的存储?
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//1: Local (Default)
//2: AliyunOSS
//3：TencentOSS
//4: Qiniu
//- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//Select the appropriate number [1-4] then [enter] (press 'c' to cancel): `)
//	_, _ = fmt.Fscanln(os.Stdin, &st)
//
//	if len(st) == 0 || st == "" {
//		storageType = StorageTypeName(1)
//	} else {
//		if len(st) == 1 {
//			a, _ := strconv.ParseInt(st, 10, 32)
//			switch a {
//			case 1:
//				storageType = StorageTypeName(1)
//				break
//			case 2:
//				storageType = StorageTypeName(2)
//				break
//			case 3:
//				storageType = StorageTypeName(3)
//				break
//			case 4:
//				storageType = StorageTypeName(4)
//				break
//			default:
//				l.Fatalf("un support storage type: %s", a)
//				os.Exit(1)
//			}
//		}
//	}
//
//	for i := 0; i < 8; i++ {
//		fmt.Print("\033[1A")
//		fmt.Print("\033[2K")
//	}
//	fmt.Println("您选择的存储是：" + storageType)
//}
