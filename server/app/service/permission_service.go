package service

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/core/global"
	"blog/core/logger"
	"errors"
	"gorm.io/gorm"
)

type PermissionService struct {
	log logger.Logger
}

func NewPermissionService() *PermissionService {
	return &PermissionService{
		log: global.Logger,
	}
}

func (s *PermissionService) SelectOne(permission *domain.Permissions) error {
	return global.DB.Model(&domain.Permissions{}).Where("id=?", permission.ID).First(&permission).Error
}

func (s *PermissionService) CreateOne(permission *domain.Permissions) error {
	// 查询是否存在相同的权限
	err := global.DB.Model(&domain.Permissions{}).Where("object_type=? and action_type=?", permission.ObjectType, permission.ActionType).First(&permission).Error

	if err == gorm.ErrRecordNotFound {
		return global.DB.Model(&domain.Permissions{}).Create(permission).Error
	}

	if err != nil {
		return err
	}
	return errors.New("该权限已经存在. ")
}

func (s *PermissionService) SelectAll(p *pager.Pager) error {
	offset := (p.PageNo - 1) * p.PageSize
	limit := p.PageSize
	db := global.DB.Model(&domain.Permissions{})

	var permissions []*domain.Permissions
	var count int64

	if err := db.Count(&count).Error; err != nil {
		return err
	}

	if err := db.Offset(offset).Limit(limit).Find(&permissions).Error; err != nil {
		return err
	}

	p.TotalRows = count
	if count == 0 {
		p.PageCount = 0
		p.List = make([]string, 0)
	} else {
		p.PageCount = int((count + int64(p.PageSize) - 1) / int64(p.PageSize))
		p.List = &permissions
	}

	return nil
}

func (s *PermissionService) DeleteOne(p *domain.Permissions) error {
	return global.DB.Model(&domain.Permissions{}).Where("id=?", p.ID).Delete(p).Error
}

func (s *PermissionService) UpdateOne(permission *domain.Permissions) error {
	// 1、查询是否存在
	var newPermission *domain.Permissions
	if err := global.DB.Model(&domain.Permissions{}).Where("id=?", permission.ID).First(&newPermission).Error; err != nil {
		return err
	}

	// 2、判断更新的权限是否存在
	newPermission = nil
	err := global.DB.Model(&domain.Permissions{}).Where("object_type=? and action_type=? and id!=?", permission.ObjectType, permission.ActionType, permission.ID).First(&newPermission).Error

	if err == gorm.ErrRecordNotFound {
		// 3、如果不存在的话，更新新权限
		if err := global.DB.Model(&domain.Permissions{}).Where("id=?", permission.ID).Updates(permission).Error; err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	return errors.New("该权限类型已存在. ")
}
