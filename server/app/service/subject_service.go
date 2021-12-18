package service

import (
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/model/vo"
	"blog/app/pager"
	"blog/app/response"
	"blog/core/global"
	"blog/core/logger"
	"errors"
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

	db := global.DB.Model(&domain.Subject{})

	// 判断是否登录
	isLogin, exists := c.Get("is_login")
	if exists {
		if !isLogin.(bool) {
			if filter.UserId == 0 {
				db.Where("visibility=2")
			} else {
				db.Where("user_id=? and visibility=2", filter.UserId)
			}
		} else {
			userId, _ := c.Get("current_user_id")
			if filter.UserId == 0 {
				db.Where("visibility=2").Or("user_id=? and visibility=1", userId.(int))
			}else {
				if filter.UserId == userId.(int) {
					db.Where("user_id=?", filter.UserId)
				} else {
					db.Where("user_id=? and visibility=?", filter.UserId, 2)
				}
			}
		}
	}

	if filter.Search != "" {
		db.Where("name like ?", fmt.Sprintf("%%%s%%", filter.Search))
	}

	if err := db.Count(&count).Error; err != nil {
		return err
	}

	if err := db.Order("created_at desc").Offset(offset).Limit(limit).Select("id").Find(&subjects).Error; err != nil {
		return err
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

func (s *SubjectService) SelectAll(c *gin.Context,page *pager.Pager, filter *dto.ListSubjects) error {
	var subjects []*domain.Subject
	var vSubjects []*vo.VSubject

	offset := (filter.PageNo - 1) * filter.PageSize
	limit := filter.PageSize
	var count int64

	userId, _ := c.Get("current_user_id")
	db := global.DB.Model(&domain.Subject{}).Where("user_id=?", userId)

	if filter.Visibility != 0 {
		db.Where("visibility=?", filter.Visibility)
	}

	// 搜索
	if filter.Search != "" {
		db.Where("name like %s", fmt.Sprintf("%%%s%%", filter.Search))
	}

	if err := db.Count(&count).Error; err != nil {
		return err
	}

	if err := db.Order("created_at desc").Offset(offset).Limit(limit).Select("id").Find(&subjects).Error; err != nil {
		return err
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


func (s *SubjectService) DeleteOne(c *gin.Context, subjectId int) error {
	var subject *domain.Subject

	// 1、根据user_id id查询subject
	userId, _ := c.Get("current_user_id")
	db := global.DB.Model(&domain.Subject{}).Where("user_id=? and id=?", userId, subjectId)
	if err := db.First(&subject).Error; err  != nil {
		return err
	}

	// 2、根据id删除subject
	if err:= db.Delete(subject).Error; err != nil {
		return err
	}

	return nil
}

func (s *SubjectService) CreateOne(c *gin.Context, param *dto.AddSubjects) (*domain.Subject, error) {
	// 1、获取当前用户
	userId, _ := c.Get("current_user_id")

	// 2、查询是否存在同名专题
	db := global.DB.Model(&domain.Subject{})

	var subject *domain.Subject
	if err := db.Where("title=?", param.Title).First(&subject).Error; err != gorm.ErrRecordNotFound {
		return nil, errors.New("该专题已存在. ")
	}

	// 3、创建专题
	subject = &domain.Subject{
		Title:       param.Title,
		Avatar:      param.AvatarId,
		CoverImage:  param.CoverImageId,
		Description: param.Description,
		Visibility:  param.Visibility,
		UserID:      userId.(int),
		Views:       0,
		CreatedAt:   time.Now(),
	}
	if err := db.Create(subject).Error; err != nil {
		return nil, err
	}
	return subject, nil
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
