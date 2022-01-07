package web

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Comment struct {
	log         logger.Logger
	postService *service.PostService
	commentService *service.CommentService
}

func NewComment() *Post {
	return &Post{
		log:         global.Logger,
		postService: service.NewPostService(),
		commentService: service.NewCommentService(),
	}
}


// 根据post id 获取comment
func (c *Comment) List(ctx *gin.Context) (*response.Response, error){
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
func (c *Comment) PostComment(ctx *gin.Context) (*response.Response, error) {
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
	return response.Success(&comment),nil
}

// 删除评论
func (comment *Comment) DeleteComment(c *gin.Context) (*response.Response, error) {
	// 1、获取博客ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、判断是否登录

	// 3、删除评论
}