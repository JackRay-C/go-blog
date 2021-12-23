package service

import (
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/pager"
	"blog/core/global"
	"blog/core/logger"
	"errors"
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

func (t *TagService) SelectOne(tag *domain.Tag)  error{
	db := global.DB.Model(&domain.Tag{}).Where("id=? ", tag.ID)

	if tag.UserId != 0 {
		db.Where("user_id=?", tag.UserId)
	}

	err := db.First(&tag).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该标签不存在. ")
	}
	return err
}

func (t *TagService) SelectAll(c *gin.Context, page *pager.Pager) error {
	var tags []*domain.Tag

	offset := (page.PageNo - 1) * page.PageSize
	limit := page.PageSize

	db := global.DB.Model(&domain.Tag{})

	if currentUserId, ok := c.Get("current_user_id"); ok {
		db.Where("user_id=?", currentUserId.(int))
	}

	if search, ok := c.GetQuery("search"); ok {
		db.Where("name like %%?%%", search)
	}

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

func (t *TagService) DeleteOne(userId int, id int) error {
	// 1、查询是否存在id的tag
	var tag *domain.Tag
	err := global.DB.Model(&domain.Tag{}).Where("id=?", id).First(&tag).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该标签不存在. ")
	}

	if err != nil {
		return err
	}

	// 2、删除tag
	return global.DB.Model(&domain.Tag{}).Where("id=? and user_id=?", id, userId).Delete(&domain.Tag{ID: id, UserId: userId}).Error
}


func (t *TagService) CreateOne(c *gin.Context, param *dto.AddTags) (tag *domain.Tag, err error) {
	userId, _ := c.Get("current_user_id")

	// 查询该用户下是否用同名的tag
	var newTag *domain.Tag
	err = global.DB.Model(&domain.Tag{}).Where("user_id=? and name=?", userId.(int), param.Name).First(&newTag).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
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
	return nil, errors.New(fmt.Sprintf("该标签[%s]已存在", param.Name))
}

func (t *TagService) UpdateOne(c *gin.Context, param *dto.PutTags) (tag *domain.Tag, err error) {
	userId, _ := c.Get("current_user_id")

	var newTag *domain.Tag
	err = global.DB.Model(&domain.Tag{}).Where("id=? and user_id=?", param.ID, userId).First(&newTag).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该标签不存在. ")
	}

	if err != nil {
		return nil, err
	}

	tag = &domain.Tag{
		ID: param.ID,
		Name: param.Name,
		UserId: userId.(int),
		CoverImage: param.CoverImage,
		Description: param.Description,
	}

	if err := global.DB.Model(&domain.Tag{}).Where("id=?", tag.ID).Omit("id", "user_id").Updates(tag).Error; err != nil {
		return nil, err
	}

	return tag, nil
}
