package service

import (

	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
)

type RepositoryServiceImpl struct {

}

type RepositoryService interface {
	ISelectAll(page *vo.Pager, repository *po.Repository) error
	ICreateOne(r *po.Repository) error
	ISelectOne()
	IDeleteOne()
	IUpdateOne()
}

func NewRepositoryService() RepositoryService {
	return &RepositoryServiceImpl{}
}

func (s *RepositoryServiceImpl) ICreateOne(r *po.Repository) error {
	return global.DB.Model(&po.Repository{}).Create(r).Error
}

func (s *RepositoryServiceImpl) ISelectAll(page *vo.Pager, repository *po.Repository) error {
	// 根据headid获取所有的repository
	var repositories []*po.Repository
	offset := (page.PageNo - 1) * page.PageSize
	limit := page.PageSize

	db := global.DB.Model(&po.Repository{})

	if repository.UserId!=0 {
		db.Where("user_id=?", repository.UserId)
	}

	if err := db.Count(&page.TotalRows).Error; err != nil {
		return err
	}
	if err := db.Order("created_at desc").Offset(offset).Limit(limit).Find(&repositories).Error;err != nil {
		return err
	}

	page.MustList(&repositories)

	return nil
}

func (s *RepositoryServiceImpl) ISelectOne() {
	panic("implement me")
}

func (s *RepositoryServiceImpl) IDeleteOne() {
	panic("implement me")
}

func (s *RepositoryServiceImpl) IUpdateOne() {
	panic("implement me")
}