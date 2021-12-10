package mysql

import (
	"blog/core/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func NewEngine(mysqlSetting *setting.Mysql) (*gorm.DB, error) {
	dsn := getDsn(mysqlSetting)
	logMode := getDbLogMode(mysqlSetting)
	db, err := getDb(dsn, logMode)
	if err != nil {
		return nil, err
	}

	if s, err := db.DB(); err != nil {
		return nil, err
	} else {
		s.SetMaxIdleConns(mysqlSetting.MaxIdleConns)
		s.SetMaxOpenConns(mysqlSetting.MaxOpenConns)
		s.SetConnMaxLifetime(20 * time.Hour)
	}
	return db, nil
}

func Close(db *gorm.DB) {
	s, _ := db.DB()
	s.Close()
}

func getDb(dsn string, logMode logger.Interface) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{SingularTable: true}, // 单数表名，不加s
		PrepareStmt:            true,
		SkipDefaultTransaction: false,
		Logger:                 logMode,
		FullSaveAssociations:   true,
	})

	return db, err
}

func getDbLogMode(setting *setting.Mysql) logger.Interface {
	var logMode logger.Interface
	switch setting.LogMode {
	case "debug":
		logMode = logger.Default.LogMode(logger.Info)
	case "error":
		logMode = logger.Default.LogMode(logger.Error)
	default:
		logMode = logger.Default.LogMode(logger.Warn)
	}
	return logMode
}

func getDsn(mysql *setting.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.DbName, mysql.Charset)
}
