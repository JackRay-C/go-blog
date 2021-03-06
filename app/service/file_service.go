package service

import (
	"blog/app/encrypt"
	"blog/app/pager"
	"blog/internal/logger"
	"blog/internal/storage"
	"blog/pkg/global"
	"blog/pkg/model/po"

	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strings"
)

type FileService struct {
	log logger.Logger
}

func NewFileService() *FileService {
	return &FileService{log: global.Log}
}

func (service *FileService) SelectOne(file *po.File) error {
	// 查询文件是否存在
	db := global.DB.Model(&po.File{}).Where("id=?", file.ID)

	if file.UserID != 0 {
		db.Where("user_id=?", file.UserID)
	}

	err := db.First(&file).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该文件不存在. ")
	}

	// 3、返回查询结果
	return err
}

func (service *FileService) SelectAll(c *gin.Context, page *pager.Pager, file *po.File) error {
	var files []*po.File

	// 1、获取用户ID
	currentUserId, _ := c.Get("current_user_id")

	// 2、统计文件个数
	db := global.DB.Model(&po.File{}).Where("user_id=?", currentUserId.(int))
	if err := db.Count(&page.TotalRows).Error; err != nil {
		return err
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &files

	// 3、获取文件列表
	return db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&files).Error
}

func (service *FileService) DeleteOne(c *gin.Context, file *po.File) error {
	// 1、查询文件是否存在
	currentUserId, _ := c.Get("current_user_id")
	err := global.DB.Model(&po.File{}).Where("user_id=? and id=?", currentUserId.(int), file.ID).First(&file).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		service.log.Errorf("删除文件[%d]失败：文件不存在. ", file.ID)
		return errors.New("该文件不存在. ")
	}

	// 2、删除存储
	if _, err := global.Storage.Delete(file.Path + "/" + file.Name); err != nil {
		service.log.Errorf("删除存储文件失败： %s", err)
		return errors.New(fmt.Sprintf("删除存储文件失败： %s", err))
	}

	// 3、删除数据库数据
	if err := global.DB.Model(&po.File{}).Where("user_id=? and id=?",currentUserId.(int), file.ID).Delete(&file); err != nil {
		service.log.Errorf("删除数据库记录失败: %s", err)
		return errors.New(fmt.Sprintf("删除数据库记录失败： %s", err))
	}

	return nil
}


func (service *FileService) CreateOne(c *gin.Context, file multipart.File, header *multipart.FileHeader, f *po.File) error {
	// 1、判断文件大小
	content, _ := ioutil.ReadAll(file)

	if int64(len(content)) >= int64(global.App.Storage.AllowUploadMaxSize) {
		return errors.New(fmt.Sprintf("文件大小： %d，超出文件最大限制：%s", int64(len(content)), global.App.Storage.AllowUploadMaxSize))
	}

	// 2、根据文件后缀判断文件类型
	ext := path.Ext(header.Filename)
	var flag bool
	for _, e := range global.App.Storage.AllowUploadFileExts {
		if strings.ToUpper(e) == strings.ToUpper(ext) {
			flag = true
			break
		}
	}
	if !flag {
		return errors.New(fmt.Sprintf("不支持该文件类型：%s !", ext))
	}

	// 3、先从数据库查询文件，判断是否存在
	currentUserId, _ := c.Get("current_user_id")
	f.UserID = currentUserId.(int)
	f.Ext = ext
	f.Name = encrypt.MD5(strings.TrimSuffix(header.Filename, path.Ext(header.Filename))) + ext
	f.Type = storage.GetFileTypeByExt(ext)
	f.Path = storage.GetFilePrefix(f.Type)
	f.AccessUrl, _ = global.Storage.GetAccessUrl(f.Path, f.Name)

	switch {
	case strings.ToUpper(global.App.AppStorageType) == "LOCAL":
		f.Host = c.Request.Host
	case strings.ToUpper(global.App.AppStorageType) == "QINIU":
		f.Host = global.App.Storage.Qiniu.ImgPath + "/" + global.App.Storage.Qiniu.Bucket
	case strings.ToUpper(global.App.AppStorageType) == "ALIYUN-OSS":
		f.Host = global.App.Storage.AliyunOSS.BucketUrl + "/" + global.App.Storage.AliyunOSS.BasePath
	case strings.ToUpper(global.App.AppStorageType) == "TENCENT-OSS":
		f.Host = global.App.Storage.TencentOSS.BucketUrl
	default:
		f.Host = c.Request.Host
	}

	var newFile *po.File
	err := global.DB.Model(&po.File{}).Where("user_id=? and name=? and ext=?", f.UserID, f.Name, f.Ext).First(&newFile).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 插入数据库
		if err := global.DB.Model(&po.File{}).Create(&f).Error; err != nil {
			return err
		}

		// 上传至存储
		if err := global.Storage.Save(path.Join(f.Path, f.Name), header); err != nil {
			// 上传失败补偿，删除数据库记录
			if err := global.DB.Model(&po.File{}).Where("id=?", f.ID).Delete(&f).Error; err != nil {
				return err
			}
			return errors.New("上传存储失败，请重新上传. ")
		}

	}
	if err != nil {
		return err
	}

	return errors.New("该文件已存在. ")
}