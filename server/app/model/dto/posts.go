package dto

import (
	"blog/app/domain"
	"encoding/json"
	"time"
)

// list post列表参数
type ListPosts struct {
	PageNo     int    `json:"page_no" form:"page_no" binding:"gt=0" `
	PageSize   int    `json:"page_size" form:"page_size"`
	Visibility int    `json:"visibility" form:"visibility" `
	Status     int    `json:"status"  form:"status"`
	UserId     int    `json:"user_id" form:"user_id"`
	SubjectId  int    `json:"subject_id"  form:"subject_id"`
	OrderBy    int    `json:"order_by" form:"order_by"`
	TagId      int    `json:"tag_id" form:"tag_id"`
	Search     string `json:"search" form:"search"`
}

func (l *ListPosts) String() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}

// 添加post
type AddPosts struct {
	Title           string       `json:"title" binding:"required"`                                      // 标题
	MarkdownContent string       `json:"markdown_content" `                                             // markdown
	HtmlContent     string       `json:"html_content" `                                                 // html
	CoverImageId    int          `json:"cover_image_id"`                                                // 封面图片id
	Description     string       `json:"description"`                                                   // 描述
	Visibility      int          `json:"visibility" binding:"oneof=1 2"`                                // 1、私有 2、公开
	Status          int          `json:"status" binding:"oneof=1 2"`                                    // 1、草稿 2、发布
	SubjectId       int          `json:"subject_id"`                                                    // 专题ID
	ImageIds        string       `json:"image_ids"`                                                     // 所有图片的列表
	Tags            []domain.Tag `json:"tags"`                                                          // 所有标签的列表
	UserId          int          `binding:"-"`                                                          // 用户ID
	CreatedAt       time.Time    `json:"created_at" time_format:"2006-01-02T15:04:05.999999999Z07:00" ` // 创建时间
}

func (l *AddPosts) String() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}

// put posts
type PutPosts struct {
	ID              int           `json:"id" binding:"required,gte=1"`                                    // 更新post id
	Title           string        `json:"title"`                                       // 标题
	MarkdownContent string        `json:"markdown_content" `                                              // markdown
	HtmlContent     string        `json:"html_content" `                                                  // html
	CoverImageId    int           `json:"cover_image_id"`                                                 // 封面图片id
	Description     string        `json:"description"`                                                    // 描述
	Visibility      int           `json:"visibility" binding:"required,oneof=1 2"`                        // 1、私有 2、公开
	Status          int           `json:"status" binding:"required,oneof=1 2"`                            // 1、草稿 2、发布
	SubjectId       int           `json:"subject_id"`                                                     // 专题ID
	ImageIds        string        `json:"image_ids"`                                                      // 所有图片的列表
	Tags            []*domain.Tag `json:"tags"`                                                           // 所有标签的列表
	Likes           int           `json:"likes"`                                                          // 点赞
	Views           int           `json:"views"`                                                          // 阅读量
	UserId          int           `binding:"-"`                                                           // 用户ID，不允许前端传，后端根据session生成
	CreatedAt       time.Time     `json:"created_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"`   // 创建时间
	PublishedAt     time.Time     `json:"published_at" time_format:"2006-01-02T15:04:05.999999999Z07:00"` // 发布时间
}

func (l *PutPosts) String() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}

// patch 修改post某项参数
type PatchPosts struct {
	Title       string `json:"title"`
	SubjectId   int    `json:"subject_id"`
	Content     string `json:"content"`
	HtmlContent string `json:"html_content"`
	CoverImage  string `json:"cover_image"`
	Description string `json:"description"`
	Type        uint8  `json:"type"`
}

// 为某个专题添加posts
type AddOneSubjectsAllPosts struct {
	Posts []*domain.Post `json:"posts"`
}
