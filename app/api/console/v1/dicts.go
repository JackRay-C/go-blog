package v1

import (
	"blog/app/model/po"
	"blog/app/model/vo"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Dict struct {
	log logger.Logger
	dictService service.DictService
}

func NewDict() *Dict {
	return &Dict{dictService: service.NewDictService(), log: global.Logger}
}

func (d *Dict) Get(c *gin.Context) (*response.Response, error) {
	d.log.Infof("根据ID查看字典")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	dict := po.Dict{ID: id}
	if err = d.dictService.ISelectOne(c, &dict); err != nil {
		d.log.Errorf("根据ID查看字典失败： error: %s", err)
		return nil, err
	}
	return response.Success(dict), nil
}

func (d *Dict) List(c *gin.Context) (*response.Response, error) {
	p := vo.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	d.log.Infof("分页查询字典列表")

	var dict *po.Dict
	name := c.Query("name")

	if name ==  "" {
		dict = &po.Dict{}
	} else {
		dict = &po.Dict{Name: name}
	}

	if err := d.dictService.ISelectAll(c, &p, dict); err != nil {
		return nil, err
	}

	return response.Success(&p), nil
}

func (d *Dict) Post(c *gin.Context) (*response.Response, error) {
	d.log.Infof("新建字典")

	dict := &po.Dict{}
	if err := c.ShouldBindJSON(&dict); err != nil {
		d.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := d.dictService.ICreateOne(c, dict); err != nil {
		d.log.Errorf("新建字典失败：error: %s", err)
		return nil, err
	}

	d.log.Infof("新建字典成功")
	return response.Success(dict), nil
}

func (d *Dict) Delete(c *gin.Context) (*response.Response, error) {
	d.log.Infof("删除字典")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := d.dictService.IDeleteOne(c, &po.Dict{ID: id}); err != nil {
		return nil, err
	}
	return response.Success("删除成功. "), nil
}

func (d *Dict) Patch(c *gin.Context) (*response.Response, error) {
	panic("implement me")
}

func (d *Dict) Put(c *gin.Context) (*response.Response, error) {
	d.log.Infof("修改字典")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var dict po.Dict
	if err := c.ShouldBindJSON(&dict); err != nil {
		d.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	dict.ID = id
	if err := d.dictService.IUpdateOne(c, &dict); err != nil {
		return nil, err
	}

	return response.Success(&dict), nil
}