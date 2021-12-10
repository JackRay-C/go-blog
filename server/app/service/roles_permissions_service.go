package service

import (
	"blog/app/domain"
	"blog/core/global"
	"blog/core/logger"
)

type RolesPermissionService struct {
	log logger.Logger
}

func NewRolesPermissionService() *RolesPermissionService {
	return &RolesPermissionService{
		log: global.Logger,
	}
}

// 根据用户id获取所有角色
func (rps RolesPermissionService) SelectPermissionByRoleId(permissions *[]*domain.Permissions, roleId... int) error {
	rps.log.Infof("根据角色ID【%d】获取所有权限列表", roleId)
	for _, id := range roleId {
		var ps []*domain.Permissions
		if err := global.DB.Table("permissions").Joins("left join roles_permissions as ur on permissions.id=ur.permission_id").Where("ur.role_id = ?", id).Find(&ps).Error;err!=nil {
			rps.log.Errorf("获取用户权限列表失败：%s", err)
			return err
		}
		*permissions = append(*permissions, ps...)
	}

	return nil
}




