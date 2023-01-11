package common

import (
	"blog/pkg/global"
	"blog/pkg/model/vo"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/token"
	"blog/pkg/utils/transform"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

type BaseService interface {
	ISelectOne(c *gin.Context, obj interface{}) error
	ISelectList(c *gin.Context, pager *vo.Pager, obj interface{}) error
	ICreateOne(c *gin.Context, obj interface{}) error
	IDeleteOne(c *gin.Context, obj interface{}) error
	IUpdateOne(c *gin.Context, obj interface{}, updateObj interface{}) error
	IUnscopeDelete(c *gin.Context, obj interface{}) error

	ISelectOneWeb(c *gin.Context, obj interface{}) error
	ISelectListWeb(c *gin.Context, pager *vo.Pager, obj interface{}) error
}

type BaseServiceImpl struct{}

// ISelectOne select obj data from database
func (b *BaseServiceImpl) ISelectOne(c *gin.Context, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
	}

	db := global.DB.Model(objV.Interface())

	if err := db.Where(objV).First(objV.Interface()).Error; err != nil {
		return err
	}

	return nil
}

// ISelectList select page data from database
func (b *BaseServiceImpl) ISelectList(c *gin.Context, pager *vo.Pager, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
	}

	objV := reflect.ValueOf(obj)

	var results []map[string]interface{}
	global.Log.Info(objV)

	db := global.DB.Debug().Model(objV.Interface()).Where(objV.Interface())

	if userID, ok := c.Get(global.SessionUserIDKey); !ok {
		return errors.New("require login. ")
	} else {
		if _, ok := objT.FieldByName("UserID"); ok {
			db.Where("user_id=?", userID)
		}
	}

	if pager.Search != "" && pager.FullIndex {
		// 全文索引
		// 1、判断obj是否有name
		if _, ok := objT.FieldByName("Name"); ok {
			db.Where(fmt.Sprintf("name like '%%%s%%'", pager.Search))
		}

		// 2、判断obj是否有title
		if _, ok := objT.FieldByName("Title"); ok {
			db.Or(fmt.Sprintf("title like '%%%s%%'", pager.Search))
		}
		// 3、判断obj是否有content
		if _, ok := objT.FieldByName("Content"); ok {
			db.Or(fmt.Sprintf("content like '%%%s%%'", pager.Search))
		}
		// 4、判断obj是否有description
		if _, ok := objT.FieldByName("Description"); ok {
			db.Or(fmt.Sprintf("description like '%%%s%%'", pager.Search))
		}

		// 5、判断obj是否有comment
		if _, ok := objT.FieldByName("Comment"); ok {
			db.Or(fmt.Sprintf("comment like '%%%s%%'", pager.Search))
		}

		// 6、判断obj是否有markdown_content
		if _, ok := objT.FieldByName("MarkdownContent"); ok {
			db.Or(fmt.Sprintf("markdown_content like '%%%s%%'", pager.Search))
		}

		// 7、判断obj是否有html_content
		if _, ok := objT.FieldByName("HtmlContent"); ok {
			db.Or(fmt.Sprintf("html_content like '%%%s%%'", pager.Search))
		}
	} else if pager.Search != "" && !pager.FullIndex {
		// 1、判断obj是否有name
		if _, ok := objT.FieldByName("Name"); ok {
			db.Where(fmt.Sprintf("name like '%%%s%%'", pager.Search))
		}

		// 2、判断obj是否有title
		if _, ok := objT.FieldByName("Title"); ok {
			db.Where(fmt.Sprintf("title like '%%%s%%'", pager.Search))
		}
	}

	if err := db.Count(&pager.TotalRows).Error; err != nil {
		return err
	}

	offset := (pager.PageNo - 1) * pager.PageSize
	limit := pager.PageSize
	if pager.SortBy == "" {
		pager.MustSort(c)
	}

	if err := db.Order(fmt.Sprintf("%s %s", pager.SortBy, pager.SortOrder)).Offset(offset).Limit(limit).Find(&results).Error; err != nil {
		return err
	}

	pager.PageCount = int((pager.TotalRows + int64(pager.PageSize) - 1) / int64(pager.PageSize))
	pager.MustList(&results)

	return nil
}

