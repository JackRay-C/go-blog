package service

import (
	"blog/app/model/po"
	"blog/core/global"
	"blog/core/logger"
	"errors"
	"gorm.io/gorm"
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
func (rps RolesPermissionService) SelectPermissionByRoleId(permissions *[]*po.Permissions, roleId... int) error {
	for _, id := range roleId {
		var ps []*po.Permissions
		if err := global.DB.Table("permissions").Joins("left join roles_permissions as ur on permissions.id=ur.permission_id").Where("ur.role_id = ?", id).Find(&ps).Error;err!=nil {
			rps.log.Errorf("获取用户权限列表失败：%s", err)
			return err
		}
		*permissions = append(*permissions, ps...)
	}

	return nil
}

func (rps *RolesPermissionService) SelectRolePermissions(role *po.Role, permissions *[]*po.Permissions) error {
	// 1、查询角色是否存在
	err := global.DB.Model(&po.Role{}).Where("id=?", role.ID).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在. ")
	}
	if err != nil {
		return err
	}

	// 2、根据角色ID查询权限列表
	if err := global.DB.Table("permissions").Joins("left join roles_permissions as ur on permissions.id=ur.permission_id").Where("ur.role_id=?", role.ID).Find(&permissions).Error; err != nil {
		return err
	}
	return nil
}

func (rps *RolesPermissionService) UpdateRolePermissions(role *po.Role, permissions []*po.Permissions) error  {
	// 1、查询角色是否存在
	err := global.DB.Model(&po.Role{}).Where("id=?", role.ID).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在. ")
	}

	if err != nil {
		return err
	}

	// 2、更新角色权限列表
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 1、删除角色的原来的菜单
		tx = tx.Model(&po.RolesPermissions{}).Where("role_id=?", role.ID)
		if err := tx.Delete(role).Error; err != nil {
			return err
		}

		// 2、添加新的角色-菜单的关系
		var rolePermissions []*po.RolesPermissions
		for _, permission := range permissions {
			rolePermissions = append(rolePermissions, &po.RolesPermissions{PermissionId: permission.ID, RoleId: role.ID})
		}
		if err := tx.Create(rolePermissions).Error; err != nil {
			return err
		}
		return nil
	})
}

func (rps *RolesPermissionService) SelectPermissionByRoles(permissions *[]*po.Permissions, roles ...*po.Role) error {
	for _, role := range roles {
		var ms []*po.Permissions
		if err := global.DB.Table("permissions").Joins("left join roles_permissions as rm on permissions.id=rm.permission_id").Where("rm.role_id = ?", role.ID).Find(&ms).Error;err!=nil {
			return err
		}
		*permissions = append(*permissions, ms...)
	}
	return nil
}


