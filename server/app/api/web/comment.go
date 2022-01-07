package web

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Comment struct {
	log            logger.Logger
	postService    *service.PostService
	commentService *service.CommentService
}

func NewComment() *Comment {
	return &Comment{
		log:            global.Logger,
		postService:    service.NewPostService(),
		commentService: service.NewCommentService(),
	}
}

// 根据post id 获取comment
func (c *Comment) List(ctx *gin.Context) (*response.Response, error) {
	// 1、获取博客ID
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、根据博客ID查询所有的评论
	var comments []*domain.Comment
	if err := c.commentService.SelectPostComments(&domain.Post{ID: id}, &comments); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&comments), nil
}

// 提交评论
func (c *Comment) Post(ctx *gin.Context) (*response.Response, error) {
	// 1、获取博客ID
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、绑定评论信息
	var comment domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 3、判断是否登录，未登录需要填写邮箱和昵称
	if !api.CheckLogin(ctx) {
		if comment.Nickname == "" || comment.Email == "" {
			return nil, response.InvalidParams.SetMsg("评论昵称或邮箱不能为空. ")
		}
	} else {
		// 4、如果登录，获取当前用户的信息
		currentUserId, _ := ctx.Get("current_user_id")
		comment.UserID = currentUserId.(int)
	}

	if err := c.commentService.CreateOne(&comment); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success(&comment), nil
}

// 删除评论
func (comment *Comment) Delete(c *gin.Context) (*response.Response, error) {
	// 1、获取评论ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	cid, err := strconv.Atoi(c.Param("cid"))
	if err != nil || cid == 0 {
		return nil, response.InvalidParams.SetMsg("Comment ID is required. ")
	}

	// 2、判断是否登录
	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("删除失败，未登录. ")
	}

	// 3、获取用户ID
	currentUserId, _ := c.Get("current_user_id")

	// 4、删除评论
	if err := comment.commentService.DeleteOne(&domain.Comment{ID: cid, UserID: currentUserId.(int), PostId: id}); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success("删除成功."), nil
}
