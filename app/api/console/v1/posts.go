package v1

import (
	"blog/app/model/po"
	"blog/app/model/vo"
	"blog/app/response"
	"blog/app/service"
	"blog/app/utils/auth"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
)

type Post struct {
	log         logger.Logger
	postService service.PostService
}

func NewPost() *Post {
	return &Post{
		log:         global.Logger,
		postService: service.NewPostService(),
	}
}

func (p *Post) Get(c *gin.Context) (*response.Response, error) {
	// 1、判断是否有参数

	// 2、判断是否登录

	// 3、判断是否有权限

	// 4、调用service
	var post *po.Post
	if err := p.postService.ISelectOne(c, post); err != nil {
		return nil, err
	}
	return response.Success(post), nil
}

func (p *Post) List(c *gin.Context) (*response.Response, error) {
	// 1、获取参数

	// 2、判断是否登录

	// 3、判断是否有权限

	// 4、调用service
	pager := vo.Pager{
		PageNo:    0,
		PageSize:  0,
		PageCount: 0,
		TotalRows: 0,
		List:      nil,
	}
	if err := p.postService.ISelectAll(c, &pager); err != nil {
		return nil, err
	}
	return response.Success(pager), nil
}

func (p *Post) Post(c *gin.Context) (*response.Response, error) {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、判断是否有权限
	if !auth.CheckPermission(c, "posts", "add") {
		return nil, vo.Forbidden
	}

	// 3、获取参数
	var post *po.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 4、调用service
	if err := p.postService.ICreateOne(c, post); err != nil {
		return nil, err
	}
	return response.Success(post), nil
}

func (p *Post) Delete(c *gin.Context) (*response.Response, error) {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、判断是否有权限
	if !auth.CheckPermission(c, "posts", "delete") {
		return nil, vo.Forbidden
	}

	// 3、获取参数
	var post *po.Post
	if err := c.ShouldBindQuery(&post); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 4、调用service
	if err := p.postService.IDeleteOne(c, post); err != nil {
		return nil, err
	}

	return response.Success("deleted"), nil
}

func (p *Post) Patch(c *gin.Context) (*response.Response, error) {
	panic("implement me")
}

func (p *Post) Put(c *gin.Context) (*response.Response, error) {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、判断是否有权限
	if !auth.CheckPermission(c, "posts", "update") {
		return nil, vo.Forbidden
	}

	// 3、获取参数
	var post *po.Post

	// 4、调用service
	if err := p.postService.IUpdateOne(c, post); err != nil {
		return nil, err
	}

	return response.Success(post),nil
}
