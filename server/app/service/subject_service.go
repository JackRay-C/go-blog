package service

import (
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/model/vo"
	"blog/app/pager"
	"blog/app/response"
	"blog/core/global"
	"blog/core/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type SubjectService struct {
	Log    logger.Logger
	engine *gorm.DB
}

func NewSubjectService() *SubjectService {
	return &SubjectService{
		Log:    global.Logger,
		engine: global.DB,
	}
}

func (s *SubjectService) DeleteAll(ids []int) error {
	subject := &domain.Subject{}

	if err := subject.DeleteIds(ids); err != nil {
		return fmt.Errorf("failed delete all tag [%s]: %s", ids, err)
	}
	return nil
}

func (s *SubjectService) SelectOneById(id int) (*vo.VSubject, error) {
	var subject *domain.Subject
	if err := global.DB.Model(&domain.Subject{}).Where("id=?", id).First(&subject).Error; err != nil {
		return nil, err
	}

	var avatar *domain.File
	if err := global.DB.Model(&domain.File{}).Where("id=?", subject.Avatar).First(&avatar).Error; err != nil {
		return nil, err
	}

	var coverImage *domain.File
	if err := global.DB.Model(&domain.File{}).Where("id=?", subject.CoverImage).First(&coverImage).Error; err != nil {
		return nil, err
	}

	service := NewUserService()
	if user, err := service.SelectOneById(subject.UserID); err != nil {
		return nil, err
	} else {
		return &vo.VSubject{
			ID:          subject.ID,
			Title:       subject.Title,
			Avatar:      avatar,
			CoverImage:  coverImage,
			Description: subject.Description,
			Visibility:  subject.Visibility,
			UserID:      subject.UserID,
			User:        user,
			Views:       subject.Views,
			CreatedAt:   subject.CreatedAt,
		}, nil
	}
}

func (s *SubjectService) SelectAllWeb(c *gin.Context, page *pager.Pager, filter *dto.ListSubjects) error {
	var subjects []*domain.Subject
	var vSubjects []*vo.VSubject

	offset := (filter.PageNo - 1) * filter.PageSize
	limit := filter.PageSize
	var count int64
	// 1、判断是否是搜索
	if filter.Search != "" {
		db := global.DB.Model(&domain.Subject{}).Where("name like %s", filter.Search).Or("description like %s", filter.Search)
		isLogin, exists := c.Get("is_login")
		if exists {
			if !isLogin.(bool) {
				if filter.UserId == 0 {
					db.Where("visibility=2")
				} else {
					db.Where("user_id =? and visibility=2", filter.UserId)
				}
			} else {
				userId, e := c.Get("current_user_id")
				if e {
					if filter.UserId == 0 {
						db.Where("visibility=2").Or("user_id=? and visibility=1", userId.(int))
					} else {
						if filter.UserId == userId.(int) {
							db.Where("user_id=?", filter.UserId)
						} else {
							db.Where("user_id=? and visibility=?", filter.UserId, 2)
						}
					}
				}
			}
		}

		if err := db.Count(&count).Error; err != nil {
			return err
		}

		if err := db.Offset(offset).Limit(limit).Select("id").Find(&subjects).Error; err != nil {
			return err
		}

	} else {
		// 查询所有公开专题
		db := global.DB.Model(&domain.Subject{})
		isLogin, exists := c.Get("is_login")
		if exists {
			if !isLogin.(bool) {
				if filter.UserId == 0 {
					db.Where("visibility=2")
				} else {
					db.Where("user_id =? and visibility=2", filter.UserId)
				}
			} else {
				userId, e := c.Get("current_user_id")
				if e {
					if filter.UserId == 0 {
						db.Where("visibility=2").Or("user_id=? and visibility=1", userId.(int))
					} else {
						if filter.UserId == userId.(int) {
							db.Where("user_id=?", filter.UserId)
						} else {
							db.Where("user_id=? and visibility=?", filter.UserId, 2)
						}
					}
				}
			}
		}

		if err := db.Count(&count).Error; err != nil {
			return err
		}
		if err := db.Order("created_at desc").Offset(offset).Limit(limit).Select("ID").Find(&subjects).Error; err != nil {
			return err
		}
	}

	for _, subject := range subjects {
		if vSubject, err := s.SelectOneById(subject.ID); err != nil {
			return err
		} else {
			vSubjects = append(vSubjects, vSubject)
		}
	}

	page.PageNo = filter.PageNo
	page.PageSize = filter.PageSize
	page.TotalRows = count
	page.PageCount = int((count + int64(page.PageSize) - 1)/int64(page.PageSize))
	page.List = &vSubjects

	return nil
}

func (s *SubjectService) SelectAll(page *pager.Pager, subject *domain.Subject) error {
	var subjects []domain.Subject

	if err := subject.Count(&page.TotalRows); err != nil {
		s.Log.Errorf("%s", err)
		return response.DatabaseSelectError.SetMsg("query subject count failed: %s", err)
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &subjects
	if err := subject.List(&subjects, (page.PageNo-1)*page.PageSize, page.PageSize); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	return nil
}

func (s *SubjectService) DeleteOne(subject *domain.Subject) error {

	if err := subject.Delete(); err != nil {
		return response.DatabaseDeleteError.SetMsg("%s", err)
	}
	return nil
}

func (s *SubjectService) CreateOne(subject *domain.Subject) error {
	// 1、根据名称查询是否存在
	if err := subject.Select(); err == gorm.ErrRecordNotFound {
		// 2、如果不存在，插入
		if err := subject.Insert(); err != nil {
			return response.DatabaseInsertError.SetMsg("%s", err)
		}
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	} else {
		return response.RecoreExisted.SetMsg("该专题已经存在. ")
	}
	return nil
}

func (s *SubjectService) UpdateOne(subject *domain.Subject) error {
	if err := subject.Update(); err != nil {
		return response.DatabaseUpdateError.SetMsg("%s", err)
	}

	if err := subject.Select(); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return nil
}

func (s *SubjectService) SaveOne(subject *domain.Subject) error {
	if err := subject.Select(); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	subject.UpdatedAt = time.Now()
	if err := subject.Save(); err != nil {
		return response.DatabaseUpdateError.SetMsg("%s", err)
	}
	return nil
}