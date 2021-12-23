package domain

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	ID          int            `json:"id" gorm:"type:int;primary_key;auto_increment;comment:'主键ID'"`
	Name        string         `json:"name" gorm:"type:varchar(255);uniqueIndex;comment:'标签名称'"`
	//Slug        string         `json:"slug" gorm:"type:varchar(255);uniqueIndex;comment:slug"`
	UserId      int            `json:"user_id" gorm:"type:int"`
	CoverImage  int            `json:"cover_image" gorm:"type:int;default:1;comment:背景图，文件id"` // 背景图
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (t *Tag) String() string {
	marshal, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*Tag) TableName() string {
	return "tags"
}