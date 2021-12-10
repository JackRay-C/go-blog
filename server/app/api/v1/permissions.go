package v1

import (
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
)

type Permission struct {
	log logger.Logger
	permissionService *service.PermissionService
}

func NewPermission() *Permission {
	return &Permission{
		log: global.Logger,
		permissionService: service.NewPermissionService(),
	}
}

func (p *Permission) Get(c *gin.Context) (*response.Response, error) {
	return nil,nil
}

func (p *Permission) List(c *gin.Context) (*response.Response, error) {
	return nil,nil
}

func (p *Permission) Post(c *gin.Context) (*response.Response, error) {
	return nil,nil
}

func (p *Permission) Delete(c *gin.Context) (*response.Response, error) {
	return nil,nil
}

func (p *Permission) Patch(c *gin.Context) (*response.Response, error) {
	return nil,nil
}

func (p *Permission) Put(c *gin.Context) (*response.Response, error) {
	return nil,nil
}



