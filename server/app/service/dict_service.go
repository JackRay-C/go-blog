package service

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/app/response"
	"gorm.io/gorm"
)

type DictService struct {
	
}


func NewDictService() *DictService {
	return &DictService{}
}

func (d *DictService) SelectOne(dict *domain.Dict) error {
	if err := dict.Select(); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return nil
}

func (d *DictService) SelectAll(page *pager.Pager, dict *domain.Dict) error {
	var dicts []*domain.Dict

	if err := dict.Count(&page.TotalRows); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &dicts

	if err := dict.List(&dicts, (page.PageNo-1)*page.PageSize, page.PageSize); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	return nil
}

func (d *DictService) DeleteOne(dict *domain.Dict) error {
	if err := dict.Select(); err == gorm.ErrRecordNotFound {
		return response.RecordNotFound.SetMsg("该模块不存在！")
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	if err := dict.Delete(); err != nil {
		return response.DatabaseDeleteError.SetMsg("%s", err)
	}
	return nil
}

func (d *DictService) DeleteAll(ids []int) error {
	panic("implement me")
}

func (d *DictService) CreateOne(dict *domain.Dict) error {
	if err := dict.Select(); err == gorm.ErrRecordNotFound {
		if err := dict.Insert(); err != nil {
			return response.DatabaseInsertError.SetMsg("%s", err)
		} else {
			return nil
		}
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return response.RecoreExisted.SetMsg("该字典已存在！")
}

func (d *DictService) UpdateOne(dict *domain.Dict) error {
	dict1 := &domain.Dict{ID: dict.ID}
	if err := dict1.Select(); err == gorm.ErrRecordNotFound {
		return response.RecordNotFound.SetMsg("该记录不存在. ")
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	if err := dict.Update(); err != nil {
		return response.DatabaseUpdateError.SetMsg("%s", err)
	}

	if err := dict.Select(); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return nil
}

func (d *DictService) SaveOne() error {
	panic("implement me")
}



