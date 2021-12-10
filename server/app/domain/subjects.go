package domain

import (
	"blog/core/global"
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
	UserID      int            `json:"user_id"`                                                      // 所属用户
	Views       int            `json:"views" gorm:"type:int;default:0"`                              // 阅读次数
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
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

func (s *Subject) Select() error {
	return global.DB.Model(s).Where(s).First(s).Error
}

func (s *Subject) List(list *[]Subject, offset int, limit int) error {
	return global.DB.Model(s).Where(s).Offset(offset).Limit(limit).Find(list).Error
}

func (s *Subject) Insert() error {
	return global.DB.Model(s).Create(s).Error
}

func (s *Subject) Save() error {
	return global.DB.Omit("created_at", "delete_at").Save(s).Error
}

func (s *Subject) Update() error {
	return global.DB.Model(s).Updates(s).Error
}

func (s *Subject) Delete() error {
	return global.DB.Model(s).Where("ID=?", s.ID).Delete(s).Error
}
func (s *Subject) DeleteIds(ids []int) error {
	return global.DB.Model(s).Delete(s, ids).Error
}

func (s *Subject) Count(count *int64) error {
	return global.DB.Debug().Model(s).Where(s).Count(count).Error
}
