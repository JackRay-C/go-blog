package service

import (

	"blog/app/pager"
	"blog/pkg/global"
	"blog/pkg/model/po"

	"errors"
	"gorm.io/gorm"
	"log"
)

type HeadService struct {

}

// NewHeadService constructor function
func NewHeadService() *HeadService {
	return &HeadService{}
}

// CreateOne create one head to database
func (s *HeadService) CreateOne(head *po.Head) error {
	return global.DB.Model(&po.Head{}).Create(head).Error
}


// SelectOne
func (s *HeadService) SelectOne(head *po.Head) error {
	db := global.DB.Model(&po.Head{}).Where("id=?", head.ID)
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

func (s *HeadService) SelectList(page *pager.Pager, head *po.Head) error {
	var heads []*po.Head
	offset := (page.PageNo - 1) * page.PageSize
	limit := page.PageSize
	var count int64

	db := global.DB.Model(&po.Head{})
	if head.Status != 0 {
		db.Where("status = ?", head.Status)
	}
	if head.Visibility != 0 {
		db.Where("visibility=?", head.Visibility)
	}
	if head.SubjectID != 0 {
		db.Where("subject_id=?", head.SubjectID)
	}
	if head.UserID != 0 {
		db.Where("user_id=?", head.UserID)
	}
	if err := db.Count(&count).Error; err != nil {
		return err
	}
	if err := db.Order("created_at desc").Offset(offset).Limit(limit).Find(&heads).Error;err != nil {
		return err
	}

	page.TotalRows = count
	log.Println(&heads)
	page.MustList(&heads)

	log.Println(page)
	return nil
}

func (s *HeadService) UpdateOne(head *po.Head) error {
	var nHead *po.Head
	err := global.DB.Model(&po.Head{}).Where("id=?", head.ID).Find(&nHead).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该记录不存在。")
	}
	if err != nil {
		return err
	}
	if nHead.UserID != head.UserID  {
		return errors.New("没有权限更改. ")
	}

	if err := global.DB.Model(&po.Head{}).Where("id=?", head.ID).Updates(head).Error; err != nil {
		return err
	}
	return nil
}

func (s *HeadService) DeleteOne(head *po.Head) error {
	var nHead *po.Head
	err := global.DB.Model(&po.Head{}).Where("id=? and user_id=?", head.ID, head.UserID).Find(&nHead).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该记录不存在。")
	}
	if err != nil {
		return err
	}
	if err := global.DB.Model(&po.Head{}).Where("id=?", head.ID).Delete(&head).Error; err != nil {
		return err
	}
	return nil
}