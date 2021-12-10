package v1

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
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
	i.log.Infof("根据ID查看文件")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	file := domain.File{ID: id}
	if err = i.fileService.SelectOne(&file); err != nil {
		i.log.Errorf("根据ID查看文件失败： error: %s", err)
		return nil, err
	}
	return response.Success(file), nil
}

func (i *File) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	i.log.Infof("分页查询文件列表")
	if err := i.fileService.SelectAll(&p, &domain.File{}); err != nil {
		return nil, err
	}

	return response.PagerResponse(&p), nil
}

func (i *File) Post(c *gin.Context) (*response.Response, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		i.log.Errorf("参数错误：%s", err)
		return nil, response.InvalidParams.SetMsg("文件参数解析错误：%s", err)
	}

	i.log.Infof("上传文件：%s", header.Filename)
	var f domain.File
	if currentUserId, ok := c.Get("current_user_id"); ok {
		f.UserID = currentUserId.(int)
	}

	if err := i.fileService.CreateOne(c, file, header, &f); err != nil {
		i.log.Errorf("上传失败: %s", err)
		return nil, response.UploadFailed.SetMsg("上传失败: %s", err)
	}

	i.log.Infof("上传成功: %s", &f)
	return response.Success(f), nil
}

func (i *File) Delete(c *gin.Context) (*response.Response, error) {
	var ids []int
	query := strings.Split(c.Query("ids"), ",")

	for _, id := range query {
		id, _ := strconv.Atoi(id)
		ids = append(ids, id)
	}

	if len(ids) <= 0 {
		return nil, response.InvalidParams.SetMsg("ids is required. ")
	}

	if err := i.fileService.DeleteAll(ids); err != nil {
		return nil, err
	}

	return response.Success("删除成功. "), nil
}

func (i *File) Patch(c *gin.Context) (*response.Response, error) {
	// 修改文件名
	panic("implement me")
}

func (i *File) Put(c *gin.Context) (*response.Response, error) {
	panic("implement me")
}
