package web

import (
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Dict struct {
	service common.BaseService
}

func NewDict() *Dict {
	return &Dict{service: &common.BaseServiceImpl{}}
}

func (d *Dict) Get(c *gin.Context) (*vo.Response, error)   {
	// 1、获取id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	dict := &po.Dict{ID: id}
	if err := d.service.ISelectOneWeb(c, dict); err != nil {
		return nil, err
	}
	return vo.Success(dict), nil
}

func (d *Dict) List(c *gin.Context) (*vo.Response, error)  {
	dict := &po.Dict{}
	if name, ok := c.GetQuery("name"); ok {
		dict.Name = name
	}

	pager := &vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}
	if err := d.service.ISelectListWeb(c, pager, dict); err != nil {
		return nil, err
	}

	return vo.Success(pager), nil
}