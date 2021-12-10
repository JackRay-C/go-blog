package domain

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type PostType int

const (
	Public PostType = iota + 1 // 发布
	Draft                      // 草稿箱
)

type Post struct {
	ID              int            `json:"id" gorm:"type:int;primary_key;autoIncrement;comment:主键ID;" form:"id"`
	Title           string         `json:"title" gorm:"type:varchar(100);default:'文章标题';not null;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"  form:"title"`
	MarkdownContent string         `json:"markdown_content" gorm:"type:longtext;comment:文章内容;;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	HtmlContent     string         `json:"html_content" gorm:"type:longtext;comment:文章html内容;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	CoverImageId    int            `json:"cover_image_id" gorm:"type:int;comment:封面图片"`
	Description     string         `json:"description" gorm:"type:text;comment:文章描述;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	Visibility      int            `json:"visibility" gorm:"type:tinyint;default:1;comment:私有1 公开2" form:"visibility"`
	Status          int            `json:"status" gorm:"type:tinyint;default:1;comment:草稿1，发布2" form:"status"`
	UserId          int            `json:"user_id" gorm:"type:int" form:"user_id"`
	SubjectId       int            `json:"subject_id" gorm:"type:int" form:"subject_id"`
	ImageIds        string         `json:"image_ids" gorm:"type:varchar(255);comment:图片id列表"`
	Views           int            `json:"views" gorm:"type:int;default:0"`
	Likes           int            `json:"likes" gorm:"type:int;default:0"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	PublishedAt     time.Time      `json:"published_at" gorm:"type:datetime;default:null"`
	DeleteAt        gorm.DeletedAt `json:"delete_at"`
}

func (p *Post) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*Post) TableName() string {
	return "posts"
}