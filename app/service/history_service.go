package service

import (
	"blog/pkg/global"
	"blog/pkg/model/po"
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
func (s *HistoryService) CreateOne(history *po.History) error {
	if err := checkHeadId(history); err != nil {
		return err
	}
	if err := checkRepositoryId(history); err != nil {
		return err
	}

	// 查询是否存在headId并且已经发布的版本，如果存在则设置prev为已发布的版本?
	var head *po.Head
	if err := global.DB.Model(&po.Head{}).Where("id=?", history.HeadID).First(&head).Error; err != nil {
		return err
	}
	if head.Status == 3 {
		history.PrevRepositoryID = head.RepositoryID
	}

	return global.DB.Model(&po.History{}).Create(history).Error
}

// SelectOne select one history from database filter by param history
// sql: select * from history where head_id=? and repository_id=?
func (s *HistoryService) SelectOne(history *po.History) error {
	if err := checkHeadId(history); err != nil {
		return err
	}
	if err := checkRepositoryId(history); err != nil {
		return err
	}
	return global.DB.Model(&po.History{}).Where("head_id=? and repository_id=?", history.HeadID, history.RepositoryID).First(&history).Error
}

// SelectList select all histories from database filter by history
// sql: select * from history where head_id=?
// todo: 考虑是否需要分页或者滚动获取
func (s *HistoryService) SelectList(history *po.History, histories *[]*po.History) error {
	if err := checkHeadId(history); err != nil {
		return err
	}
	return global.DB.Model(&po.History{}).Where("head_id=?", history.HeadID).Find(&histories).Error
}

// UpdateOne update one history (exp: staged、commit、publish)
func (s *HistoryService) UpdateOne(history *po.History) error {
	if err := checkHeadId(history); err != nil {
		return err
	}
	if err := checkRepositoryId(history); err != nil {
		return err
	}
	return global.DB.Model(&po.History{}).Where("head_id=? and repository_id=?", history.HeadID, history.RepositoryID).Updates(history).Error
}

func checkRepositoryId(history *po.History) error {
	if history.RepositoryID == 0 {
		return errors.New("history's repository_id is not nil. ")
	}
	return nil
}
func checkHeadId(history *po.History) error {
	if history.HeadID == 0 {
		return errors.New("history's head_id is not nil. ")
	}
	return nil
}
