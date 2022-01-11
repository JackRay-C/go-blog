package service

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/core/global"
	"errors"
	"gorm.io/gorm"
)

type MenuService struct {
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (m *MenuService) SelectOne(menu *domain.Menu) error {
	err := global.DB.Model(&domain.Menu{}).Where("id=?", menu.ID).First(&menu).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该菜单不存在. ")
	}

	return err
}

func (m *MenuService) SelectAll(page *pager.Pager, menu *domain.Menu) error {
	menus := make([]*domain.Menu, 0)

	db := global.DB.Model(&domain.Menu{})

	if err := db.Count(&page.TotalRows).Error; err != nil {
		return err
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &menus

	if err := db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&menus).Error; err != nil {
		return err
	}
	return nil
}

func (m *MenuService) DeleteOne(menu *domain.Menu) error {
	var newMenu *domain.Menu
	err := global.DB.Model(&domain.Menu{}).Where("id=?", menu.ID).First(&newMenu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该用户不存在. ")
	}

	if err != nil {
		return err
	}

	return global.DB.Model(&domain.Menu{}).Where("id=?", menu.ID).Delete(menu).Error
}

func (m *MenuService) CreateOne(menu *domain.Menu) error {
	// 1、根据名字和路由查询是否存在
	var newMenu *domain.Menu
	err := global.DB.Model(&domain.Menu{}).Where("name=? and path=?", menu.Name, menu.Path).First(&newMenu).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 2、不存在则创建
		if err := global.DB.Model(&domain.Menu{}).Create(&menu).Error; err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return errors.New("该菜单已经存在. ")
}

// 更新
func (m *MenuService) SaveOne(menu *domain.Menu) error {
	// 根据ID查询是否存在
	var newMenu *domain.Menu

	err := global.DB.Model(&domain.Menu{}).Where("id=?", menu.ID).First(&newMenu).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该菜单不存在. ")
	}

	return global.DB.Model(&domain.Menu{}).Save(&menu).Error
}
