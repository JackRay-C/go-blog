package impl

import (
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostServiceImpl struct {
	common.BaseServiceImpl
}

func (p PostServiceImpl) ILike(c *gin.Context, post *po.Post) error {
	return global.DB.Model(&po.Post{}).Select("likes").Omit("updated_at").Where("id=?", post.ID).Update("likes", gorm.Expr("likes + 1")).Error
}

func (p PostServiceImpl) IIncrementView(c *gin.Context, post *po.Post) error {
	return global.DB.Model(&po.Post{}).Omit("updated_at").Where("id=?", post.ID).Update("views", gorm.Expr("views + 1")).Error
}

func (p PostServiceImpl) IUnLike(c *gin.Context, post *po.Post) error {
	return global.DB.Model(&po.Post{}).Select("likes").Omit("updated_at").Where("id=?", post.ID).Update("likes", gorm.Expr("likes - 1")).Error
}

func (p PostServiceImpl) IDecrementViews(c *gin.Context, post *po.Post) error {
	return global.DB.Model(&po.Post{}).Omit("updated_at").Where("id=?", post.ID).Update("views", gorm.Expr("views - 1")).Error
}

//func (p *PostServiceImpl) CreateVersion(c *gin.Context, post *po.Post, version *po.Version) error {
//	if err := global.DB.Model(&po.Version{}).Create(version).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (p *PostServiceImpl) ResetVersion(c *gin.Context, post *po.Post) error {
//	// 1、获取post version信息
//	version := po.Version{}
//	if err := global.DB.Model(&po.Version{}).Where("post_id=? and id=?", post.ID, post.Version).First(&version).Error; err != nil {
//		return err
//	}
//
//	// 2、将post的信息更新为version的信息
//	if err := global.DB.Model(&po.Post{}).Where("id=?", post.ID).Updates(&po.Post{MarkdownContent: version.MarkdownContent, HtmlContent: version.HtmlContent, SubjectID: version.SubjectID, Version: version.ID}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (p *PostServiceImpl) ClearVersions(c *gin.Context, post *po.Post) error {
//	// 将post的所有版本清除
//	if err := global.DB.Model(&po.Version{}).Where("post_id=?", post.ID).Delete(&po.Version{}).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (p *PostServiceImpl) DeleteVersion(c *gin.Context, post *po.Post, version int64) error {
//	if err := global.DB.Model(&po.Version{}).Where("post_id=? and id=?", post.ID, version).Delete(&po.Version{}).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (p *PostServiceImpl) GetVersions(c *gin.Context, post *po.Post) error {
//	var versions *po.Version
//	if err := global.DB.Model(&po.Version{}).Where("post_id=?").Find(&versions).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (p *PostServiceImpl) CurrentVersion(c *gin.Context, post *po.Post) error {
//	if err := global.DB.Model(&po.Post{}).Where("id=?", post.ID).First(&post).Error; err != nil {
//		return err
//	}
//	if err := global.DB.Model(&po.Version{}).Where("id=?", post.Version).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (p *PostServiceImpl) CreateDraft(c *gin.Context, draft *po.Draft) error {
//	// 创建草稿
//	if err := global.DB.Model(&po.Draft{}).Create(&draft).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (p *PostServiceImpl) DeleteDraft(c *gin.Context, draft *po.Draft) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) PublishDraft(c *gin.Context,post *po.Post, draft *po.Draft) error {
//	// 发布草稿
//	// 更新post信息为草稿信息
//	values := &po.Post{Title: draft.Title, MarkdownContent: draft.MarkdownContent, HtmlContent: draft.HtmlContent, Status: global.Publish}
//	if err := global.DB.Model(&po.Post{}).Where("id=?", draft.PostId).Updates(values).Error; err != nil {
//		return err
//	}
//
//	// 创建version为草稿信息
//	if err := p.CreateVersion(c, values, &po.Version{PostID: draft.PostId, MarkdownContent: draft.MarkdownContent, HtmlContent: draft.HtmlContent}); err != nil {
//		return err
//	}
//
//	// 删除草稿信息
//	if err := p.DeleteDraft(c, draft); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (p *PostServiceImpl) ILike(c *gin.Context, post *bo.Post) error {
//	go func() {
//		if err := global.DB.Model(&po.Head{}).Where("id=?", post.Head.ID).Update("Likes", gorm.Expr("Likes+?", 1)).Error; err != nil {
//			global.Log.Error("更新head.like失败: %s", err)
//		}
//	}()
//
//	return nil
//}
//
//func (p *PostServiceImpl) ISelectOneWeb(c *gin.Context, obj interface{}) error {
//	objT := reflect.TypeOf(obj)
//	objV := reflect.ValueOf(obj)
//
//	if objT.String() != "*bo.Post" {
//		return errors.New("obj must be *bo.Post")
//	}
//
//	// 1、查询head
//	headV := objV.Elem().FieldByName("Head")
//	if err := global.DB.Model(&po.Head{}).Where("id=?", headV.Elem().FieldByName("ID").Int()).First(headV.Addr().Interface()).Error; err != nil {
//		return err
//	}
//
//	// 2、判断head是否是公开的
//	if headV.Elem().FieldByName("Visibility").Int() == global.Private {
//		// 2.1 判断是否登录
//		if userID, ok := c.Get(global.SessionUserIDKey); !ok {
//			return errors.New("require login. ")
//		} else {
//			// 2.2 登录判断userid是否相同
//			if userID.(int64) != headV.Elem().FieldByName("UserID").Int() {
//				return errors.New("forbidden. ")
//			}
//		}
//		if auth.CheckLogin(c) {
//			claims, err := token.ParseAccessToken(c.GetHeader(global.RequestHeaderTokenKey))
//			if err != nil {
//				return err
//			}
//			// 2.2 登录判断userid是否相同
//			if int64(claims.UserId) != headV.Elem().FieldByName("UserID").Int() {
//				return errors.New("forbidden. ")
//			}
//		}
//	}
//
//	// 3、查询repository
//	if err := global.DB.Model(&po.Repository{}).Where("id=?", headV.Elem().FieldByName("RepositoryID").Int()).Find(objV.Elem().FieldByName("Repositories").Addr().Interface()).Error; err != nil {
//		return err
//	}
//
//	// 4、查询history
//	if err := global.DB.Model(&po.History{}).Where("head_id=?", headV.Elem().FieldByName("ID").Int()).Find(objV.Elem().FieldByName("Histories").Addr().Interface()).Error; err != nil {
//		return err
//	}
//
//	// 5、更新post的visit
//	go func() {
//		if err := global.DB.Model(&po.Head{}).Where("id=?", headV.Elem().FieldByName("ID").Int()).Update("Views", gorm.Expr("Views+?", 1)).Error; err != nil {
//			global.Log.Error("更新Head.Views失败: %s", err)
//		}
//	}()
//
//	return nil
//}
//
//func (p *PostServiceImpl) ISelectOne(c *gin.Context, obj interface{}) error {
//	// 查询单个博客
//	objT := reflect.TypeOf(obj)
//	objV := reflect.ValueOf(obj)
//
//	if objT.String() != "*bo.Post" {
//		return errors.New("obj must be *bo.Post")
//	}
//
//	// 1、查询head
//	headV := objV.Elem().FieldByName("Head")
//
//	claims, err := token.ParseAccessToken(c.GetHeader(global.RequestHeaderTokenKey))
//	if err != nil {
//		return err
//	}
//	if err := global.DB.Model(&po.Head{}).Where("id=? and user_id=?", headV.Elem().FieldByName("ID").Int(), claims.UserId).First(headV.Addr().Interface()).Error; err != nil {
//		return err
//	}
//
//	// 3、查询repository
//	if err := global.DB.Model(&po.Repository{}).Where("id=?", headV.Elem().FieldByName("RepositoryID").Int()).Find(objV.Elem().FieldByName("Repositories").Addr().Interface()).Error; err != nil {
//		return err
//	}
//
//	// 4、查询history
//	if err := global.DB.Model(&po.History{}).Where("head_id=?", headV.Elem().FieldByName("ID").Int()).Find(objV.Elem().FieldByName("Histories").Addr().Interface()).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// ISelectList 查询所有博客列表
//func (p *PostServiceImpl) ISelectList(c *gin.Context, pager *vo.Pager, obj interface{}) error {
//
//	objT := reflect.TypeOf(obj)
//	objV := reflect.ValueOf(obj)
//
//	if objT.String() != "*bo.Post" {
//		return errors.New("obj must be *bo.Post")
//	}
//
//	db := global.DB.Model(&po.Head{})
//
//	if pager.Search != "" {
//		//Where("MATCH (`title`,`markdown_content`,`html_content`,`description`) AGAINST ( ? IN NATURAL LANGUAGE MODE)", filter.Search)
//		// 全文搜索
//		//db.Where("MATCH (`title`,`markdown_content`,`html_content`,`description`) AGAINST ( ? IN NATURAL LANGUAGE MODE)", pager.Search)
//	}
//
//	headV := objV.Elem().FieldByName("Head")
//	if headV.Elem().FieldByName("SubjectID").Int() != 0 {
//		db.Where("subject_id=?", headV.Elem().FieldByName("SubjectID").Int())
//	}
//
//	if headV.Elem().FieldByName("Status").Int() != 0 {
//		db.Where("status=?", headV.Elem().FieldByName("Status").Int())
//	}
//
//	claims, err := token.ParseAccessToken(c.GetHeader(global.RequestHeaderTokenKey))
//	if err != nil {
//		return err
//	}
//
//	db.Where("user_id=?", claims.UserId)
//
//	if err := db.Count(&pager.TotalRows).Error; err != nil {
//		return err
//	}
//
//	var heads []*po.Head
//	offset := (pager.PageNo - 1) * pager.PageSize
//	limit := pager.PageSize
//	if err := db.Order(fmt.Sprintf("%s %s", pager.SortBy, pager.SortOrder)).Offset(offset).Limit(limit).Find(&heads).Error; err != nil {
//		return errors.New(fmt.Sprintf("PostService find heads error: %s", err))
//	}
//
//	// 根据heads查找repository 和 history
//	var posts []*bo.Post
//	for _, head := range heads {
//		post := &bo.Post{}
//		post.Head = head
//		if err := global.DB.Model(&po.Repository{}).Where("id=?", head.RepositoryID).Find(&post.Repositories).Error; err != nil {
//			return errors.New(fmt.Sprintf("PostService find repository error: %s", err))
//		}
//		if err := global.DB.Model(&po.History{}).Where("head_id=?", head.ID).Find(&post.Histories).Error; err != nil {
//			return errors.New(fmt.Sprintf("PostService find history error: %s", err))
//		}
//		posts = append(posts, post)
//	}
//
//	pager.MustList(&posts)
//
//	return nil
//}
//
//func (p *PostServiceImpl) ICreateOne(c *gin.Context, obj interface{}) error {
//	return nil
//}
//
//func (p *PostServiceImpl) IDeleteOne(c *gin.Context, obj interface{}) error {
//	return nil
//}
//
//func (p *PostServiceImpl) IUpdateOne(c *gin.Context, obj interface{}, updateObj interface{}) error {
//	return nil
//}
//
//func (p *PostServiceImpl) IUnscopeDelete(c *gin.Context, obj interface{}) error {
//	return nil
//}
//
//func (p *PostServiceImpl) ISelectListWeb(c *gin.Context, pager *vo.Pager, obj interface{}) error {
//	// 查询所有博客
//
//	objT := reflect.TypeOf(obj)
//	objV := reflect.ValueOf(obj)
//
//	if objT.String() != "*bo.Post" {
//		return errors.New("obj must be *bo.Post")
//	}
//
//	db := global.DB.Model(&po.Head{})
//
//	if pager.Search != "" {
//		// 1、全文搜索repository，查找出repository id
//
//		// 2、根据repository_id查找head
//
//		// 3、根据head查找history
//
//		//db.Where("MATCH (`title`,`markdown_content`,`html_content`,`description`) AGAINST ( ? IN NATURAL LANGUAGE MODE)", pager.Search)
//	}
//
//	headV := objV.Elem().FieldByName("Head")
//
//	if headV.Elem().FieldByName("SubjectID").Int() != 0 {
//		db.Where("subject_id=?", headV.Elem().FieldByName("SubjectID").Int())
//	}
//
//	if headV.Elem().FieldByName("UserID").Int() != 0 {
//		db.Where("user_id=?", headV.Elem().FieldByName("UserID").Int())
//	}
//
//	db.Where("visibility=?", global.Public).Where("status=?", global.Publish)
//
//	if auth.CheckLogin(c) {
//		claims, err := token.ParseAccessToken(c.GetHeader(global.RequestHeaderTokenKey))
//		if err != nil {
//			return err
//		}
//		db.Or("user_id=? and visibility=?", claims.UserId, global.Private)
//	}
//
//	if err := db.Count(&pager.TotalRows).Error; err != nil {
//		return err
//	}
//
//	var heads []*po.Head
//	offset := (pager.PageNo - 1) * pager.PageSize
//	limit := pager.PageSize
//	if err := db.Order(fmt.Sprintf("%s %s", pager.SortBy, pager.SortOrder)).Offset(offset).Limit(limit).Find(&heads).Error; err != nil {
//		return errors.New(fmt.Sprintf("PostService find heads error: %s", err))
//	}
//
//	// 根据heads查找repository 和 history
//	var posts []*bo.Post
//	for _, head := range heads {
//		post := &bo.Post{}
//		post.Head = head
//		if err := global.DB.Model(&po.Repository{}).Where("id=?", head.RepositoryID).Find(&post.Repositories).Error; err != nil {
//			return errors.New(fmt.Sprintf("PostService find repository error: %s", err))
//		}
//		if err := global.DB.Model(&po.History{}).Where("head_id=?", head.ID).Find(&post.Histories).Error; err != nil {
//			return errors.New(fmt.Sprintf("PostService find history error: %s", err))
//		}
//		posts = append(posts, post)
//	}
//
//	pager.MustList(&posts)
//
//	return nil
//}
