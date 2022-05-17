package service

import (
	"blog/pkg/model/common"
	"blog/pkg/service/impl"
)

type RoleService interface {
	common.BaseService
}


func NewRoleService() RoleService {
	return &impl.RoleServiceImpl{}
}


