package service

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/app/response"
	"gorm.io/gorm"
)

type CommentService struct {
	
}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (c *CommentService) SelectOne(comment *domain.Comment) error {
	if err := comment.Select(); err == gorm.ErrRecordNotFound {
		return response.RecordNotFound
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return nil
}

func (c *CommentService) SelectAll(page *pager.Pager, comment *domain.Comment) error {
	var comments []domain.Comment

	if err := comment.Count(&page.TotalRows); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &comments

	if err := comment.List(&comments, (page.PageNo-1)*page.PageSize, page.PageSize); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	return nil
}

func (c *CommentService) DeleteOne(comment *domain.Comment) error {
	if err := comment.Select(); err == gorm.ErrRecordNotFound {
		return response.RecordNotFound
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	if err := comment.Delete(); err != nil {
		return response.DatabaseDeleteError.SetMsg("%s", err)
	}
	return nil
}

func (c *CommentService) DeleteAll(ids []int) error {
	panic("implement me")
}

func (c *CommentService) CreateOne(comment *domain.Comment) error {
	panic("implement me")
}

func (c *CommentService) UpdateOne() error {
	panic("implement me")
}

func (c *CommentService) SaveOne() error {
	panic("implement me")
}

