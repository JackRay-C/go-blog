package impl

import (
	"blog/pkg/global"
	"blog/pkg/model/bo"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/utils/auth"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"reflect"
)

type PostServiceImpl struct {
}

func (p *PostServiceImpl) ILike(c *gin.Context, post *bo.Post) error {
	go func() {
		if err := global.DB.Model(&po.Head{}).Where("id=?", post.Head.ID).Update("Likes", gorm.Expr("Likes+?", 1)).Error; err != nil {
			global.Log.Error("更新head.like失败: %s", err)
		}
	}()

	return nil
}

func (p *PostServiceImpl) ISelectOneWeb(c *gin.Context, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	if objT.String() != "*bo.Post" {
		return errors.New("obj must be *bo.Post")
	}

	// 1、查询head
	headV := objV.Elem().FieldByName("Head")
	if err := global.DB.Model(&po.Head{}).Where("id=?", headV.Elem().FieldByName("ID").Int()).First(headV.Addr().Interface()).Error; err != nil {
		return err
	}

	// 2、判断head是否是公开的
	if headV.Elem().FieldByName("Visibility").Int() == global.Private {
		// 2.1 判断是否登录
		if userID, ok := c.Get(global.SessionUserIDKey); !ok {
			return errors.New("require login. ")
		} else {
			// 2.2 登录判断userid是否相同
			if userID.(int64) != headV.Elem().FieldByName("UserID").Int() {
				return errors.New("forbidden. ")
			}
		}
	}

	// 3、查询repository
	if err := global.DB.Model(&po.Repository{}).Where("id=?", headV.Elem().FieldByName("RepositoryID").Int()).Find(objV.Elem().FieldByName("Repositories").Addr().Interface()).Error; err != nil {
		return err
	}

	// 4、查询history
	if err := global.DB.Model(&po.History{}).Where("head_id=?", headV.Elem().FieldByName("ID").Int()).Find(objV.Elem().FieldByName("Histories").Addr().Interface()).Error; err != nil {
		return err
	}

	// 5、更新post的visit
	go func() {
		if err := global.DB.Model(&po.Head{}).Where("id=?", headV.Elem().FieldByName("ID").Int()).Update("Views", gorm.Expr("Views+?", 1)).Error; err != nil {
			global.Log.Error("更新Head.Views失败: %s", err)
		}
	}()

	return nil
}

func (p *PostServiceImpl) ISelectOne(c *gin.Context, obj interface{}) error {
	// 查询单个博客
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	if objT.String() != "*bo.Post" {
		return errors.New("obj must be *bo.Post")
	}

	// 1、查询head
	headV := objV.Elem().FieldByName("Head")
	if err := global.DB.Model(&po.Head{}).Where("id=?", headV.Elem().FieldByName("ID").Int()).First(headV.Addr().Interface()).Error; err != nil {
		return err
	}

	// 2、判断head是否是公开的
	if headV.Elem().FieldByName("Visibility").Int() == global.Private {
		// 2.1 判断是否登录
		if userID, ok := c.Get(global.SessionUserIDKey); !ok {
			return errors.New("require login. ")
		} else {
			// 2.2 登录判断userid是否相同
			if userID.(int64) != headV.Elem().FieldByName("UserID").Int() {
				return errors.New("forbidden. ")
			}
		}
	}

	// 3、查询repository
	if err := global.DB.Model(&po.Repository{}).Where("id=?", headV.Elem().FieldByName("RepositoryID").Int()).Find(objV.Elem().FieldByName("Repositories").Addr().Interface()).Error; err != nil {
		return err
	}

	// 4、查询history
	if err := global.DB.Model(&po.History{}).Where("head_id=?", headV.Elem().FieldByName("ID").Int()).Find(objV.Elem().FieldByName("Histories").Addr().Interface()).Error; err != nil {
		return err
	}

	return nil
}

// ISelectList 查询所有博客列表
func (p *PostServiceImpl) ISelectList(c *gin.Context, pager *vo.Pager, obj interface{}) error {

	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	if objT.String() != "*bo.Post" {
		return errors.New("obj must be *bo.Post")
	}

	db := global.DB.Model(&po.Head{})

	if pager.Search != "" {
		//Where("MATCH (`title`,`markdown_content`,`html_content`,`description`) AGAINST ( ? IN NATURAL LANGUAGE MODE)", filter.Search)
		// 全文搜索
		//db.Where("MATCH (`title`,`markdown_content`,`html_content`,`description`) AGAINST ( ? IN NATURAL LANGUAGE MODE)", pager.Search)
	}

	headV := objV.Elem().FieldByName("Head")
	if headV.Elem().FieldByName("SubjectID").Int() != 0 {
		db.Where("subject_id=?", headV.Elem().FieldByName("SubjectID").Int())
	}

	return nil
}

func (p *PostServiceImpl) ICreateOne(c *gin.Context, obj interface{}) error {
	return nil
}

func (p *PostServiceImpl) IDeleteOne(c *gin.Context, obj interface{}) error {
	return nil
}

func (p *PostServiceImpl) IUpdateOne(c *gin.Context, obj interface{}, updateObj interface{}) error {
	return nil
}

func (p *PostServiceImpl) IUnscopeDelete(c *gin.Context, obj interface{}) error {
	return nil
}

