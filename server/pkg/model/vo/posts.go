package vo

import (
	"blog/pkg/model/po"
	"encoding/json"
)

type VPosts struct {
	ID              int       `json:"id"`                // 更新post id
	Title           string    `json:"title"`             // 标题
	MarkdownContent string    `json:"markdown_content" ` // markdown
	HtmlContent     string    `json:"html_content" `     // html
	CoverImageId    int       `json:"cover_image_id"`    // 封面图片id
	CoverImage      *po.File  `json:"cover_image"`       // 封面图片
	Description     string    `json:"description"`       // 描述
	Visibility      int       `json:"visibility"`        // 1、私有 2、公开
	Status          int       `json:"status"`            // 1、草稿 2、发布
	SubjectId       int       `json:"subject_id"`        // 专题ID
	Subject         *VSubject `json:"subject"`           // 专题
	ImageIds        string    `json:"image_ids"`         // 所有图片的列表
	Tags            []*po.Tag `json:"tags"`              // 所有标签的列表
	Likes           int       `json:"likes"`             // 点赞
	Views           int       `json:"views"`             // 阅读量
	UserId          int       `json:"user_id"`           // 用户ID
	User            *VUser    `json:"user"`              // 用户
	CreatedAt       *int64    `json:"created_at"`        // 创建时间
	UpdatedAt       *int64    `json:"updated_at"`        // 更新时间
	PublishedAt     *int64    `json:"published_at"`      // 发布时间
}

func (l *VPosts) string() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}
