package service

import (
	"blog/pkg/model/common"
)

type PermissionService interface {
	common.BaseService
}


type PermissionServiceImpl struct {
	common.BaseServiceImpl
}

func NewPermissionService() PermissionService {
	return &PermissionServiceImpl{}
}