func (p *PostServiceImpl) ISelectListWeb(c *gin.Context, pager *vo.Pager, obj interface{}) error {
	// 查询所有博客

	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	if objT.String() != "*bo.Post" {
		return errors.New("obj must be *bo.Post")
	}

	db := global.DB.Model(&po.Head{})

	if pager.Search != "" {
		// 全文搜索
	}

	headV := objV.Elem().FieldByName("Head")

	if headV.Elem().FieldByName("SubjectID").Int() != 0 {
		db.Where("subject_id=?", headV.Elem().FieldByName("SubjectID").Int())
	}

	if headV.Elem().FieldByName("UserID").Int() != 0 {
		db.Where("user_id=?", headV.Elem().FieldByName("UserID").Int())
	}

	db.Where("visibility=?", global.Public).Where("status=?", global.Publish)

	if auth.CheckLogin(c) {
		userId, _ := c.Get(global.SessionUserIDKey)
		db.Or("user_id=? and visibility=?", userId, global.Private)
	}

	if err := db.Count(&pager.TotalRows).Error; err != nil {
		return err
	}

	var heads []*po.Head
	offset := (pager.PageNo - 1) * pager.PageSize
	limit := pager.PageSize
	if err := db.Order(fmt.Sprintf("%s %s", pager.SortBy, pager.SortOrder)).Offset(offset).Limit(limit).Find(&heads).Error; err != nil {
		return errors.New(fmt.Sprintf("PostService find heads error: %s", err))
	}

	pager.PageCount = int((pager.TotalRows + int64(pager.PageSize) - 1) / int64(pager.PageSize))


	// 根据heads查找repository 和 history
	var posts []*bo.Post
	for _, head := range heads {
		post := &bo.Post{}
		post.Head = head
		if err := global.DB.Model(&po.Repository{}).Where("id=?", head.RepositoryID).Find(&post.Repositories).Error; err != nil {
			return errors.New(fmt.Sprintf("PostService find repository error: %s", err))
		}
		if err := global.DB.Model(&po.History{}).Where("head_id=?", head.ID).Find(&post.Histories).Error; err != nil {
			return errors.New(fmt.Sprintf("PostService find history error: %s", err))
		}
		posts = append(posts, post)
	}

	pager.MustList(&posts)

	return nil
}

//
//func (p *PostServiceImpl) ISelectAllWeb(c *gin.Context, page *vo.Pager, post *bo.Post) error {
//	var posts []*po.Post
//	db := global.DB.Model(&po.Post{})
//
//	if !auth.CheckLogin(c) {
//		db.Where("visibility=2")
//	} else {
//		db.Where("(user_id=? and visibility=1) or (visibility=2)")
//	}
//
//	if err := db.Count(&page.TotalRows).Error; err != nil {
//		return err
//	}
//
//	if err := db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&posts).Error; err != nil {
//		return err
//	}
//
//	page.MustList(&posts)
//
//	return nil
//}
//
//func (p *PostServiceImpl) IUpdateOneWeb(c *gin.Context, post *bo.Post) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) ISearchWeb(c *gin.Context, page *vo.Pager) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) IStaged(c *gin.Context, post *bo.Post) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) ICommit(c *gin.Context, post *bo.Post) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) IPublish(c *gin.Context, post *bo.Post) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) IPull(c *gin.Context, post *bo.Post) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) ISelectAll(c *gin.Context, page *vo.Pager, post *bo.Post) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) ISelectOne(c *gin.Context, post *bo.Post) error {
//	// 1、通过ID查询head
//	if err := global.DB.Model(&po.Head{}).Where("id=?", post.Head.ID).First(&post.Head).Error; err != nil {
//		return err
//	}
//
//	// 2、通过ID查询所有repository
//	if err := global.DB.Model(&po.Repository{}).Where("head_id=?", post.Head.ID).Find(&post.Repositories).Error; err != nil {
//		return err
//	}
//
//	// 3、通过ID查询history
//	if err := global.DB.Model(&po.History{}).Where("head_id=?", post.Head.ID).Find(&post.Histories).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// ICreateOne 创建博客，先创建repository，将head的repository id 指向创建的repository，并写入history，创建
//func (p *PostServiceImpl) ICreateOne(c *gin.Context, post *bo.Post) error {
//	// 获取userid
//	userId, _ := c.Get("current_user_id")
//
//	if err := global.DB.Transaction(func(tx *gorm.DB) error {
//		post.Repositories[0].UserId = userId.(int)
//		post.Repositories[0].CreatedAt = time.Now()
//		if err := tx.Model(&po.Repository{}).Create(post.Repositories[0]).Error; err != nil {
//			return err
//		}
//
//		post.Head.UserID = userId.(int)
//		post.Head.CreatedAt = time.Now()
//		post.Head.Status = global.Staged
//		post.Head.Visibility = global.Public
//		post.Head.RepositoryID = post.Repositories[0].ID
//		post.Head.Likes = 0
//		post.Head.Views = 0
//		post.Head.CoverImageId = po.DefaultPostCoverImageId
//		post.Head.SubjectID = 0
//
//		if err := tx.Model(&po.Head{}).Create(post.Head).Error; err != nil {
//			return err
//		}
//
//		post.Histories[0].UserID = userId.(int)
//		post.Histories[0].HeadID = post.Head.ID
//		post.Histories[0].RepositoryID = post.Repositories[0].ID
//		post.Histories[0].StagedAt = sql.NullTime{Time: time.Now()}
//		post.Histories[0].CommitedAt = sql.NullTime{Time: time.Now()}
//		post.Histories[0].PublishedAt = sql.NullTime{Valid: false}
//
//		if err := tx.Model(&po.History{}).Create(post.Histories[0]).Error; err != nil {
//			return err
//		}
//
//		return nil
//	}); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (p *PostServiceImpl) IUpdateOne(c *gin.Context, post *bo.Post) error {
//	panic("implement me")
//}
//
//func (p *PostServiceImpl) IDeleteOne(c *gin.Context, post *bo.Post) error {
//	panic("implement me")
//}
