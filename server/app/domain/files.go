package domain

import (
	"blog/core/global"
	"blog/core/storage"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type File struct {
	ID        int              `json:"id" gorm:"type:int;primary_key;auto_increment;comment:主键ID;"`
	Name      string           `json:"name" gorm:"type:varchar(255);index:idx_filename,unique;"`
	Type      storage.FileType `json:"type" gorm:"type:tinyint(2);"`
	Ext       string           `json:"ext" gorm:"type:varchar(50);index:idx_filename,unique;comment:文件后缀"`
	Host      string           `json:"host" gorm:"type:varchar(255)"`
	AccessUrl string           `json:"access_url" gorm:"type:varchar(255)"`
	Path      string           `json:"path" gorm:"type:varchar(255)"`
	UserID    int              `json:"user_id" grom:"type:int"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
}

func (i *File) TableName() string {
	return "files"
}

func (i *File) String() string {
	marshal, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (i *File) Select() error {
	return global.DB.Model(i).Where(i).First(i).Error
}

func (i *File) List(list *[]File, offset int, limit int) error {
	return global.DB.Model(i).Where(i).Offset(offset).Limit(limit).Find(list).Error
}

func (i *File) Insert() error {
	return global.DB.Model(i).Clauses(clause.OnConflict{DoNothing: true}).Create(i).Error
}

func (i *File) InsertAll(files []File) error {
	return global.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(files, 1000).Error
}

func (i *File) Save() error {
	return global.DB.Save(i).Error
}

func (i *File) Update(condition map[string]interface{}) error {
	if err := global.DB.Model(i).Where("ID=?", i.ID).First(i).Error; err != nil {
		return err
	}
	return global.DB.Model(i).Updates(condition).Error
}

func (i *File) Delete() error {
	if err := global.DB.Model(i).Where("ID=?", i.ID).First(i).Error; err != nil {
		return err
	}
	return global.DB.Delete(i).Error
}

func (i *File) DeleteIds(ids []int) error {
	return global.DB.Delete(i, ids).Error
}

func (i *File) Count(count *int64) error {
	return global.DB.Model(i).Where(i).Count(count).Error
}
