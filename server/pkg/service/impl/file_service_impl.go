package impl

import (
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/encrypt"
	"blog/pkg/utils/upload"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
	"path"
	"reflect"
)

type FileServiceImpl struct {
	common.BaseServiceImpl
}

func (f *FileServiceImpl) IDeleteAll(c *gin.Context, files []*po.File) error {
	panic("implement me")
}

// ISelectOne 重写接口，获取access url
func (f *FileServiceImpl) ISelectOne(c *gin.Context, obj interface{}) error  {
	objT := reflect.TypeOf(obj)

	if objT.String() != "*po.File" {
		return errors.New("obj is not *po.File. ")
	}
	file := obj.(*po.File)
	db := global.DB.Model(&po.File{})
	if err := db.Where("id=? and user_id=?", file.ID, auth.GetCurrentUserId(c)).First(file).Error; err != nil {
		return err
	}

	if url, err := global.Storage.GetAccessUrl(file.Name); err != nil {
		return err
	} else {
		log.Println(url)
		file.AccessUrl = url
	}

	return nil
}

func (f *FileServiceImpl) IUploadFile(c *gin.Context, files multipart.File, header *multipart.FileHeader, file *po.File) error {
	// 1、判断文件大小
	if !upload.CheckMaxSize(files) {
		return errors.New(fmt.Sprintf("file size execeded maximum， limited is：%s", global.App.Storage.AllowUploadMaxSize))
	}

	// 2、判断文件类型
	if !upload.CheckExts(header.Filename) {
		return errors.New("file exts is not support. ")
	}

	// 3、保存到oss
	name, err := global.Storage.Save(header)
	if err != nil {
		return err
	}
	// 4、获取access url
	url, err := global.Storage.GetAccessUrl(name)
	if err != nil {
		return errors.New(fmt.Sprintf("get access url failed: %s", err))
	}

	// 5、对文件内容进行md5
	content, _ := ioutil.ReadAll(files)
	file.Md5 = encrypt.MD5(string(content))

	// 6、设置数据库其他属性
	file.UserID = auth.GetCurrentUserId(c)
	file.Ext = path.Ext(header.Filename)
	file.Name = name
	file.AccessUrl = url

	// 插入数据库
	if err := f.ICreateOne(c, file); err != nil {
		// 数据库插入失败，删除文件
		go func() {
			_, e := global.Storage.Delete(name)
			if e != nil {
				global.Log.Errorf("delete oss file failed: %s. ", e)
			}
		}()
		return errors.New(fmt.Sprintf("upload file failed: save database error: %s", err))
	}

	return nil
}

func (f *FileServiceImpl) IDeleteFile(c *gin.Context, file *po.File) error {
	// 1、查询文件是否存在
	var count int64
	if err := global.DB.Model(&po.File{}).Where("user_id=? and id=?", auth.GetCurrentUserId(c), file.ID).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return errors.New(fmt.Sprintf("this file: %d is not found", file.ID))
	}

	if err := global.DB.Model(&po.File{}).Where("user_id=? and id=?", auth.GetCurrentUserId(c), file.ID).First(&file).Error; err != nil {
		return err
	}

	// 2、删除存储
	if _, err := global.Storage.Delete(file.Name); err != nil {
		return err
	}

	// 3、删除数据库数据
	if err := global.DB.Model(&po.File{}).Where("user_id=? and id=?",auth.GetCurrentUserId(c), file.ID).Delete(&file).Error; err != nil {
		return err
	}

	return nil
}
