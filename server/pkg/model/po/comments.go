package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        int             `json:"id" gorm:"type:int;primary_key;autoIncrement;common:主键ID;"`
	Comment   string          `json:"comment" gorm:"type:text;"`
	Email     string          `json:"email" gorm:"type:varchar(255)"`
	Nickname  string          `json:"nickname" gorm:"type:varchar(255)"`
	PostId    int             `json:"post_id" gorm:"type:int"`
	ParentID  int             `json:"parent_id" gorm:"type:int"`
	UserID    int             `json:"user_id" gorm:"type:int"`
	Child     []*Comment      `json:"child" gorm:"-"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

func (c *Comment) String() string {
	marshal, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (c *Comment) TableName() string {
	return "comments"
}
