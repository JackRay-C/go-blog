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

func (p *Pager) MustList(list interface{})  {
	if p.TotalRows == 0 {
		p.PageCount = 0
		p.List = make([]string, 0)
	} else {
		p.PageCount = (p.PageCount + p.PageSize - 1)/p.PageSize
		p.List = list
	}
}
