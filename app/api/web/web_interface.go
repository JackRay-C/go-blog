package web

import (
	"blog/app/response"
	"github.com/gin-gonic/gin"
)

type Web interface {
	Get(c *gin.Context) (*response.Response, error)
	List(c *gin.Context) (*response.Response, error)

}
