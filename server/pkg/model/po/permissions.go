package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Permissions struct {
	ID          int64            `json:"id" gorm:"type:int;primary_key;autoIncrement;common:主键ID;"`
	Name        string         `json:"name" gorm:"type:varchar(255);index:idx_name,unique;"`
	ObjectType  string         `json:"object_type" gorm:"type:varchar(255);index:idx_p,unique;"`
	ActionType  string         `json:"action_type" gorm:"type:varchar(255);index:idx_p,unique;"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"delete_at"`
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