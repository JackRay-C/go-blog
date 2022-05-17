package impl

import (
	"blog/pkg/model/po"
	"github.com/gin-gonic/gin"
)

type RolePermissionServiceImpl struct {}

func (r *RolePermissionServiceImpl) ISelectPermissionByRoles(c *gin.Context, permissions *[]*po.Permissions, roles ...*po.Role) error {
	panic("implement me")
}

func (r *RolePermissionServiceImpl) IUpdateRolePermissions(c *gin.Context, role *po.Role, permissions []*po.Permissions) error {
	panic("implement me")
}

