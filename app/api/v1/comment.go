package v1

import (
	"blog/app/model/po"
	"blog/app/model/vo"
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
	commentService service.CommentService
}

func NewComment() *Comment {
	return &Comment{
		log: global.Logger,
		commentService: service.NewCommentService(),
	}
}

// Get 根据ID查询评论信息
func (c *Comment) Get(ctx *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	comment := po.Comment{ID: id}
	if err := c.commentService.ISelectOne(ctx, &comment); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success(&comment), nil
}

// List 分页所有评论信息
func (c *Comment) List(ctx *gin.Context) (*response.Response, error) {
	// 1、获取分页参数
	page := vo.Pager{
		PageNo:   request.GetPageNo(ctx),
		PageSize: request.GetPageSize(ctx),
	}

	// 2、查询所有博客
	if err := c.commentService.ISelectAll(ctx, &page, &po.Comment{}); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	// 3、返回查询结果
	return response.Success(&page), nil
}

func (c *Comment) Post(ctx *gin.Context) (*response.Response, error) {
	c.log.Infof("新增评论")

	comment := &po.Comment{}
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

	if err := c.commentService.ICreate(ctx, comment); err != nil {
		c.log.Errorf("评论失败：error: %s", err)
		return nil, err
	}

	c.log.Infof("评论成功")
	return response.Success(comment), nil
}

// Delete 根据ID删除单条评论
func (c *Comment) Delete(ctx *gin.Context) (*response.Response, error) {
	c.log.Infof("删除评论")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	comment := &po.Comment{ID: id}
	if err := c.commentService.IDelete(ctx, comment); err != nil {
		return nil, err
	}
	c.log.Infof("删除成功.")
	return response.Success("delete success. "), nil
}




