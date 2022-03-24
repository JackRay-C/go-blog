package service

import (
	"blog/app/model/po"
	"blog/app/model/vo"
	"blog/app/service/impl"
	"github.com/gin-gonic/gin"
)

type CommentService interface {
	ISelectOne(c *gin.Context, comment *po.Comment) error
	ISelectAll(c *gin.Context, pager *vo.Pager, comment *po.Comment) error
	ICreate(c *gin.Context, comment *po.Comment) error
	IDelete(c *gin.Context, comment *po.Comment) error
	ISelectAllByPostId(ctx *gin.Context) error
}



func NewCommentService() CommentService {
	return &impl.CommentServiceImpl{}
}

//func (c CommentService) SelectOne(common *domain.Comment) error {
//	db := global.DB.Model(&domain.Comment{})
//	if common.UserID != 0 {
//		db.Where("user_id=?", common.UserID)
//	}
//	if common.ID != 0 {
//		db.Where("id=?", common.ID)
//	}
//	if common.PostId != 0 {
//		db.Where("post_id=?", common.PostId)
//	}
//
//	err := db.First(&common).Error
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return errors.New("该评论不存在. ")
//	}
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (c *CommentService) SelectAll(page *pager.Pager, common *domain.Comment) error {
//	var comments []domain.Comment
//
//	db := global.DB.Model(&domain.Comment{}).Where(common)
//
//	if err := db.Count(&page.TotalRows).Error; err != nil {
//		return err
//	}
//
//	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
//	page.List = &comments
//
//	if err := db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&comments).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (c *CommentService) DeleteOne(common *domain.Comment) error {
//	db := global.DB.Model(&domain.Comment{})
//
//	if common.UserID != 0 {
//		db.Where("user_id=?", common.UserID)
//	}
//	if common.ID != 0 {
//		db.Where("id=?", common.ID)
//	}
//
//	var nc *domain.Comment
//	err := db.First(&nc).Error
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return errors.New("该评论不存在. ")
//	}
//	if err != nil {
//		return err
//	}
//
//	err = db.Delete(&common).Error
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// CreateOne 创建评论
//func (c *CommentService) CreateOne(common *domain.Comment) error {
//	return global.DB.Model(&domain.Comment{}).Create(&common).Error
//}
//
//// SelectPostComments 根据博客查询所有评论
//// todo: 暂时全部查询评论树，后续根据情况，如果评论过多可以改为分页查询评论
//func (c *CommentService) SelectPostComments(p *domain.Post, comments *[]*domain.Comment) error {
//	//1、 查询post ID是否存在
//	var common *domain.Comment
//	err := global.DB.Model(&domain.Post{}).Where("id=?", p.ID).First(&common).Error
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return errors.New("该博客不存在. ")
//	}
//	if err != nil {
//		return err
//	}
//
//	// 2、根据post id 查询所有评论
//	var all []*domain.Comment
//	if err := global.DB.Model(&domain.Comment{}).Where("post_id=?", p.ID).Find(&all).Error; err != nil {
//		return err
//	}
//
//	// 3、构造树形结构
//	m := make(map[int]*domain.Comment)
//
//	for _, common := range all {
//		if common.ParentID == 0 {
//			*comments = append(*comments, common)
//		}
//		m[common.ID] = common
//	}
//
//	for _, common := range all {
//		if common.ParentID != 0 {
//			if com, ok := m[common.ParentID]; ok {
//				com.Child = append(com.Child, common)
//			}
//		}
//	}
//
//	return nil
//}
