package domain

import (
	"blog/core/global"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	ID        int            `json:"id" gorm:"type:int;primary_key;autoIncrement;comment:主键ID;"`
	Comment   string         `json:"comment" gorm:"type:text;"`
	Email     string         `json:"email" gorm:"type:varchar(255)"`
	Nickname  string         `json:"nickname" gorm:"type:varchar(255)"`
	PostId    int            `json:"post_id" gorm:"type:int"`
	ParentID  int            `json:"parent_id" gorm:"type:int"`
	UserID    int            `json:"user_id" gorm:"type:int"`
	Child     []*Comment     `json:"child" gorm:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
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

func (c *Comment) Select() error {
	return global.DB.Model(c).Where(c).First(c).Error
}

func (c *Comment) List(list *[]Comment, offset int, limit int) error {
	return global.DB.Model(c).Offset(offset).Limit(limit).Find(list).Error
}

func (c *Comment) Insert() error {
	return global.DB.Create(c).Error
}

func (c *Comment) Save() error {
	return global.DB.Save(c).Error
}

func (c *Comment) Update() error {
	return global.DB.Debug().Model(c).Select("comment").Omit("*").Where("ID=?", c.ID).Updates(c).Error
}

func (c *Comment) Delete() error {
	return global.DB.Delete(c).Where(c).Error
}

func (c *Comment) DeleteIds(ids []int) error {
	return global.DB.Delete(c, ids).Error
}

func (c *Comment) Count(count *int64) error {
	return global.DB.Model(c).Count(count).Error
}

func (c *Comment) SelectAll(comments1 *[]*Comment) error {
	return global.DB.Model(c).Where(c).Order("created_at").Find(comments1).Error
}
