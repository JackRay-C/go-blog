package v1

import (
	"blog/app/model/po"
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

	// 2、根据用户ID和文件ID查询文件
	file := po.File{ID: id}
	if err = i.fileService.SelectOne(&file); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success(file), nil
}