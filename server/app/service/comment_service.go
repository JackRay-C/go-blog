package service

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/app/response"
	"blog/core/global"
	"errors"

	"gorm.io/gorm"
)

type CommentService struct {
}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (c *CommentService) SelectOne(comment *domain.Comment) error {
	db := global.DB.Model(&domain.Comment{})
	if comment.UserID != 0 {
		db.Where("user_id=?", comment.UserID)
	}
	if comment.ID != 0 {
		db.Where("id=?", comment.ID)
	}
	if comment.PostId != 0 {
		db.Where("post_id=?", comment.PostId)
	}

	var nc *domain.Comment
	err := db.First(&nc).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该评论不存在. ")
	}
	if err != nil {
		return err
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
	db := global.DB.Model(&domain.Comment{})

	if comment.UserID != 0 {
		db.Where("user_id=?", comment.UserID)
	}
	if comment.ID != 0 {
		db.Where("id=?", comment.ID)
	}

	var nc *domain.Comment
	err := db.First(&nc).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该评论不存在. ")
	}
	if err != nil {
		return err
	}

	err = db.Delete(&comment).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *CommentService) CreateOne(comment *domain.Comment) error {
	return global.DB.Model(&domain.Comment{}).Create(&comment).Error
}

// SelectPostComments 根据博客查询所有评论
// todo: 暂时全部查询评论树，后续根据情况，如果评论过多可以改为分页查询评论
func (c *CommentService) SelectPostComments(p *domain.Post, comments *[]*domain.Comment) error {
	//1、 查询post ID是否存在
	var comment *domain.Comment
	err := global.DB.Model(&domain.Post{}).Where("id=?", p.ID).First(&comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该博客不存在. ")
	}
	if err != nil {
		return err
	}

	// 2、根据post id 查询所有评论
	var all []*domain.Comment
	if err := global.DB.Model(&domain.Comment{}).Where("post_id=?", p.ID).Find(&all).Error; err != nil {
		return err
	}

	// 3、构造树形结构
	m := make(map[int]*domain.Comment)

	for _, comment := range all {
		if comment.ParentID == 0 {
			*comments = append(*comments, comment)
		}
		m[comment.ID] = comment
	}

	for _, comment := range all {
		if comment.ParentID != 0 {
			if com, ok := m[comment.ParentID]; ok {
				com.Child = append(com.Child, comment)
			}
		}
	}

	return nil
}
