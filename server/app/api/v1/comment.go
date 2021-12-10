package v1

import (
	"blog/app/domain"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Comment struct {
	log logger.Logger
	commentService *service.CommentService
}

func NewComment() *Comment {
	return &Comment{
		log: global.Logger,
	}
}

func (c *Comment) Get(ctx *gin.Context) (*response.Response, error) {
	c.log.Infof("根据ID查询评论")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	comment := domain.Comment{ID: id}
	if err := c.commentService.SelectOne(&comment); err != nil {
		c.log.Errorf("根据ID查询评论失败： error: %s", err)
		return nil, err
	}
	return response.Success(&comment), nil
}

func (c *Comment) List(ctx *gin.Context) (*response.Response, error) {
	// 获取一个博客下所有评论
	c.log.Infof("获取所有评论")
	// 1、获取分页参数
	page := pager.Pager{
		PageNo:   request.GetPageNo(ctx),
		PageSize: request.GetPageSize(ctx),
	}
	c.log.Infof("分页查询评论: pageNo：%d, pageSize: %d", page.PageNo, page.PageSize)

	if err := c.commentService.SelectAll(&page, &domain.Comment{}); err != nil {
		c.log.Infof("分页查询评论失败. ")
		return nil, err
	}

	// 4、返回查询结果
	c.log.Infof("分页查询评论成功: %s", &page)
	return response.PagerResponse(&page), nil
}

func (c *Comment) Post(ctx *gin.Context) (*response.Response, error) {
	c.log.Infof("新增评论")

	comment := &domain.Comment{}
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		c.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if comment.PostId == 0 {
		return nil, response.InvalidParams.SetMsg("评论博客的ID不能为空. ")
	}

	// 判断是否登录，如果登录，设置userid，否则必须有昵称和邮箱
	if currentUserId, ok := ctx.Get("current_user_id"); ok {
		comment.UserID = currentUserId.(int)
	} else {
		c.log.Infof("%s", comment)
		if comment.Nickname == "" || comment.Email == "" {
			return nil, response.InvalidParams.SetMsg("昵称和邮箱是必须的.")
		}
	}

	if err := c.commentService.CreateOne(comment); err != nil {
		c.log.Errorf("评论失败：error: %s", err)
		return nil, err
	}

	c.log.Infof("评论成功")
	return response.Success(comment), nil
}

func (c *Comment) Delete(ctx *gin.Context) (*response.Response, error) {
	c.log.Infof("删除评论")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	comment := &domain.Comment{ID: id}
	if err := c.commentService.DeleteOne(comment); err != nil {
		return nil, err
	}
	c.log.Infof("删除成功.")
	return response.Success("delete success. "), nil
}

func (c *Comment) Patch(ctx *gin.Context) (*response.Response, error) {
	panic("implement me")
}

func (c *Comment) Put(ctx *gin.Context) (*response.Response, error) {
	panic("implement me")
}



