package service

import (
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/pager"
	"blog/core/global"
	"errors"
	"gorm.io/gorm"
)

type HeadService struct {

}

// NewHeadService constructor function
func NewHeadService() *HeadService {
	return &HeadService{}
}

// CreateOne create one head to database
func (s *HeadService) CreateOne(head *domain.Head) error {
	return global.DB.Model(&domain.Head{}).Create(head).Error
}

func (s *HeadService) SelectOne(head *domain.Head) error {
	db := global.DB.Model(&domain.Head{}).Where("id=?", head.ID)
	if head.UserID != 0 {
		db.Where("user_id=?", head.UserID)
	}
	if head.RepositoryID!= 0 {
		db.Where("repository_id=?", head.RepositoryID)
	}
	if head.Status != 0 {
		db.Where("status=?", head.Status)
	}
	if head.Visibility != 0 {
		db.Where("visibility=?", head.Visibility)
	}
	if head.SubjectID != 0 {
		db.Where("subject_id=?", head.SubjectID)
	}
	err := db.First(&head).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("不存在该博客. ")
	}
	return err
}

func (s *HeadService) SelectList(page *pager.Pager, query dto.Query) error {
	return nil
}