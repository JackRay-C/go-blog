package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID          int64          `json:"id" gorm:"type:int;primary_key;auto_increment;common:'主键ID'"`
	Name        string         `json:"name" gorm:"type:varchar(100);index:idx_name,unique;"`
	Description string         `json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"delete_at"`
}


func (i *Role) TableName() string {
	return "roles"
}

func (i *Role) String() string {
	marshal, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(marshal)
}

const (
	ROLE_ADMIN = iota + 1
	ROLE_EDITOR
	ROLE_VIEWER
)
