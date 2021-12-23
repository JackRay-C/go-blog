package service

import (
	"blog/app/domain"
	"blog/core/global"
	"blog/core/logger"
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