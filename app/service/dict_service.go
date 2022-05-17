package service

import (

	"blog/app/service/impl"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
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



