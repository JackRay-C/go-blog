package domain

import (
	"blog/core/global"
	"encoding/json"
	"gorm.io/gorm/clause"
)

type Dict struct {
	ID          int    `json:"id" gorm:"type:int;primary_key;auto_increment;comment:主键ID"`
	Name        string `json:"name" gorm:"type:varchar(255);index:idx_type_code,unique;"`
	Code        int    `json:"code" gorm:"type:int;index:idx_type_code,unique;"`
	Value       string `json:"value" gorm:"type:varchar(255);index:idx_type_code,unique;"`
	Description string `json:"description" gorm:"type:varchar(255)"`
}

func (d *Dict) TableName() string  {
	return "dicts"
}

func (d *Dict) String() string {
	marshal, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (d *Dict) Select() error {
	return global.DB.Model(d).Where(d).First(d).Error
}

func (d *Dict) List(list *[]*Dict, offset int, limit int) error {
	return global.DB.Model(d).Where(d).Offset(offset).Limit(limit).Find(list).Error
}

func (d *Dict) Insert() error {
	return global.DB.Model(d).Create(d).Error
}

func (d *Dict) InsertAll(dicts []Dict) error {
	return global.DB.Model(d).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dicts, 1000).Error
}

func (d *Dict) Save() error {
	return global.DB.Save(d).Error
}

func (d *Dict) Update() error {
	return global.DB.Model(d).Updates(d).Error
}

func (d *Dict) Delete() error {
	return global.DB.Where(d).Delete(d).Error
}

func (d *Dict) DeleteIds(ids []int) error {
	return global.DB.Delete(d, ids).Error
}

func (d *Dict) Count(count *int64) error {
	return global.DB.Debug().Model(d).Where(d).Count(count).Error
}