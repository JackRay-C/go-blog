package impl

import (
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/utils/upload"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mime/multipart"
	"path"
)

type FileServiceImpl struct {
	common.BaseServiceImpl
}

func (f *FileServiceImpl) IDeleteAll(c *gin.Context, files []*po.File) error {
	panic("implement me")
}

//func (f *FileServiceImpl) ISelectOne(c *gin.Context, file *po.File) error {
//	// 查询文件是否存在
//	db := global.DB.Model(&po.File{}).Where("id=?", file.ID)
//
//	if currentUserId, ok := c.Get(common.SessionUserIDKey); ok {
//		db.Where("user_id=?", currentUserId)
//
//
//	if err := db.First(&file).Error; err != nil {
//		return vo.DatabaseSelectError
//	}
//
//	return nil
//}

// ISelectAll 查询所有文件列表，只能登录且查询自己名下的所有文件
//func (f *FileServiceImpl) ISelectAll(c *gin.Context, page *vo.Pager) error {
//	var files []*po.File
//
//	// 1、获取用户ID
//	db := global.DB.Model(&po.File{})
//	if currentUserId, ok := c.Get(common.SessionUserIDKey); ok {
//		db.Where("user_id=?", currentUserId.(int))
//	} else {
//		return vo.NotLogin
//	}
//
//	// 2、统计文件个数
//	if err := db.Count(&page.TotalRows).Error; err != nil {
//		return err
//	}
//
//	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
//	page.List = &files
//
//	// 3、获取文件列表
//	return db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&files).Error
//}

func (f *FileServiceImpl) IUploadFile(c *gin.Context, files multipart.File, header *multipart.FileHeader, file *po.File) error {
	// 1、判断文件大小

	if !upload.CheckMaxSize(files) {
		return errors.New(fmt.Sprintf("file size execeded maximum， limited is：%s", global.App.Storage.AllowUploadMaxSize))
	}

	// 2、判断文件类型
	if !upload.CheckExts(header.Filename) {
		return errors.New("file exts is not support. ")
	}

	ext := path.Ext(header.Filename)
	name, err := global.Storage.Save(header)
	if err != nil {
		return err
	}
	url, err := global.Storage.GetAccessUrl(name)
	if err != nil {
		return err
	}

	currentUserId, _ := c.Get("current_user_id")
	file.UserID = currentUserId.(int)
	file.Ext = ext
	file.Name = name
	file.AccessUrl  = url

	if err := f.ICreateOne(c, &file); err != nil {
		// 数据库插入失败，删除文件
		_, _ = global.Storage.Delete(name)
		global.Log.Errorf("upload file failed: save database error: %s", err)
		return vo.UploadFailed
	}

	return nil
}

func (f *FileServiceImpl) IDeleteFile(c *gin.Context, file *po.File) error {
	// 1、查询文件是否存在
	currentUserId, _ := c.Get("current_user_id")
	err := global.DB.Model(&po.File{}).Where("user_id=? and id=?", currentUserId.(int), file.ID).First(&file).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(fmt.Sprintf("this file id: %d is not found", file.ID))
	}

	// 2、删除存储
	if _, err := global.Storage.Delete(file.Name); err != nil {
		return err
	}

	// 3、删除数据库数据
	if err := global.DB.Model(&po.File{}).Where("user_id=? and id=?",currentUserId.(int), file.ID).Delete(&file); err != nil {
		return errors.New(fmt.Sprintf("删除数据库记录失败： %s", err))
	}

	return nil
}
