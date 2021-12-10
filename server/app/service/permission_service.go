package service

import (
	"blog/app/domain"
	"blog/core/global"
	"blog/core/logger"
)

type PermissionService struct {
	log logger.Logger
}

func NewPermissionService() *PermissionService {
	return &PermissionService{
		log: global.Logger,
	}
}

func (s PermissionService) SelectOne(permission *domain.Permissions) error {
	return nil
}

