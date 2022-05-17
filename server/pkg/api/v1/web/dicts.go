package web

import (
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Dict struct {
	dictService service.DictService
}

func NewDict() *Dict {
	return &Dict{dictService: service.NewDictService()}
}

func (d *Dict) Get(c *gin.Context) (*vo.Response, error)   {
	// 1、获取id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	dict := &po.Dict{ID: id}
	if err := d.dictService.ISelectOne(c, dict); err != nil {
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
	if err := d.dictService.ISelectList(c, pager, dict); err != nil {
		return nil, err
	}

	return vo.Success(pager), nil
}