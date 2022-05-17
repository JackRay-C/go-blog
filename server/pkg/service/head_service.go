package service

import (
	"blog/pkg/model/common"
	"blog/pkg/service/impl"
)

type HeadService interface {
	common.BaseService
}

// NewHeadService constructor function
func NewHeadService() HeadService {
	return &impl.HeadServiceImpl{}
}