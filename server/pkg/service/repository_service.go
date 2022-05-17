package service

import (
	"blog/pkg/model/common"
	"blog/pkg/service/impl"
)


type RepositoryService interface {
	common.BaseService
}

func NewRepositoryService() RepositoryService {
	return &impl.RepositoryServiceImpl{}
}

