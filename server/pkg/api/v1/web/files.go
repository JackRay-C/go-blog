package web

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
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

	// 2、根据用户ID和文件ID查询文件
	file := po.File{ID: id}
	if err = i.fileService.ISelectOne(c,&file); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}
	return vo.Success(file), nil
}