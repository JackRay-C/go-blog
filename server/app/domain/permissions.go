package domain

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Permissions struct {
	ID          int    `json:"id" gorm:"type:int;primary_key;autoIncrement;comment:主键ID;"`
	Name        string `json:"name"`
	ObjectType  string `json:"object_type"`
	ActionType  string `json:"action_type"`
	Description string `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeleteAt        gorm.DeletedAt `json:"delete_at"`
}

func (p *Permissions) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (p *Permissions) TableName() string {
	return "permissions"
}
