package service

import (
	"blog/app/domain"
	"blog/core/global"
)

type RepositoryService struct {

}

func NewRepositoryService() *RepositoryService {
	return &RepositoryService{}
}



func (s *RepositoryService) CreateOne(r *domain.Repository) error {
	return global.DB.Model(&domain.Repository{}).Create(r).Error
}
