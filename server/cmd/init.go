package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"time"
)

var (
	initCommand = &cobra.Command{
		Use:   "init [flags]",
		Short: "Init go blog directory",
		Run: func(cmd *cobra.Command, args []string) {
			// 初始化viper配置
			initViperDefaultConfiguration()
			// 写入配置文件
			writeViperDefaultConfigurationToInitPath()
		},
	}
	initPath       string = ""
	configFileName        = ".go-log"
	configFileType        = "yaml"
)

func init() {
	rootCmd.AddCommand(initCommand)

	initCommand.AddCommand(mysqlCommand)
	initCommand.AddCommand(storageCommand)
	initCommand.AddCommand(serverCommand)
	initCommand.AddCommand(logCommand)
	initCommand.AddCommand(dbCommand)

	initCommand.Flags().StringVarP(&initPath, "directory", "d", "", "init directory")
	initCommand.Flags().StringVar(&configFileName, "config-file-name", ".go-blog", "If noe-empty, config file name is this")
	initCommand.Flags().StringVar(&configFileType, "config-file-type", "yaml", "viper config file type")
}

func initViperDefaultConfiguration() {
	initViperServerDefaultConfiguration()
	initViperAppDefaultConfiguration()
	initViperSnowflakeDefaultConfiguration()
	initViperMysqlDefaultConfiguration()
	initViperMongodbDefaultConfiguration()
	initViperStorageDefaultConfiguration()
	initViperZapDefaultConfiguration()
	initViperSmtpDefaultConfiguration()
}

func initViperServerDefaultConfiguration() {
	viper.SetDefault("server.port", serverPort)
	viper.SetDefault("server.readTimeout", serverReadTimeout)
	viper.SetDefault("server.writeTimeout", serverWriteTimeout)
}
func initViperAppDefaultConfiguration() {
	viper.SetDefault("app.name", "go-blog")
	viper.SetDefault("app.version", "v1.0.0")
	viper.SetDefault("app.run-mode", "release")
	viper.SetDefault("app.db-type", dbType)
	viper.SetDefault("app.storage-type", storageType)
	viper.SetDefault("app.upload-max-size", uploadMaxSize)
	viper.SetDefault("app.upload-allow-exts", []string{
		".jpg", ".jpeg", ".png", ".bmp", ".icon", ".gif", ".psd",
		".mp3", ".wav", ".mid", ".flac", ".mp4", ".mov", ".avi", ".flv", ".m4v", ".rmvb",
		".zip", ".tar", ".rar", ".gz", ".bz2",
		".docx", ".xlsx", ".pdf", ".ppt", ".pptx", ".doc", "xls",
		".md", ".html", ".js", ".css", ".java", ".class", ".py", ".go", ".sh", ".logs", ".yaml", ".yml", ".json", ".ini",
		".xmind", ".txt",
	})
	viper.SetDefault("app.access-token-expire", 60*time.Second)
	viper.SetDefault("app.refresh-token-expire", 60*7*24)
}

func initViperSnowflakeDefaultConfiguration() {
	viper.SetDefault("snowflake.work-id", 1)
	viper.SetDefault("snowflake.data-center-id", 1)
}

func initViperZapDefaultConfiguration() {
	viper.SetDefault("zap.level", logging.level)
	viper.SetDefault("zap.format", logging.format)
	viper.SetDefault("zap.directory", logging.directory)
	viper.SetDefault("zap.linkName", logging.filename)
	viper.SetDefault("zap.showLine", logging.showLine)
	viper.SetDefault("zap.encodeLevel", logging.encodeLevel)
	viper.SetDefault("zap.stacktraceKey", logging.stacktraceKey)
	viper.SetDefault("zap.logInConsole", logging.logInConsole)
	viper.SetDefault("zap.logMaxSize", logging.logMaxSize)
	viper.SetDefault("zap.logMaxBackups", logging.logMaxBackups)
	viper.SetDefault("zap.logMaxAge", logging.logMaxAge)

}

func initViperMysqlDefaultConfiguration() {
	viper.SetDefault("mysql.username", mysql.Username)
	viper.SetDefault("mysql.password", mysql.Password)
	viper.SetDefault("mysql.host", mysql.Host)
	viper.SetDefault("mysql.port", mysql.Port)
	viper.SetDefault("mysql.db-name", mysql.DbName)
	viper.SetDefault("mysql.parse-time", mysql.ParseTime)
	viper.SetDefault("mysql.charset", mysql.Charset)
	viper.SetDefault("mysql.log-mode", mysql.LogMode)
	viper.SetDefault("mysql.max-idle-conns", mysql.MaxIdleConns)
	viper.SetDefault("mysql.max-open-conns", mysql.MaxOpenConns)
}

func initViperMongodbDefaultConfiguration() {
	viper.SetDefault("mongodb.username", "")
	viper.SetDefault("mongodb.password", "")
	viper.SetDefault("mongodb.host", "")
	viper.SetDefault("mongodb.port", "")
	viper.SetDefault("mongodb.db-name", "")

}

func initViperSmtpDefaultConfiguration() {
	viper.SetDefault("smtp.host", "")
	viper.SetDefault("smtp.port", 564)
	viper.SetDefault("smtp.username", "")
	viper.SetDefault("smtp.password", "")
	viper.SetDefault("smtp.isSSL", "")
	viper.SetDefault("smtp.form", "")
}

func initViperStorageDefaultConfiguration() {
	viper.SetDefault("local.path", "static/uploads")

	viper.SetDefault("qiniu.zone", "zone")
	viper.SetDefault("qiniu.bucket", "")

	viper.SetDefault("aliyun-oss.endpoint", "")

	viper.SetDefault("tencent-oss.bucket", "")
}

func writeViperDefaultConfigurationToInitPath() {
	if initPath == "" {
		homeDir, _ := os.UserHomeDir()
		initPath = homeDir
	}






	// 1、判断文件是否存在
	f := path.Join(initPath, configFileName + "." + configFileType)
	_, err := os.Stat(f)

	// 2、不存在的话直接写入
	if os.IsNotExist(err) {
		viper.AddConfigPath(initPath)
		viper.SetConfigName(configFileName)
		viper.SetConfigType(configFileType)
		if err := viper.SafeWriteConfig(); err != nil {
			log.Fatalln(err)
		}
		_, _ = fmt.Fprintf(os.Stdout, "generate default config file to [%s]", f)
		return
	}

	// 3、存在报错
	_, _ = fmt.Fprintf(os.Stderr, "config file \"%s\" already exists", f)
}
