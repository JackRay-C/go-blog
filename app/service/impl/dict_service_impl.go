package impl

import (
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"github.com/gin-gonic/gin"
)

type DictServiceImpl struct {
	
}

func (d *DictServiceImpl) IUpdateOne(c *gin.Context, dict *po.Dict) error {
	panic("implement me")
}

func (d *DictServiceImpl) ICreateOne(c *gin.Context, dict *po.Dict) error {
	panic("implement me")
}

func (d *DictServiceImpl) IDeleteOne(c *gin.Context, dict *po.Dict) error {
	panic("implement me")
}

func (d *DictServiceImpl) ISelectOne(c *gin.Context, dict *po.Dict) error {
	panic("implement me")
}

func (d *DictServiceImpl) ISelectAll(c *gin.Context, pager *vo.Pager, dict *po.Dict) error {
	panic("implement me")
}

