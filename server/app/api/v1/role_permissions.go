package v1

import (
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
)

type RolePermission struct {
	log           logger.Logger
	roleService *service.RoleService
}

func NewRolePermission() *RolePermission {
	return &RolePermission{
		log: global.Logger,
		roleService: service.NewRoleService(),
	}
}

func (r *RolePermission) Get(c *gin.Context) (*response.Response, error)  {

	return response.Success("success"), nil
}

func (r *RolePermission) Put(c *gin.Context) (*response.Response, error)  {

	return response.Success("success"), nil
}
