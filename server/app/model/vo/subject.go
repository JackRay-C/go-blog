package vo

import (
	"blog/app/model/po"
	"encoding/json"
	"time"
)

type VSubject struct {
	ID          int          `json:"id" `          // 主键
	Title       string       `json:"title" `       // 专题题目
	Avatar      *po.File `json:"image"`        // 头像
	CoverImage  *po.File `json:"cover_image"`  // 背景图
	Description string       `json:"description" ` // 描述
	Visibility  int          `json:"visibility"`   // 是否公开
	UserID      int          `json:"user_id"`      // 所属用户
	User        *VUser       `json:"user"`         // 用户
	Views       int          `json:"views"`        // 阅读次数
	CreatedAt   time.Time    `json:"created_at"`
}

func (l *VSubject) string() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}
