package vo

import (
	"blog/pkg/utils/convert"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Pager struct {
	PageNo    int         `json:"page_no" form:"page_no"`
	PageSize  int         `json:"page_size" form:"page_size"`
	PageCount int         `json:"total_page"`
	TotalRows int64       `json:"total_rows"`
	List      interface{} `json:"list"`
	Search    string      `json:"search" form:"search"`
	SortBy    string      `json:"sort_by" form:"sort_by"`
	SortOrder string      `json:"sort_order" form:"sort_order"`
}

func (p *Pager) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (p *Pager) MustList(list interface{}) {
	if p.TotalRows == 0 {
		p.PageCount = 0
		p.List = make([]string, 0)
	} else {
		p.PageCount = (p.PageCount + p.PageSize - 1) / p.PageSize
		p.List = list
	}
}

// MustPageNo bind query page_no or set default value 1
func (p *Pager) MustPageNo(c *gin.Context)  {
	pageNo := c.DefaultQuery("page_no", "1")
	p.PageNo = convert.StrTo(pageNo).MustInt()
}

// MustPageSize bind query page_size or set default 10
func (p *Pager) MustPageSize(c *gin.Context) {
	pageSize := c.DefaultQuery("page_size", "10")
	p.PageSize =  convert.StrTo(pageSize).MustInt()
}

// MustSort bind sort meta from query or set default {"sort_by": "created_by", "sort_order":"desc"}
func (p *Pager) MustSort(c *gin.Context)  {
	p.SortBy = c.DefaultQuery("sort_by", "created_at")
	p.SortOrder = c.DefaultQuery("sort_order", "desc")
}

// MustSearch bind search query or set default ""
func (p *Pager) MustSearch(c *gin.Context)  {
	p.Search = c.DefaultQuery("search", "")
}