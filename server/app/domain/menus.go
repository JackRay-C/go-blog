package domain

import (
	"blog/core/global"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Menu struct {
	ID        int    `json:"id" gorm:"type:int;primary_key;auto_increment;comment:主键ID;"`
	Name      string `json:"name" gorm:"type:varchar(100);"`
	Path      string `json:"path" gorm:"type:varchar(100);index:idx_path,unique;"`
	Component string `json:"component" gorm:"type:varchar(255);"`
	Meta      string `json:"meta" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *Menu) String() string {
	marshal, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*Menu) TableName() string {
	return "menus"
}

func (m *Menu) Select() error {
	return global.DB.Model(m).Where(m).First(m).Error
}

func (m *Menu) List(list *[]Menu, offset int, limit int) error {
	return global.DB.Model(m).Where(m).Offset(offset).Limit(limit).Find(list).Error
}

func (m *Menu) Insert() error {
	return global.DB.Create(m).Error
}

func (m *Menu) InsertAll(menus ...*Menu) error {
	return global.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(menus, 1000).Error
}

func (m *Menu) Save() error {
	return global.DB.Save(m).Error
}

func (m *Menu) Update() error {
	if err := global.DB.Model(m).Where("ID=?", m.ID).First(m).Error; err != nil {
		return err
	}
	return global.DB.Model(m).Updates(m).Error
}

func (m *Menu) Delete() error {
	if err := global.DB.Model(m).Where("ID=?", m.ID).First(m).Error; err != nil {
		return err
	}
	return global.DB.Delete(m).Error
}

func (m *Menu) DeleteIds(ids []int) error {
	return global.DB.Delete(m, ids).Error
}

func (m *Menu) Count(count *int64) error {
	return global.DB.Model(m).Where(m).Count(count).Error
}
