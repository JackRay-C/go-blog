package service

import (
	"blog/app/domain"
	"blog/core/global"
	"errors"
)

type HistoryService struct {

}

// NewHistoryService constructor function
func NewHistoryService() *HistoryService {
	return &HistoryService{}
}

// CreateOne create one history to databases
// sql: insert into history(head_id, repository_id, prev_repository_id, staged_at, commited_at, published_at) values(?,?,?,?,?,?)
func (s *HistoryService) CreateOne(history *domain.History) error {
	if err := checkHeadId(history); err != nil {
		return err
	}
	if err := checkRepositoryId(history); err != nil {
		return err
	}
	return global.DB.Model(&domain.History{}).Create(history).Error
}

// SelectOne select one history from database filter by param history
// sql: select * from history where head_id=? and repository_id=?
func (s *HistoryService) SelectOne(history *domain.History) error  {
	if err := checkHeadId(history); err != nil {
		return err
	}
	if err := checkRepositoryId(history); err != nil {
		return err
	}
	return global.DB.Model(&domain.History{}).Where("head_id=? and repository_id=?", history.HeadID, history.RepositoryID).First(&history).Error
}

// SelectList select all histories from database filter by history
// sql: select * from history where head_id=?
// todo: 考虑是否需要分页或者滚动获取
func (s *HistoryService) SelectList(history *domain.History, histories *[]*domain.History) error {
	if err := checkHeadId(history); err != nil {
		return err
	}
	return global.DB.Model(&domain.History{}).Where("head_id=?", history.HeadID).Find(&histories).Error
}

func (s *HistoryService)   {

}




func checkRepositoryId(history *domain.History) error {
	if history.RepositoryID == 0 {
		return errors.New("history's repository_id is not nil. ")
	}
	return nil
}
func checkHeadId(history *domain.History) error  {
	if history.HeadID == 0 {
		return errors.New("history's head_id is not nil. ")
	}
	return nil
}