func (b *BaseServiceImpl) ICreateOne(c *gin.Context, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
	}

	objV := reflect.ValueOf(obj)

	db := global.DB.Model(objV.Interface())

	if userID, ok := c.Get(global.SessionUserIDKey); !ok {
		return errors.New("require login. ")
	} else {
		if _, ok := objT.FieldByName("UserID"); ok {
			if objV.Kind() == reflect.Ptr {
				objV.Elem().FieldByName("UserID").SetInt(userID.(int64))
			} else {
				objV.FieldByName("UserID").SetInt(userID.(int64))
			}
		}
	}

	timestamp := time.Now().UnixMilli()
	if _, ok := objT.FieldByName("CreatedAt"); ok {
		if objV.Kind() == reflect.Ptr {
			objV.Elem().FieldByName("CreatedAt").Set(reflect.ValueOf(&timestamp))
		} else {
			objV.FieldByName("CreatedAt").Set(reflect.ValueOf(&timestamp))
		}
	}
	if _, ok := objT.FieldByName("UpdatedAt"); ok {
		if objV.Kind() == reflect.Ptr {
			objV.Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(&timestamp))
		} else {
			objV.FieldByName("UpdatedAt").Set(reflect.ValueOf(&timestamp))
		}
	}

	if err := db.Create(objV.Interface()).Error; err != nil {
		return err
	}

	return nil
}

func (b *BaseServiceImpl) IDeleteOne(c *gin.Context, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
	}

	objV := reflect.ValueOf(obj)

	if userID, ok := c.Get(global.SessionUserIDKey); !ok {
		return errors.New("require login. ")
	} else {
		if _, ok := objT.FieldByName("UserID"); ok {
			if objV.Kind() == reflect.Ptr {
				objV.Elem().FieldByName("UserID").SetInt(userID.(int64))
			} else {
				objV.FieldByName("UserID").SetInt(userID.(int64))
			}
		}
	}

	if err := global.DB.Model(objV.Interface()).Delete(objV.Interface()).Error; err != nil {
		return err
	}

	return nil
}

func (b *BaseServiceImpl) IUpdateOne(c *gin.Context, obj interface{}, updateObj interface{}) error {
	objT := reflect.TypeOf(obj)
	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
	}

	if err := transform.Transition(updateObj, obj); err != nil {
		return err
	}

	objV := reflect.ValueOf(obj)

	if userID, ok := c.Get(global.SessionUserIDKey); !ok {
		return errors.New("require login. ")
	} else {
		if _, ok := objT.FieldByName("UserID"); ok {
			if objV.Kind() == reflect.Ptr {
				objV.Elem().FieldByName("UserID").SetInt(userID.(int64))
			} else {
				objV.FieldByName("UserID").SetInt(userID.(int64))
			}
		}
	}
	timestamp := time.Now().UnixMilli()
	if _, ok := objT.FieldByName("UpdatedAt"); ok {
		if objV.Kind() == reflect.Ptr {
			objV.Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(&timestamp))
		} else {
			objV.FieldByName("UpdatedAt").Set(reflect.ValueOf(&timestamp))
		}
	}

	if err := global.DB.Model(objV.Interface()).Updates(objV.Interface()).Error; err != nil {
		return err
	}

	return nil
}

func (b *BaseServiceImpl) ISelectOneWeb(c *gin.Context, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
	}

	db := global.DB.Model(objV.Interface())

	if _, ok := objT.FieldByName("Visibility"); ok {
		db.Where("visibility = ?", global.Public)
	}

	if auth.CheckLogin(c) {
		claims, err := token.ParseAccessToken(c.GetHeader(global.RequestHeaderTokenKey))
		if err != nil {
			return err
		}

		if _, ok := objT.FieldByName("UserID"); ok {
			db.Or("user_id=? and visibility=?", claims.UserId, global.Private)
		}
	}

	if err := db.First(objV.Interface()).Error; err != nil {
		return err
	}

	return nil
}

