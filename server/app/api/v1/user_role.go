package v1

import (
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	log           logger.Logger
	roleService *service.RoleService
}

func NewUserRole() *UserRole {
	return &UserRole{
		log: global.Logger,
		roleService: service.NewRoleService(),
	}
}

func (ur *UserRole) Get(c *gin.Context) (*response.Response, error)  {

	return response.Success("success"), nil
}

func (ur *UserRole) List(c *gin.Context) (*response.Response, error)  {

	return response.Success("success"), nil
}