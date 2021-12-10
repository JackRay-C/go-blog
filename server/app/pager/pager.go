package pager

import (
	"encoding/json"
)

type Pager struct {
	PageNo    int         `json:"page_no" form:"page_no"`
	PageSize  int         `json:"page_size" form:"page_size"`
	PageCount int         `json:"total_page"`
	TotalRows int64       `json:"total_rows"`
	List      interface{} `json:"list"`
}

func (p *Pager) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}