// IUnscopeDelete 直接删除记录
func (b *BaseServiceImpl) IUnscopeDelete(c *gin.Context, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
		objV = objV.Elem()
	}

	if err := global.DB.Model(objV.Interface()).Where("id=?", objV.FieldByName("ID").Int()).Delete(reflect.New(objT).Interface()).Error; err != nil {
		return err
	}
	return nil
}

// ISelectListWeb 查询所有公开的结果
func (b *BaseServiceImpl) ISelectListWeb(c *gin.Context, pager *vo.Pager, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
	}

	objV := reflect.ValueOf(obj)

	var results []map[string]interface{}

	db := global.DB.Debug().Model(objV.Interface())

	if pager.Search != "" && pager.FullIndex {
		// 全文索引
		// 1、判断obj是否有name
		if _, ok := objT.FieldByName("Name"); ok {
			db.Where(fmt.Sprintf("name like '%%%s%%'", pager.Search))
		}

		// 2、判断obj是否有title
		if _, ok := objT.FieldByName("Title"); ok {
			db.Or(fmt.Sprintf("title like '%%%s%%'", pager.Search))
		}
		// 3、判断obj是否有content
		if _, ok := objT.FieldByName("Content"); ok {
			db.Or(fmt.Sprintf("content like '%%%s%%'", pager.Search))
		}
		// 4、判断obj是否有description
		if _, ok := objT.FieldByName("Description"); ok {
			db.Or(fmt.Sprintf("description like '%%%s%%'", pager.Search))
		}

		// 5、判断obj是否有comment
		if _, ok := objT.FieldByName("Comment"); ok {
			db.Or(fmt.Sprintf("comment like '%%%s%%'", pager.Search))
		}

		// 6、判断obj是否有markdown_content
		if _, ok := objT.FieldByName("MarkdownContent"); ok {
			db.Or(fmt.Sprintf("markdown_content like '%%%s%%'", pager.Search))
		}

		// 7、判断obj是否有html_content
		if _, ok := objT.FieldByName("HtmlContent"); ok {
			db.Or(fmt.Sprintf("html_content like '%%%s%%'", pager.Search))
		}
	} else if pager.Search != "" && !pager.FullIndex {
		// 1、判断obj是否有name
		if _, ok := objT.FieldByName("Name"); ok {
			db.Where(fmt.Sprintf("name like '%%%s%%'", pager.Search))
		}

		// 2、判断obj是否有title
		if _, ok := objT.FieldByName("Title"); ok {
			db.Where(fmt.Sprintf("title like '%%%s%%'", pager.Search))
		}
	}

	if _, ok := objT.FieldByName("Visibility"); ok {
		db.Where("visibility = ?", global.Public)
	}

	if auth.CheckLogin(c) {
		claims, err := token.ParseAccessToken(c.GetHeader(global.RequestHeaderTokenKey))
		if err != nil {
			return err
		}

		if _, ok := objT.FieldByName("UserID"); ok {
			db.Or("user_id=? and visibility=?", claims.UserId, global.Private)
		}
	}

	if pager.SortBy == "" || pager.SortOrder == "" {
		pager.MustSort(c)
	}

	if err := db.Count(&pager.TotalRows).Error; err != nil {
		return err
	}

	offset := (pager.PageNo - 1) * pager.PageSize
	limit := pager.PageSize
	if err := db.Order(fmt.Sprintf("%s %s", pager.SortBy, pager.SortOrder)).Offset(offset).Limit(limit).Find(&results).Error; err != nil {
		return err
	}

	pager.PageCount = int((pager.TotalRows + int64(pager.PageSize) - 1) / int64(pager.PageSize))
	pager.MustList(&results)

	return nil
}
