package service

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/core/global"
	"errors"
	"gorm.io/gorm"
	"time"
)

type RoleService struct {
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (s *RoleService) SelectOne(role *domain.Role) error {
	if err := global.DB.Model(&domain.Role{}).Where("id=?", role.ID).First(&role).Error; err != nil || err == gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (s *RoleService) SelectAll(page *pager.Pager) error {
	var roles []domain.Role

	db := global.DB.Model(&domain.Role{})

	if err := db.Count(&page.TotalRows).Error; err != nil {
		return err
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &roles

	if err := db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&roles).Error; err != nil {
		return err
	}

	return nil
}

func (s *RoleService) DeleteOne(role *domain.Role) error {
	db := global.DB.Model(&domain.Role{}).Where("id=?", role.ID)
	// 查询是否存在
	if err := db.First(&role).Error; err != nil || err == gorm.ErrRecordNotFound {
		return err
	}
	// 删除记录
	if err := db.Delete(role).Error; err != nil {
		return err
	}
	return nil
}


func (s *RoleService) CreateOne(role *domain.Role) error {
	db := global.DB.Model(&domain.Role{})
	// 查询是否存在该角色
	var r *domain.Role
	if err := db.Where("name=?", role.Name).First(&r).Error; err != nil {
		return err
	}
	// 判断如果有
	global.Logger.Println(r.String())
	if r != nil {
		return errors.New("该角色已存在。 ")
	}
	if err := db.Create(&domain.Role{
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *RoleService) UpdateOne(role *domain.Role) error {
	db := global.DB.Model(&domain.Role{})
	var r *domain.Role
	if err := db.Where("id=?", role.ID).First(&r).Error; err != nil || err == gorm.ErrRecordNotFound{
		return err
	}

	role.UpdatedAt = time.Now()
	if err := db.Where("id=?", role.ID).Updates(role).Error; err != nil {
		return err
	}
	return nil
}


/*
	查询角色菜单
*/
func (s *RoleService) SelectMenus(menus *[]*domain.Menu, role *domain.Role) error {
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
func (s *RoleService) UpdateMenus(role *domain.Role, menus []*domain.Menu) error {
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
