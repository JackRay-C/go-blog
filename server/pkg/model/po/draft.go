package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Draft struct {
	ID              int64          `json:"id" gorm:"type:int;primary_key;autoIncrement;common:主键ID;" form:"id"` // 主键ID，博客的唯一ID
	Title           string         `json:"title" gorm:"type:varchar(100);default:'文章标题';not null;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"  form:"title"`
	MarkdownContent string         `json:"markdown_content" gorm:"type:longtext;common:文章内容;;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	CoverImageId    int64          `json:"cover_image_id" gorm:"type:int;common:封面图片"`
	Description     string         `json:"description" gorm:"type:text;common:文章描述;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	ImageIds        string         `json:"image_ids" gorm:"type:varchar(255);common:图片id列表"`
	Visibility      int            `json:"visibility" gorm:"type:tinyint(2);default:1" form:"visibility"`     // 是否公开 1、私有 2、公开
	SubjectID       int64          `json:"subject_id" gorm:"type:int;index:idx_subject_id" form:"subject_id"` // 专题ID
	PostID          int64          `json:"post_id" gorm:"type:int;" form:"post_id"`
	UserID          int64          `json:"user_id" gorm:"type:int;index:idx_user_id" form:"user_id"`     // 用户ID
	CreatedAt       time.Time      `json:"created_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // 创建时间
	UpdatedAt       time.Time      `json:"updated_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // 更新时间
	DeletedAt       gorm.DeletedAt `json:"deleted_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // 删除时间
}

func (i *Draft) TableName() string {
	return "drafts"
}

func (i *Draft) String() string {
	marshal, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(marshal)
}
