package common

import (
	"blog/pkg/global"
	"blog/pkg/model/vo"
	"blog/pkg/utils/transform"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
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

	if userID, ok := c.Get(global.SessionUserIDKey); !ok {
		return errors.New("require login. ")
	} else {
		if _, ok := objT.FieldByName("UserID"); ok {
			db.Where("user_id=?", userID)
		}
	}

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

	db := global.DB.Model(objV.Interface()).Where(objV.Interface())

	if userID, ok := c.Get(global.SessionUserIDKey); !ok {
		return errors.New("require login. ")
	} else {
		if _, ok := objT.FieldByName("UserID"); ok {
			db.Where("user_id=?", userID)
		}
	}

	if pager.Search != "" {
		// 1、判断obj是否有name
		if _, ok := objT.FieldByName("Name"); ok {
			db.Where(fmt.Sprintf("name like %%%s%%", pager.Search))
		}

		// 2、判断obj是否有title
		if _, ok := objT.FieldByName("Title"); ok {
			db.Where(fmt.Sprintf("title like %%%s%%", pager.Search))
		}
		// 3、判断obj是否有content
		if _, ok := objT.FieldByName("Content"); ok {
			db.Where(fmt.Sprintf("content like %%%s%%", pager.Search))
		}
		// 4、判断obj是否有description
		if _, ok := objT.FieldByName("Description"); ok {
			db.Where(fmt.Sprintf("description like %%%s%%", pager.Search))
		}

		// 5、判断obj是否有comment
		if _, ok := objT.FieldByName("Comment"); ok {
			db.Where(fmt.Sprintf("comment like \"%%%s%%\"", pager.Search))
		}
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
			objV.FieldByName("UserID").SetInt(userID.(int64))
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
			objV.FieldByName("UserID").SetInt(userID.(int64))
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
			objV.FieldByName("UserID").SetInt(userID.(int64))
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

	if userID, ok := c.Get(global.SessionUserIDKey); ok {
		// 登录
		if _, ok := objT.FieldByName("UserID"); ok {
			db.Or("user_id=? and visibility=?", userID, global.Private)
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
	}

	if err := global.DB.Model(objV.Interface()).Unscoped().Update("delete_at", nil).Error; err != nil {
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

	db := global.DB.Model(objV.Interface())

	if pager.Search != "" {
		// 1、判断obj是否有name
		if _, ok := objT.FieldByName("Name"); ok {
			db.Where(fmt.Sprintf("name like %%%s%%", pager.Search))
		}

		// 2、判断obj是否有title
		if _, ok := objT.FieldByName("Title"); ok {
			db.Where(fmt.Sprintf("title like %%%s%%", pager.Search))
		}
		// 3、判断obj是否有content
		if _, ok := objT.FieldByName("Content"); ok {
			db.Where(fmt.Sprintf("content like %%%s%%", pager.Search))
		}
		// 4、判断obj是否有description
		if _, ok := objT.FieldByName("Description"); ok {
			db.Where(fmt.Sprintf("description like %%%s%%", pager.Search))
		}

		// 5、判断obj是否有comment
		if _, ok := objT.FieldByName("Comment"); ok {
			db.Where(fmt.Sprintf("comment like \"%%%s%%\"", pager.Search))
		}
	}

	if _, ok := objT.FieldByName("Visibility"); ok {
		db.Where("visibility = ?", global.Public)
	}

	if userID, ok := c.Get(global.SessionUserIDKey); ok {
		// 登录
		if _, ok := objT.FieldByName("UserID"); ok {
			db.Or("user_id=? and visibility=?", userID, global.Private)
		}
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
