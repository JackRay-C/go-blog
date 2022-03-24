package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id" gorm:"type:int;primary_key;auto_increment"`
	Username  string         `json:"username"  gorm:"type:varchar(255);uniqueIndex"`
	Nickname  string         `json:"nickname"  gorm:"type:varchar(255);uniqueIndex"`
	Password  string         `json:"password"  gorm:"type:varchar(255)"`
	Active    int8           `json:"active" gorm:"type:tinyint;default:1;common:'1激活，2锁定'"`
	Email     string         `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Avatar    int            `json:"avatar" gorm:"type:int"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}


func (*User) TableName() string {
	return "users"
}

func (u *User) String() string {
	marshal, _ := json.Marshal(u)
	return string(marshal)
}
