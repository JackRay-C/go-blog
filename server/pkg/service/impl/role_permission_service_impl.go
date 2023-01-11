package impl

import (
	"blog/pkg/global"
	"blog/pkg/model/po"
	"github.com/gin-gonic/gin"
)

type RolePermissionServiceImpl struct {}

func (r *RolePermissionServiceImpl) ISelectPermissionByRoles(c *gin.Context, permissions *[]*po.Permissions, roles ...*po.Role) error {
	for _, role := range roles {
		var ps []*po.Permissions
		if err := global.DB.Table("permissions").Joins("left join roles_permissions as ur on permissions.id=ur.permission_id").Where("ur.role_id = ?", role.ID).Find(&ps).Error;err!=nil {
			return err
		}
		*permissions = append(*permissions, ps...)
	}

	return nil
}

func (r *RolePermissionServiceImpl) IUpdateRolePermissions(c *gin.Context, role *po.Role, permissions []*po.Permissions) error {
	panic("implement me")
}

