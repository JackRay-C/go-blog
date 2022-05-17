package bo

import (
	"blog/pkg/model/po"
	"encoding/json"
)

type Post struct {
	Head         *po.Head         `json:"head"`
	Repositories []*po.Repository `json:"repositories"`
	Histories    []*po.History    `json:"histories"`
}

func (p *Post) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}
