package service

import (
	"blog/app/model/po"
	"blog/app/model/vo"
	"blog/app/service/impl"
	"github.com/gin-gonic/gin"
)

type DictService interface {
	ISelectOne(c *gin.Context, dict *po.Dict) error
	ISelectAll(c *gin.Context, pager *vo.Pager, dict *po.Dict) error
	ICreateOne(c *gin.Context, dict *po.Dict) error
	IDeleteOne(c *gin.Context, dict *po.Dict) error
	IUpdateOne(c *gin.Context, dict *po.Dict) error
}


func NewDictService() DictService {
	return &impl.DictServiceImpl{}
}



