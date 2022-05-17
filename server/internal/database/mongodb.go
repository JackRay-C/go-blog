package database

import (
	"blog/internal/config"
	"gorm.io/gorm"
)

func NewMongoDBEngine(mongo *config.MongoDB) (*gorm.DB, error) {
	panic("implement me.")
}
