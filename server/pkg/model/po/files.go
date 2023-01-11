package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

const (
	DefaultPostCoverImageId = iota
	DefaultUserAvatarId
	DefaultSubjectCoverImageId
	DefaultSubjectAvatarId
	DefaultTagAvatarId
	DefaultTagCoverImageId
)

type File struct {
	ID        int64          `json:"id" gorm:"type:int;primary_key;auto_increment;common:主键ID;"`
	Name      string         `json:"name" gorm:"type:varchar(255);index:idx_filename,unique;"`
	Ext       string         `json:"ext" gorm:"type:varchar(50);index:idx_filename,unique;common:文件后缀"`
	AccessUrl string         `json:"access_url"`
	UserID    int64          `json:"user_id" grom:"type:int"`
	Md5       string         `json:"md_5" gorm:"type:varchar(32)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
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
