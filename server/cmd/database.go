package cmd

import (
	"github.com/spf13/cobra"
)



type mysqlT struct {
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

var (
	dbCommand = &cobra.Command{
		Use:   "db [options] [commands] [flags]",
		Short: "Configuration database information",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 0 {
				err := cmd.Usage()
				if err != nil {
					return 
				}
			}
		},
	}
	dbType       string
	mysqlCommand = &cobra.Command{
		Use:   "mysql [options] [flags]",
		Short: "Configuration mysql database information",
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			
		},
	}
	mongodbCommand = &cobra.Command{
		Use: "mongodb [flags]",
		Short: "not support this database",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	mysql mysqlT
)

func init() {
	dbCommand.AddCommand(mysqlCommand)
	dbCommand.AddCommand(mongodbCommand)

	dbCommand.Flags().StringVarP(&dbType, "db-type", "d","mysql", "database type")
	initMysqlFlags()
}

func initMysqlFlags() {
	mysql.Username = "root"
	mysql.Password = "root"
	mysql.Host = "127.0.0.1"
	mysql.Port = 3306
	mysql.DbName = "go_blog"
	mysql.ParseTime = true
	mysql.Charset = "utf8"
	mysql.MaxIdleConns = 10
	mysql.MaxOpenConns = 30
	mysql.LogMode = "info"

	mysqlCommand.Flags().StringVar(&mysql.Host, "host", "127.0.0.1", "mysql server host")
	mysqlCommand.Flags().StringVar(&mysql.Username, "username", "root", "mysql server username")
	mysqlCommand.Flags().StringVar(&mysql.Password, "password", "root", "mysql server password")
	mysqlCommand.Flags().IntVar(&mysql.Port, "port", 3306, "mysql server port")
	mysqlCommand.Flags().StringVar(&mysql.DbName, "db-name", "go_blog", "server database name")
	mysqlCommand.Flags().StringVar(&mysql.Charset, "charset", "utf8", "db connect charset")
	mysqlCommand.Flags().StringVar(&mysql.LogMode, "log-mode", "info", "db connect pool log mode")
	mysqlCommand.Flags().IntVar(&mysql.MaxIdleConns, "max-idle-conns", 10, "db connect pool idle conn")
	mysqlCommand.Flags().IntVar(&mysql.MaxOpenConns, "max-open-conns", 30, "db connect pool open conn")

}
