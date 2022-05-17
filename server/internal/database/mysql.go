package database

import (
	"blog/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysqlEngine(mysqlSetting *config.Mysql) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(mysqlDsn(mysqlSetting)), &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		PrepareStmt:            true,
		SkipDefaultTransaction: false,
		Logger:                 mysqlLogMode(mysqlSetting.LogMode),
		FullSaveAssociations:   true,
	})
}

func mysqlLogMode(logMode string) logger.Interface {
	var mode logger.Interface
	switch logMode {
	case "debug":
		mode = logger.Default.LogMode(logger.Info)
	case "error":
		mode = logger.Default.LogMode(logger.Error)
	default:
		mode = logger.Default.LogMode(logger.Warn)
	}
	return mode
}

func mysqlDsn(mysql *config.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.DbName, mysql.Charset)
}
