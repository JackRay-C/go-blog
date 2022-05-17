package service

import (
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/service/impl"
	"github.com/gin-gonic/gin"
)

type CommentService interface {
	common.BaseService

	// ISelectAllByPostId 查询一条博客的所有评论信息
	ISelectAllByPostId(ctx *gin.Context, comments *[]*po.Comment) error
}


func NewCommentService() CommentService {
	return &impl.CommentServiceImpl{}
}
