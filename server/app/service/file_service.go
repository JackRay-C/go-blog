package service

import (
	"blog/app/domain"
	"blog/app/encrypt"
	"blog/app/pager"
	"blog/app/response"
	"blog/core/global"
	"blog/core/logger"
	"blog/core/storage"
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
	return &FileService{log: global.Logger}
}

func (service *FileService) SelectOne(file *domain.File) error {
	service.log.Infof("查询文件")
	if err := file.Select(); err != nil {
		service.log.Errorf("查询文件错误: %s", err)
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	service.log.Infof("查询成功. ")
	return nil
}

func (service *FileService) SelectAll(page *pager.Pager, file *domain.File) error {
	service.log.Infof("分页查询文件列表")
	var files []domain.File

	if err:= file.Count(&page.TotalRows); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	page.PageCount = int((page.TotalRows + int64(page.PageSize) - 1) / int64(page.PageSize))
	page.List = &files

	if err := file.List(&files, (page.PageNo-1)*page.PageSize, page.PageSize); err != nil {
		service.log.Errorf("分页查询文件列表失败。")
		return err
	}
	service.log.Infof("分页查询文件成功. ")
	return nil
}

func (service *FileService) DeleteOne(file *domain.File) error {
	// 1、查询文件
	service.log.Infof("查询待删除文件 ")
	if err := file.Select(); err != nil {
		service.log.Infof("查询待删除文件失败: %s", err)
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	// 2、删除远程文件
	service.log.Infof("删除存储文件")
	if _, err := global.Storage.Delete(file.Path + "/" + file.Name); err != nil {
		service.log.Errorf("删除存储文件失败： %s", err)
		return response.FailedRemoveFile.SetMsg("%s", err)
	}

	// 3、删除数据库数据
	service.log.Infof("删除数据库记录")
	if err := file.Delete(); err != nil {
		service.log.Errorf("删除数据库记录失败: %s", err)
		return response.DatabaseDeleteError.SetMsg("%s", err)
	}
	service.log.Infof("删除文件成功. ")
	return nil
}

func (service *FileService) DeleteAll(ids []int) error {
	service.log.Infof("查询待删除文件 ")
	var filesNames []string
	for _, id := range ids {
		f := domain.File{ID: id}
		if err := f.Select(); err != nil {
			service.log.Errorf("查询ID为 %d 的文件失败： %s", id, err)
			//return response.DatabaseSelectError.SetMsg("%s", err)
			continue
		}
		filesNames = append(filesNames, f.Path+"/"+f.Name)
	}

	service.log.Infof("删除存储文件: %s", filesNames)
	if _, err := global.Storage.Delete(filesNames...); err != nil {
		return response.FailedRemoveFile.SetMsg("%s", err)
	}
	service.log.Infof("删除数据库记录: %v", ids)
	file := &domain.File{}
	if err := file.DeleteIds(ids); err != nil {
		service.log.Errorf("删除数据库记录失败： %s", err)
		return response.DatabaseDeleteError.SetMsg("%s", err)
	}

	service.log.Infof("删除数据库记录成功")
	return nil
}

func (service *FileService) CreateOne(c *gin.Context, file multipart.File, header *multipart.FileHeader, f *domain.File) error {
	// 1、判断文件大小
	content, _ := ioutil.ReadAll(file)
	if int64(len(content)) >= int64(global.Setting.App.UploadMaxSize) {
		return response.ExceededMaximumLimit.SetMsg("超出 %s ", global.Setting.App.UploadMaxSize)
	}

	// 2、根据文件后缀判断文件类型
	ext := path.Ext(header.Filename)
	var flag bool
	for _, e := range global.Setting.App.UploadAllowExts {
		if strings.ToUpper(e) == strings.ToUpper(ext) {
			flag = true
			break
		}
	}
	if !flag {
		return response.NotSupportedSuffix.SetMsg("不支持该文件类型！")
	}

	// 3、先从数据库查询文件，判断是否存在
	f.Ext = ext
	f.Name = encrypt.MD5(strings.TrimSuffix(header.Filename, path.Ext(header.Filename))) + ext
	f.Type = storage.GetFileTypeByExt(ext)
	f.Path = storage.GetFilePrefix(f.Type)
	f.AccessUrl  = global.Storage.GetAccessUrl(f.Path, f.Name)

	switch {
	case strings.ToUpper(global.Setting.App.StorageType) == "LOCAL":
		f.Host = c.Request.Host
	case strings.ToUpper(global.Setting.App.StorageType) == "QINIU":
		f.Host = global.Setting.Qiniu.ImgPath + "/" + global.Setting.Qiniu.Bucket
	case strings.ToUpper(global.Setting.App.StorageType) == "ALIYUN-OSS":
		f.Host = global.Setting.AliyunOSS.BucketUrl + "/" + global.Setting.AliyunOSS.BasePath
	case strings.ToUpper(global.Setting.App.StorageType) == "TENCENT-OSS":
		f.Host = global.Setting.TencentOSS.BucketUrl
	default:
		f.Host = c.Request.Host
	}
	if err := f.Select(); err == gorm.ErrRecordNotFound {
		// 3.1 如果没有找到，则插入数据库
		if err := f.Insert(); err != nil {
			service.log.Errorf("添加文件记录失败！")
			return response.DatabaseInsertError.SetMsg("%s", err)
		}
		// 3.2 上传至oss
		if err := global.Storage.Save(path.Join(f.Path, f.Name), header); err != nil {
			service.log.Errorf("创建存储文件失败: %s", err)
			if err := f.Delete(); err != nil {
				return response.DatabaseDeleteError.SetMsg("回滚删除数据库记录失败： %s， 文件ID： ", err, f.ID)
			}
			return response.UploadFailed.SetMsg("%s", err)
		}
	} else if err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	service.log.Infof("创建文件成功. %s", f)
	return nil
}

func (service *FileService) CreateAll() error {
	return response.InternalServerError.SetMsg("暂不支持批量上传文件.")
}
