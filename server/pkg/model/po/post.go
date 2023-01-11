package po

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Post struct {
	ID              int64          `json:"id" gorm:"type:int;primary_key;auto_increment;comment:主键ID;" form:"id"` // 主键ID，博客的唯一ID
	Title           string         `json:"title" gorm:"type:varchar(100);default:'文章标题';not null;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"  form:"title"`
	Slug            string         `json:"slug" gorm:"type:varchar(255)" form:"slug"`
	MarkdownContent string         `json:"markdown_content" gorm:"type:longtext;common:文章内容;;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	HtmlContent     string         `json:"html_content" gorm:"type:longtext;common:文章html内容;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	CoverImageId    int64          `json:"cover_image_id" gorm:"type:int;common:封面图片"`
	Description     string         `json:"description" gorm:"type:text;common:文章描述;index:idx_fulltext,class:FULLTEXT,option:WITH PARSER ngram ;"`
	ImageIds        string         `json:"image_ids" gorm:"type:varchar(255);common:图片id列表"`
	Visibility      int            `json:"visibility" gorm:"type:tinyint(2);default:1" form:"visibility"`        // 是否公开 1、私有 2、公开
	SubjectID       int64          `json:"subject_id" gorm:"type:bigint;index:idx_subject_id" form:"subject_id"` // 专题ID
	Likes           int            `json:"likes"  gorm:"type:int;default:0"`                                     // 点赞
	Views           int            `json:"views" gorm:"type:int;default:0"`                                      // 阅读量
	Status          int            `json:"status" gorm:"type:tinyint;default:1" form:"status"`                   // 状态：1、草稿，2、发布，3、删除
	UserID          int64          `json:"user_id" gorm:"type:bigint;index:idx_user_id" form:"user_id"`          // 用户ID
	CreatedAt       *int64         `json:"created_at" gorm:"type:bigint;default:0"`                              // 创建时间
	UpdatedAt       *int64         `json:"updated_at" gorm:"type:bigint;default:0"`                              // 更新时间
	PublishedAt     *int64         `json:"published_at" gorm:"type:bigint;default:0"`                            // 发表时间
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`                                                           // 删除时间
}

func (i *Post) TableName() string {
	return "posts"
}

func (i *Post) String() string {
	marshal, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(marshal)
}
