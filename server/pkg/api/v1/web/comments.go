package web

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Comment struct {
	log            logger.Logger
	service        common.BaseService

}

func NewComment() *Comment {
	return &Comment{
		log:            global.Log,
		service:        &common.BaseServiceImpl{},
	}
}

// Get 获取一条评论
func (c *Comment) Get(ctx *gin.Context) (*vo.Response, error) {

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	comment := &po.Comment{ID: id}
	if err := c.service.ISelectOneWeb(ctx, comment); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(comment), nil
}

// List 根据博客ID获取所有评论
func (c *Comment) List(ctx *gin.Context) (*vo.Response, error) {
	// 查询博客的评论
	pager := &vo.Pager{PageNo: page.GetPageNo(ctx), PageSize: page.GetPageSize(ctx)}


	var comments *po.Comment
	if err := ctx.ShouldBind(&comments); err != nil {
		return nil, vo.InvalidParams
	}

	if err := c.service.ISelectList(ctx, pager, &comments); err != nil {
		return nil, vo.InternalServerError
	}

	return vo.Success(pager), nil
}
