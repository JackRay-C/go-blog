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

type Comment struct {
	log            logger.Logger
	postService    service.PostService
	commentService service.CommentService
	userService service.UserService
}

func NewComment() *Comment {
	return &Comment{
		log:            global.Log,
		postService:    service.NewPostService(),
		commentService: service.NewCommentService(),
		userService: service.NewUserService(),
	}
}


// Get 获取一条评论
func (c *Comment) Get(ctx *gin.Context) (*vo.Response, error) {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	comment := &po.Comment{ID: id}
	if err := c.commentService.ISelectOne(ctx, comment); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(comment), nil
}

// List 根据博客ID获取所有评论
func (c *Comment) List(ctx *gin.Context) (*vo.Response, error) {
	// 查询博客的评论
	var comments []*po.Comment
	if err := c.commentService.ISelectAllByPostId(ctx, &comments); err != nil {
		return nil, err
	}

	return vo.Success(comments), nil
}
