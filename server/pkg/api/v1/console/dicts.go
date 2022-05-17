package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Dict struct {
	log         logger.Logger
	dictService service.DictService
}

func NewDict() *Dict {
	return &Dict{dictService: service.NewDictService(), log: global.Log}
}

func (d *Dict) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}
	d.log.Infof("Select one dict by id: %d", id)

	dict := po.Dict{ID: id}
	if err = d.dictService.ISelectOne(c, &dict); err != nil {
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

	d.log.Infof("分页查询字典列表")

	var dict *po.Dict
	name := c.Query("name")

	if name == "" {
		dict = &po.Dict{}
	} else {
		dict = &po.Dict{Name: name}
	}

	if err := d.dictService.ISelectList(c, &p, dict); err != nil {
		d.log.Errorf("list dict failed service error: %s", err)
		return nil, err
	}

	d.log.Infof("list dict success. ")
	return vo.Success(&p), nil
}

func (d *Dict) Post(c *gin.Context) (*vo.Response, error) {
	d.log.Infof("new dict. ")

	dict := &po.Dict{}
	if err := c.ShouldBindJSON(&dict); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := d.dictService.ICreateOne(c, dict); err != nil {
		d.log.Errorf("新建字典失败：error: %s", err)
		return nil, err
	}

	d.log.Infof("新建字典成功")
	return vo.Success(dict), nil
}

func (d *Dict) Delete(c *gin.Context) (*vo.Response, error) {
	d.log.Infof("删除字典")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := d.dictService.IDeleteOne(c, &po.Dict{ID: id}); err != nil {
		return nil, err
	}
	return vo.Success("删除成功. "), nil
}

func (d *Dict) Patch(c *gin.Context) (*vo.Response, error) {
	panic("implement me")
}

func (d *Dict) Put(c *gin.Context) (*vo.Response, error) {
	d.log.Infof("修改字典")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	var dict po.Dict
	if err := c.ShouldBindJSON(&dict); err != nil {
		d.log.Errorf("参数绑定错误：%s", err)
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	dict.ID = id
	if err := d.dictService.IUpdateOne(c, &dict, dict); err != nil {
		return nil, err
	}

	return vo.Success(&dict), nil
}
