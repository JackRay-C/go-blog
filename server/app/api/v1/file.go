package v1

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type File struct {
	log         logger.Logger
	fileService *service.FileService
}

func NewFile() *File {
	return &File{
		log:         global.Logger,
		fileService: service.NewFileService(),
	}
}

func (i *File) Get(c *gin.Context) (*response.Response, error) {
	// 1、获取文件ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、获取用户ID
	currentUserId, _ := c.Get("current_user_id")

	// 3、根据用户ID和文件ID查询文件
	file := domain.File{ID: id, UserID: currentUserId.(int)}
	if err = i.fileService.SelectOne(&file); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success(file), nil
}

func (i *File) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("未登录. ")
	}

	if err := i.fileService.SelectAll(c, &p, &domain.File{}); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&p), nil
}

func (i *File) Post(c *gin.Context) (*response.Response, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, response.InvalidParams.SetMsg("文件参数解析错误：%s", err)
	}

	var f domain.File
	if err := i.fileService.CreateOne(c, file, header, &f); err != nil {
		return nil, response.UploadFailed.SetMsg("上传失败: %s", err)
	}

	return response.Success(f), nil
}

func (i *File) Delete(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}


	if err := i.fileService.DeleteOne(c, &domain.File{ID: id}); err != nil {
		return nil, err
	}

	return response.Success("删除成功. "), nil
}


func (i *File) Put(c *gin.Context) (*response.Response, error) {
	panic("implement me")
}
