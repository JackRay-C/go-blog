package service

import (
	"blog/pkg/model/common"
	"blog/pkg/service/impl"
)

type TagService interface {
	common.BaseService
}

func NewTagService() TagService {
	return &impl.TagServiceImpl{}
}

