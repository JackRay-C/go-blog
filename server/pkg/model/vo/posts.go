package vo

import (
	"blog/pkg/model/po"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type VPosts struct {
	ID              int           `json:"id"`                                                             // 更新post id
	Title           string        `json:"title"`                                                          // 标题
	MarkdownContent string        `json:"markdown_content" `                                              // markdown
	HtmlContent     string        `json:"html_content" `                                                  // html
	CoverImageId    int           `json:"cover_image_id"`                                                 // 封面图片id
	CoverImage      *po.File  `json:"cover_image"`                                                    // 封面图片
	Description     string        `json:"description"`                                                    // 描述
	Visibility      int           `json:"visibility"`                                                     // 1、私有 2、公开
	Status          int           `json:"status"`                                                         // 1、草稿 2、发布
	SubjectId       int           `json:"subject_id"`                                                     // 专题ID
	Subject         *VSubject     `json:"subject"`                                                        // 专题
	ImageIds        string        `json:"image_ids"`                                                      // 所有图片的列表
	Tags            []*po.Tag `json:"tags"`                                                           // 所有标签的列表
	Likes           int           `json:"likes"`                                                          // 点赞
	Views           int           `json:"views"`                                                          // 阅读量
	UserId          int           `json:"user_id"`                                                        // 用户ID
	User            *VUser        `json:"user"`                                                           // 用户
	CreatedAt       time.Time     `json:"created_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`   // 创建时间
	UpdatedAt       time.Time     `json:"updated_at"  time_format:"2006-01-02T15:04:05.999999999Z07:00"`  // 更新时间
	PublishedAt     time.Time     `json:"published_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // 发布时间
}

func (l *VPosts) string() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type VHead struct {
	ID           int                `json:"id" form:"id"`                                                 // 主键ID，博客的唯一ID
	RepositoryID int                `json:"repository_id"  form:"repository_id"`                          // 当前存储库的id
	Repository   *po.Repository `json:"repository"`                                                   // 当前存储库内容
	History      []*po.History  `json:"history"`                                                      // 当前博客的历史记录
	Visibility   int                `json:"visibility"`                                                   // 是否公开 1、私有 2、公开
	Status       int                `json:"status"`                                                       // 博客状态 1、已暂存 2、已提交 3、已发布
	Likes        int                `json:"likes"`                                                        // 点赞量
	Views        int                `json:"views"`                                                        // 阅读量
	SubjectID    int                `json:"subject_id" `                                                  // 专题ID
	CoverImageID int                `json:"cover_image_id"`                                               // 封面图片ID
	UserID       int                `json:"user_id"`                                                      // 用户ID
	User         *VUser             `json:"user"`                                                         // 用户信息
	CreatedAt    time.Time          `json:"created_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // 创建时间
	UpdatedAt    time.Time          `json:"updated_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // 更新时间
	DeletedAt    gorm.DeletedAt     `json:"deleted_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // 删除
}

type VPost struct {
	ID           int            `json:"id"`
	Visibility   int            `json:"visibility"`
	Status       int            `json:"status"`
	Likes        int            `json:"likes"`
	Views        int            `json:"views"`
	SubjectID    int            `json:"subject_id"`
	CoverImageID int            `json:"cover_image_id"`
	UserID       int            `json:"user_id"`
	User         *VUser         `json:"user"`
	Tags         []*po.Tag  `json:"tags"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	*po.Repository
}
