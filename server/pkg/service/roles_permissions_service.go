package service

import (
	"blog/pkg/model/po"
	"blog/pkg/service/impl"
	"github.com/gin-gonic/gin"
)

type RolePermissionService interface {
	ISelectPermissionByRoles(c *gin.Context, permissions *[]*po.Permissions, roles ...*po.Role) error
	IUpdateRolePermissions(c *gin.Context, role *po.Role, permissions []*po.Permissions) error
}

func NewRolesPermissionService() RolePermissionService {
	return &impl.RolePermissionServiceImpl{}
}


