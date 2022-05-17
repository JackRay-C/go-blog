package service

import (
	"blog/pkg/model/common"
	"blog/pkg/service/impl"
)

type SubjectService interface {
	common.BaseService
}

func NewSubjectService() SubjectService {
	return &impl.SubjectServiceImpl{}
}