package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Settings struct {
	ID        int            `json:"id" gorm:"type:int;primary_key;auto_increment;common:主键ID"`
	Group     string         `json:"group" gorm:"type:varchar(255);uniqueIndex"`
	Key       string         `json:"key" gorm:"type:varchar(255)"`
	Value     string         `json:"value" gorm:"type:varchar(255)"`
	Flags     int8         `json:"flags" gorm:"type:tinyint;common:是否公开的标志"`
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