package service

import (
	"blog/app/domain"
	"blog/core/global"
	"blog/core/logger"
	"errors"
	"gorm.io/gorm"
)

type RolesMenusService struct {
	log logger.Logger
}

func NewRolesMenusService() *RolesMenusService {
	return &RolesMenusService{
		log: global.Logger,
	}
}

func (rm *RolesMenusService) SelectMenusByRoles(menus *[]*domain.Menu, roles... *domain.Role) error {
	for _, role := range roles {
		var ms []*domain.Menu
		if err := global.DB.Table("menus").Joins("left join roles_menus as rm on menus.id=rm.menu_id").Where("rm.role_id = ?", role.ID).Find(&ms).Error;err!=nil {
			return err
		}
		*menus = append(*menus, ms...)
	}
	return nil
}

/*
	查询角色菜单
*/
func (s *RolesMenusService) SelectRoleMenus(role *domain.Role, menus *[]*domain.Menu) error {
	var r *domain.Role
	err := global.DB.Model(&domain.Role{}).Where("id=?", role.ID).First(&r).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在. ")
	}

	if err != nil {
		return err
	}

	err = global.DB.Table("menus").Joins("left join roles_menus on menus.id=roles_menus.menu_id ").Joins("left join roles on roles_menus.role_id=roles.id").Where("roles.id=?", r.ID).Find(&menus).Error

	if err != nil {
		return err
	}
	return nil
}


/*
	更新角色菜单
*/
func (s *RolesMenusService) UpdateRoleMenus(role *domain.Role, menus []*domain.Menu) error {
	var r *domain.Role
	if err := global.DB.Model(&domain.Role{}).Where("id=?", role.ID).First(&r).Error; err != nil || err == gorm.ErrRecordNotFound {
		return err
	}

	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 1、删除角色的原来的菜单
		tx = tx.Model(&domain.RoleMenu{}).Where("role_id=?", role.ID)
		if err := tx.Delete(role).Error; err != nil {
			return err
		}

		// 2、添加新的角色-菜单的关系
		var roleMenus []*domain.RoleMenu
		for _, menu := range menus {
			roleMenus = append(roleMenus, &domain.RoleMenu{MenuId: menu.ID, RoleId: role.ID})
		}
		if err := tx.Create(roleMenus).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
