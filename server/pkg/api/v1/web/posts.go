package web

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Post struct {
	log         logger.Logger
	service     common.BaseService
	postService service.PostService
}

func NewPost() *Post {
	return &Post{
		log:         global.Log,
		service:     &common.BaseServiceImpl{},
		postService: service.NewPostService(),
	}
}

// Get 获取单个博客信息
func (p *Post) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	post := po.Post{ID: id}
	if err := p.service.ISelectOneWeb(c, &post); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	// 更新阅读量
	if err := p.postService.IIncrementView(c, &post); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(post), err
}

// List 分页获取所有博客信息
func (p *Post) List(c *gin.Context) (*vo.Response, error) {

	p2 := &po.Post{}

	if err := c.ShouldBindQuery(&p2); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	p2.Status = 2
	if !auth.CheckLogin(c) {
		p2.Visibility = 2
	}

	pages := &vo.Pager{}
	pages.MustPageNo(c)
	pages.MustPageSize(c)
	pages.MustSort(c)
	pages.MustSearch(c)
	pages.FullIndex = true

	if err := p.service.ISelectListWeb(c, pages, p2); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(pages), nil
}

func (p *Post) Like(c *gin.Context) (*vo.Response, error) {
	// 登录之后喜欢
	// 1、判断是否登录
	// 2、获取id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	post := po.Post{ID: id}
	if err := p.postService.ILike(c, &post); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}
	return vo.Message(200, "success"), nil
}
