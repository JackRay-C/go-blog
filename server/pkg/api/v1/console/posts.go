package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/bo"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Post struct {
	log         logger.Logger
	postService service.PostService
}

func NewPost() *Post {
	return &Post{
		log:         global.Log,
		postService: service.NewPostService(),
	}
}

func (p *Post) Get(c *gin.Context) (*vo.Response, error) {
	// 1、判断是否有参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.log.Infof("获取ID参数错误")
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 3、判断是否有权限
	if !auth.CheckPermission(c, "posts", "read") {
		return nil, vo.Forbidden
	}

	// 4、调用service
	post := &bo.Post{
		Head: &po.Head{
			ID: id,
		},
		Repositories: []*po.Repository{},
		Histories: []*po.History{},
	}
	if err := p.postService.ISelectOne(c, post); err != nil {
		return nil, err
	}
	return vo.Success(post), nil
}

func (p *Post) List(c *gin.Context) (*vo.Response, error) {
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
	if err := p.postService.ISelectList(c, &pager, &bo.Post{
		Head: &po.Head{ID: 1},
	}); err != nil {
		return nil, err
	}
	return vo.Success(pager), nil
}

func (p *Post) Post(c *gin.Context) (*vo.Response, error) {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、判断是否有权限
	if !auth.CheckPermission(c, "posts", "add") {
		return nil, vo.Forbidden
	}

	// 3、获取参数
	var post *bo.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 4、调用service
	if err := p.postService.ICreateOne(c, post); err != nil {
		return nil, err
	}
	return vo.Success(post), nil
}

func (p *Post) Delete(c *gin.Context) (*vo.Response, error) {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、判断是否有权限
	if !auth.CheckPermission(c, "posts", "delete") {
		return nil, vo.Forbidden
	}

	// 3、获取参数
	var post *bo.Post
	if err := c.ShouldBindQuery(&post); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 4、调用service
	if err := p.postService.IDeleteOne(c, post); err != nil {
		return nil, err
	}

	return vo.Success("deleted"), nil
}

func (p *Post) Patch(c *gin.Context) (*vo.Response, error) {
	panic("implement me")
}

func (p *Post) Put(c *gin.Context) (*vo.Response, error) {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、判断是否有权限
	if !auth.CheckPermission(c, "posts", "update") {
		return nil, vo.Forbidden
	}

	// 3、获取参数
	var post *bo.Post

	// 4、调用service
	if err := p.postService.IUpdateOne(c, post, post); err != nil {
		return nil, err
	}

	return vo.Success(post), nil
}
