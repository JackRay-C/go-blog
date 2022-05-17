package service

import (
	"blog/pkg/model/common"
	"blog/pkg/service/impl"
)

type DictService interface {
	common.BaseService
}


func NewDictService() DictService {
	return &impl.DictServiceImpl{}
}



