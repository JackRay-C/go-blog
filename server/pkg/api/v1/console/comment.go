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

type Comment struct {
	log            logger.Logger
	commentService service.CommentService
}

func NewComment() *Comment {
	return &Comment{
		log:            global.Log,
		commentService: service.NewCommentService(),
	}
}

func (c *Comment) Get(ctx *gin.Context) (*vo.Response, error) {
	// 1、判断是否登录
	if !auth.CheckPermission(ctx, "comments", "read") {
		return nil, vo.Forbidden
	}

	// 2、根据user id 查询comment
	userId, _ := ctx.Get(global.SessionUserIDKey)

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("invalid params ID.")
	}

	c2 := &po.Comment{ID: id, UserID: userId.(int64)}

	// 3、查询comment
	if err := c.commentService.ISelectOne(ctx, c2); err != nil {
		c.log.Errorf("select comment failed service error: %s", err)
		return nil, vo.InternalServerError
	}

	return vo.Success(c2), nil
}

func (c *Comment) List(ctx *gin.Context) (*vo.Response, error) {
	if !auth.CheckPermission(ctx, "comments", "list") {
		return nil, vo.Forbidden
	}

	p := vo.Pager{
		PageNo:   page.GetPageNo(ctx),
		PageSize: page.GetPageSize(ctx),
	}

	// 3、分页查询评论
	if err := c.commentService.ISelectList(ctx, &p, &po.Comment{}); err != nil {
		return nil, err
	}

	return vo.Success(&p), nil
}

func (c *Comment) Post(ctx *gin.Context) (*vo.Response, error) {
	return vo.Failed(500, "console not be allow post comment. "), nil
}

func (c *Comment) Delete(ctx *gin.Context) (*vo.Response, error) {
	if !auth.CheckPermission(ctx, "comments", "delete") {
		return nil, vo.Forbidden
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams
	}

	// 2、根据user id 查询comment
	comment := &po.Comment{ID: id, UserID: auth.GetCurrentUserId(ctx)}

	if err := c.commentService.IDeleteOne(ctx, comment); err != nil {
		c.log.Error(err)
		return nil, err
	}

	return vo.Success("success"), nil
}
