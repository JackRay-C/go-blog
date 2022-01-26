package domain

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Subject struct {
	ID          int            `json:"id" gorm:"type:int;primary_key;auto_increment;comment:'主键ID'"` // 主键
	Title       string         `json:"title" gorm:"type:varchar(255);uniqueIndex;not null"`          // 专题题目
	Avatar      int            `json:"image" gorm:"type:bigint;comment:'头像，文件id'"`                   // 头像
	CoverImage  int            `json:"cover_image" gorm:"type:bigint;default:1;comment:背景图，文件id"`    // 背景图
	Description string         `json:"description" gorm:"type:text"`                                 // 描述
	Visibility  int            `json:"visibility" gorm:"type:tinyint;default:1;comment:私有1，公开2"`     // 是否公开，1私有、2公开                                                   // 是否公开
	UserID      int            `json:"user_id" gorm:"type:int;index:idx_user_id"`                    // 所属用户
	Views       int            `json:"views" gorm:"type:int;default:0"`                              // 阅读次数
	CreatedAt   time.Time      `json:"created_at"`                                                   // 创建时间
	UpdatedAt   time.Time      `json:"updated_at"`                                                   // 更新时间
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`                                                   // 删除时间
}

func (s *Subject) String() string {
	marshal, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (s *Subject) TableName() string {
	return "subjects"
}
