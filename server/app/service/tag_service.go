package service

import (
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/pager"
	"blog/core/global"
	"blog/core/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type TagService struct {
	log logger.Logger
}

func NewTagService() *TagService {
	return &TagService{
		log: global.Logger,
	}
}

func (t *TagService) SelectOne(id int) (tag *domain.Tag, err error) {
	if err = global.DB.Model(&domain.Tag{}).Where("id=?", id).First(&tag).Error; err != nil || err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return tag, nil
}

func (t *TagService) SelectAll(page *pager.Pager) error {
	var tags []*domain.Tag

	offset := (page.PageNo - 1) * page.PageSize
	limit := page.PageSize

	db := global.DB.Model(&domain.Tag{})

	if err := db.Count(&page.TotalRows).Error; err != nil {
		return err
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &tags

	if err := db.Offset(offset).Limit(limit).Find(&tags).Error; err != nil {
		return err
	}

	return nil
}

func (t *TagService) DeleteOne(id int) error {
	// 1、查询是否存在id的tag
	var tag *domain.Tag
	if err := global.DB.Model(&domain.Tag{}).Where("id=?", id).First(&tag).Error; err != nil || err == gorm.ErrRecordNotFound {
		return err
	}
	// 2、删除tag
	if err := global.DB.Model(&domain.Tag{}).Where("id=?", id).Delete(tag).Error; err != nil {
		return err
	}
	return nil
}

func (t *TagService) DeleteAll(ids []int) error {
	tag := &domain.Tag{}

	if err := tag.DeleteIds(ids); err != nil {
		return fmt.Errorf("failed delete all posts [%s]: %s", ids, err)
	}
	return nil
}

func (t *TagService) CreateOne(c *gin.Context, param *dto.AddTags) (tag *domain.Tag, err error) {
	userId, _ := c.Get("current_user_id")

	tag = &domain.Tag{
		Name:        param.Name,
		UserId:      userId.(int),
		CoverImage:  param.CoverImage,
		Description: param.Description,
		CreatedAt:   time.Now(),
	}

	if err := global.DB.Model(&domain.Tag{}).Create(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}

func (t *TagService) UpdateOne(c *gin.Context, param *dto.PutTags) (tag *domain.Tag, err error) {
	userId, _ := c.Get("current_user_id")

	if err := global.DB.Model(&domain.Tag{}).Where("id=?", param.ID).First(&tag).Error; err != nil || err == gorm.ErrRecordNotFound {
		return nil, err
	}

	tag.UserId = userId.(int)
	tag.Name = param.Name
	tag.Description = param.Description
	tag.CoverImage = param.CoverImage
	tag.UpdatedAt = time.Now()
	if err := global.DB.Model(&domain.Tag{}).Where("id=?", tag.ID).Save(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}
