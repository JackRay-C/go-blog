package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)
type File struct {
	log         logger.Logger
	fileService service.FileService
}

func NewFile() *File {
	return &File{
		log:         global.Log,
		fileService: service.NewFileService(),
	}
}

func (i *File) Get(c *gin.Context) (*vo.Response, error) {
	// 1、获取文件ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、获取用户ID
	currentUserId, _ := c.Get("current_user_id")

	// 3、根据用户ID和文件ID查询文件
	file := po.File{ID: id, UserID: currentUserId.(int)}
	if err = i.fileService.ISelectOne(c, &file); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}
	return vo.Success(file), nil
}

func (i *File) List(c *gin.Context) (*vo.Response, error) {
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin.SetMsg("未登录. ")
	}

	if err := i.fileService.ISelectList(c, &p, &po.File{}); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&p), nil
}

func (i *File) Post(c *gin.Context) (*vo.Response, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("文件参数解析错误：%s", err)
	}

	var f po.File
	if err := i.fileService.IUploadFile(c, file, header, &f); err != nil {
		return nil, vo.UploadFailed.SetMsg("上传失败: %s", err)
	}

	return vo.Success(f), nil
}

func (i *File) Delete(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}


	if err := i.fileService.IDeleteOne(c, &po.File{ID: id}); err != nil {
		return nil, err
	}

	return vo.Success("删除成功. "), nil
}

