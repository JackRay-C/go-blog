package domain

import (
	"blog/core/global"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Settings struct {
	ID        int            `json:"id" gorm:"type:int;primary_key;auto_increment;comment:主键ID"`
	Group     string         `json:"group" gorm:"type:varchar(255);uniqueIndex"`
	Key       string         `json:"key" gorm:"type:varchar(255)"`
	Value     string         `json:"value" gorm:"type:varchar(255)"`
	Flags     int8         `json:"flags" gorm:"type:tinyint;comment:是否公开的标志"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"delete_at"`
}

func (s *Settings) TableName() string {
	return "settings"
}

func (s *Settings) String() string {
	marshal, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (s *Settings) Select() error {
	return global.DB.Model(s).Where(s).First(s).Error
}

func (s *Settings) Count(count *int64) error {
	return global.DB.Model(s).Where(s).Count(count).Error
}

func (s *Settings) InsertAll(settings ...*Settings) error {
	return global.DB.Model(s).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(settings, 1000).Error
}

func (s *Settings) Update() error {
	return global.DB.Model(s).Updates(s).Error
}

func (s *Settings) Save() error {
	return global.DB.Model(s).Omit("created_at").Save(s).Error
}
