package database

import (
	"blog/internal/config"
	"gorm.io/gorm"
)

const (
	Mysql = "mysql"
	Sqlite3 = "sqlite3"
	MongoDB = "mongodb"
)

func New(setting *config.App) (*gorm.DB, error)  {
	switch setting.AppDatabaseType {
	case Mysql:
		return NewMysqlEngine(setting.Database.Mysql)
	case MongoDB:
		return NewMongoDBEngine(setting.Database.MongoDB)
	case Sqlite3:
		return NewSqlite3Engine(setting.Database.Sqlite3)
	default:
		return NewMysqlEngine(setting.Database.Mysql)
	}
}
