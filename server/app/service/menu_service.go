package service

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/app/response"
	"gorm.io/gorm"
	"time"
)

type MenuService struct {

}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (m *MenuService) SelectOne(menu *domain.Menu) error {
	if err := menu.Select(); err == gorm.ErrRecordNotFound {
		return response.RecordNotFound.SetMsg("该模块不存在！")
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return nil
}

func (m *MenuService) SelectAll(page *pager.Pager, menu *domain.Menu) error {
	var menus []domain.Menu

	if err := menu.Count(&page.TotalRows); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &menus

	if err := menu.List(&menus, (page.PageNo-1)*page.PageSize, page.PageSize); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	return nil
}

func (m *MenuService) DeleteOne(menu *domain.Menu) error {
	if err := menu.Select(); err == gorm.ErrRecordNotFound {
		return response.RecordNotFound.SetMsg("该模块不存在！")
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	if err := menu.Delete(); err != nil {
		return response.DatabaseDeleteError.SetMsg("%s", err)
	}
	return nil
}

func (m *MenuService) DeleteAll(ids []int) error {
	menu := domain.Menu{}
	if err := menu.DeleteIds(ids); err != nil {
		return response.DatabaseDeleteError.SetMsg("%s", err)
	}
	return nil
}

func (m *MenuService) CreateOne(menu *domain.Menu) error {
	m1 := &domain.Menu{Path: menu.Path}
	if err := m1.Select(); err == gorm.ErrRecordNotFound {
		if err := menu.Insert(); err != nil {
			return response.DatabaseInsertError.SetMsg("%s", err)
		}
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return response.RecoreExisted.SetMsg("该菜单已存在！")
}

func (m *MenuService) UpdateOne(menu *domain.Menu) error {
	route1 := &domain.Menu{ID: menu.ID}
	if err := route1.Select(); err == gorm.ErrRecordNotFound {
		return response.RecordNotFound.SetMsg("该记录不存在. ")
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}


	menu.UpdatedAt = time.Now()
	if err := menu.Update(); err != nil {
		return response.DatabaseUpdateError.SetMsg("%s", err)
	}
	if err := menu.Select(); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return nil
}

func (m *MenuService) SaveOne(d *domain.Menu) error {
	return nil
}


