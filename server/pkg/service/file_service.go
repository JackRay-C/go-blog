package service

import (
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/service/impl"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type FileService interface {
	common.BaseService

	IUploadFile(c *gin.Context, file multipart.File, header *multipart.FileHeader, f *po.File) error
	IDeleteFile(c *gin.Context, f *po.File) error
}

func NewFileService() FileService {
	return &impl.FileServiceImpl{}
}

