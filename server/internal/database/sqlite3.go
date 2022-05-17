package database

import (
	"blog/internal/config"
	"gorm.io/gorm"
)

func NewSqlite3Engine(sqlite3 *config.Sqlite3) (*gorm.DB, error) {
	panic("implement me")
}
