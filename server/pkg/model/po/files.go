package po

import (
	"blog/internal/storage"
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
	ID        int              `json:"id" gorm:"type:int;primary_key;auto_increment;common:主键ID;"`
	Name      string           `json:"name" gorm:"type:varchar(255);index:idx_filename,unique;"`
	Type      storage.FileType `json:"type" gorm:"type:tinyint(2);"`
	Ext       string           `json:"ext" gorm:"type:varchar(50);index:idx_filename,unique;common:文件后缀"`
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
