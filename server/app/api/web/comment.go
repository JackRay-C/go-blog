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
	userService *service.UserService
}

func NewComment() *Comment {
	return &Comment{
		log:            global.Logger,
		postService:    service.NewPostService(),
		commentService: service.NewCommentService(),
		userService: service.NewUserService(),
	}
}

// Get 根据ID查询评论信息
func (c *Comment) Get(ctx *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	comment := domain.Comment{ID: id}
	if err := c.commentService.SelectOne(&comment); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success(&comment), nil
}

// List 根据post id 获取comment
func (c *Comment) List(ctx *gin.Context) (*response.Response, error) {
	// 1、绑定参数
	post_id, err := strconv.Atoi(ctx.Query("post_id"))
	if err != nil || post_id == 0 {
		return nil, response.InvalidParams.SetMsg("post_id is required. ")
	}

	// 2、根据博客ID查询所有的评论
	comments := make([]*domain.Comment, 0)
	if err := c.commentService.SelectPostComments(&domain.Post{ID: post_id}, &comments); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&comments), nil
}

// Post 提交评论
func (c *Comment) Post(ctx *gin.Context) (*response.Response, error) {
	// 1、绑定评论信息
	var comment domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断postid是否为空
	if comment.PostId == 0 {
		return nil, response.InvalidParams.SetMsg("评论post_id不能为空. ")
	}

	// 3、判断是否登录，未登录需要填写邮箱和昵称
	if !api.CheckLogin(ctx) {
		// 3.1 未登录，需要填写昵称和邮箱，方便后续发送邮件回复
		if comment.Nickname == "" || comment.Email == "" {
			return nil, response.InvalidParams.SetMsg("评论昵称或邮箱不能为空. ")
		}
	} else {
		// 3.2、如果登录，获取当前用户的信息
		currentUserId, _ := ctx.Get("current_user_id")
		comment.UserID = currentUserId.(int)

		if one, err := c.userService.SelectOne(&domain.User{ID: currentUserId.(int)}); err != nil {
			return nil, response.InternalServerError.SetMsg("查询用户失败. ")
		} else {
			comment.Nickname = one.Nickname
			comment.Email = one.Email
		}
	}

	// 4、添加评论
	if err := c.commentService.CreateOne(&comment); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&comment), nil
}

// Delete 删除评论
func (comment *Comment) Delete(c *gin.Context) (*response.Response, error) {
	// 1、获取评论ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、判断是否登录
	if !api.CheckLogin(c) {
		return nil, response.NotLogin.SetMsg("未登录. ")
	}

	// 3、获取用户ID
	currentUserId, _ := c.Get("current_user_id")

	// 4、删除评论
	if err := comment.commentService.DeleteOne(&domain.Comment{ID: id, UserID: currentUserId.(int)}); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success("删除成功."), nil
}
