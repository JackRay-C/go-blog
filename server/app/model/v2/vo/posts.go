package vo

import (
	"blog/app/domain"
	"encoding/json"
)

type Posts struct {
	Head       *domain.Head       `json:"head"`       // 博客的指针头
	History    []*domain.History  `json:"history"`    // 博客的历史库
	Repository *domain.Repository `json:"repository"` // 博客的原始内容
}

func (v *Posts) string() string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(marshal)
}
