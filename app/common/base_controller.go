package common

import (
	"blog/app/response"
	"github.com/gin-gonic/gin"
)

type BaseController interface {
	Get(c *gin.Context) (*response.Response, error)    //根据ID获取
	List(c *gin.Context) (*response.Response, error)   //分页获取所有
	Post(c *gin.Context) (*response.Response, error)   // 创建
	Delete(c *gin.Context) (*response.Response, error) // 根据id删除
	Patch(c *gin.Context) (*response.Response,error)  // 根据id更新
	Put(c *gin.Context) (*response.Response, error)    // 根据id全量更新
}
