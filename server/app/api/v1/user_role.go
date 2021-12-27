package v1

import (
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
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
