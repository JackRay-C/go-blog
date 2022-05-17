package service

import (
	"blog/pkg/model/common"
	"blog/pkg/service/impl"
)

type HistoryService interface {
	common.BaseService
}

// NewHistoryService constructor function
func NewHistoryService() HistoryService {
	return &impl.HistoryServiceImpl{}
}
