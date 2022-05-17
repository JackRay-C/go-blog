package po

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	ID              int            `json:"id" gorm:"type:int;primary_key;autoIncrement;common:主键ID;" form:"id"`
	Title           string         `json:"title" gorm:"type:varchar(100);default:'文章标题';not null;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"  form:"title"`
	MarkdownContent string         `json:"markdown_content" gorm:"type:longtext;common:文章内容;;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	HtmlContent     string         `json:"html_content" gorm:"type:longtext;common:文章html内容;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	CoverImageId    int            `json:"cover_image_id" gorm:"type:int;common:封面图片"`
	Description     string         `json:"description" gorm:"type:text;common:文章描述;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	UserId          int            `json:"user_id" gorm:"type:int;index:idx_user_id" form:"user_id"`
	ImageIds        string         `json:"image_ids" gorm:"type:varchar(255);common:图片id列表"`
	CreatedAt       time.Time      `json:"created_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`
	UpdatedAt       time.Time      `json:"updated_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`
	DeleteAt        gorm.DeletedAt `json:"delete_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`
}

func (r *Repository) String() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*Repository) TableName() string {
	return "repositories"
}