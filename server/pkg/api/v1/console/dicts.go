package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Dict struct {
	log         logger.Logger
	service 	common.BaseService
}

func NewDict() *Dict {
	return &Dict{service: &common.BaseServiceImpl{}, log: global.Log}
}

func (d *Dict) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}
	d.log.Infof("Select one dict by id: %d", id)

	dict := po.Dict{ID: id}
	if err = d.service.ISelectOne(c, &dict); err != nil {
		d.log.Errorf("Dict get by ID: %d, error: %s", id, err)
		return nil, err
	}

	return vo.Success(&dict), nil
}

func (d *Dict) List(c *gin.Context) (*vo.Response, error) {
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	d.log.Infof("select dict with page. ")

	var dict *po.Dict
	name := c.Query("name")

	if name == "" {
		dict = &po.Dict{}
	} else {
		dict = &po.Dict{Name: name}
	}

	if err := d.service.ISelectList(c, &p, dict); err != nil {
		d.log.Errorf("list dict failed service error: %s", err)
		return nil, err
	}

	d.log.Infof("list dict success. ")
	return vo.Success(&p), nil
}

func (d *Dict) Post(c *gin.Context) (*vo.Response, error) {
	d.log.Infof("create dict. ")

	dict := &po.Dict{}
	if err := c.ShouldBindJSON(&dict); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := d.service.ICreateOne(c, dict); err != nil {
		d.log.Errorf("create dict failed: %s", err)
		return nil, err
	}

	d.log.Infof("create dict success. ")
	return vo.Success(dict), nil
}

func (d *Dict) Delete(c *gin.Context) (*vo.Response, error) {
	d.log.Infof("delete dict. ")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := d.service.IDeleteOne(c, &po.Dict{ID: id}); err != nil {
		return nil, err
	}
	return vo.Success("delete success. "), nil
}



func (d *Dict) Put(c *gin.Context) (*vo.Response, error) {
	d.log.Infof("update dict. ")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	var dict po.Dict
	if err := c.ShouldBindJSON(&dict); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	dict.ID = id
	if err := d.service.IUpdateOne(c, &dict, dict); err != nil {
		return nil, err
	}

	d.log.Infof("update dict success.")
	return vo.Success(&dict), nil
}
