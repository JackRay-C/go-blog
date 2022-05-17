package service

import (

	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"errors"
	"gorm.io/gorm"
	"time"
)

type RoleService struct {
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (s *RoleService) SelectOne(role *po.Role) error {
	if err := global.DB.Model(&po.Role{}).Where("id=?", role.ID).First(&role).Error; err != nil || err == gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func (s *RoleService) SelectAll(page *vo.Pager) error {
	var roles []po.Role

	db := global.DB.Model(&po.Role{})

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

func (s *RoleService) DeleteOne(role *po.Role) error {
	db := global.DB.Model(&po.Role{}).Where("id=?", role.ID)
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


func (s *RoleService) CreateOne(role *po.Role) error {
	db := global.DB.Model(&po.Role{})
	// 查询是否存在该角色
	var r *po.Role
	if err := db.Where("name=?", role.Name).First(&r).Error; err != nil {
		return err
	}
	// 判断如果有
	global.Log.Println(r.String())
	if r != nil {
		return errors.New("该角色已存在。 ")
	}
	if err := db.Create(&po.Role{
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   time.Now(),
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *RoleService) UpdateOne(role *po.Role) error {
	db := global.DB.Model(&po.Role{})
	var r *po.Role
	if err := db.Where("id=?", role.ID).First(&r).Error; err != nil || err == gorm.ErrRecordNotFound{
		return err
	}

	role.UpdatedAt = time.Now()
	if err := db.Where("id=?", role.ID).Updates(role).Error; err != nil {
		return err
	}
	return nil
}

