package impl

import (
	"blog/pkg/model/common"
)

type SubjectServiceImpl struct {
	common.BaseServiceImpl
}

//
//func (s *SubjectServiceImpl) ISelectList(c *gin.Context, pager *vo.Pager, subject *po.Subject) error {
//	var count int64
//	var subjects []*po.Subject
//	offset := (pager.PageNo - 1) * pager.PageSize
//	limit := pager.PageSize
//
//	// 1、获取用户名
//	userId, _ := c.Get(common.SessionUserIDKey)
//
//	db := global.DB.Model(&po.Subject{}).Where("user_id=?", userId)
//
//	if subject.Visibility != 0 {
//		db.Where("visibility=?", subject.Visibility)
//	}
//
//	if err := db.Count(&pager.TotalRows).Error; err != nil {
//		return vo.DatabaseSelectError.SetMsg("%s", err)
//	}
//
//	if err := db.Offset(offset).Limit(limit).Find(&subjects).Error; err != nil {
//		return vo.DatabaseSelectError.SetMsg("%s", err)
//	}
//
//	pager.PageCount = int((count + int64(pager.PageSize) - 1)/int64(pager.PageSize))
//	pager.MustList(subjects)
//
//	return nil
//}
//
//func (s *SubjectServiceImpl) ICreateOne(c *gin.Context, subject *po.Subject) error {
//	userId, _ := c.Get(common.SessionUserIDKey)
//	subject.UserID = userId.(int)
//	if subject.Avatar == 0 {
//		subject.Avatar = po.DefaultSubjectAvatarId
//	}
//	if subject.CoverImage == 0 {
//		subject.CoverImage = po.DefaultSubjectCoverImageId
//	}
//
//	db := global.DB.Model(&po.Subject{}).Where(&po.Subject{Title: subject.Title}).FirstOrCreate(subject)
//
//	if err := db.Error; err != nil {
//		return vo.DatabaseInsertError.SetMsg("%s", err)
//	}
//
//	if db.RowsAffected == 0 {
//		return vo.RecoreExisted.SetMsg("this subject title is exists: %s", subject.Title)
//	}
//
//	return nil
//}

