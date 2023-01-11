package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Post struct {
	log            logger.Logger
	service        common.BaseService
	postTagService service.PostsTagsService
}

func NewPost() *Post {
	return &Post{
		log:            global.Log,
		service:        &common.BaseServiceImpl{},
		postTagService: service.NewPostsTagsService(),
	}
}

func (p *Post) Get(c *gin.Context) (*vo.Response, error) {
	// 1、判断是否有参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		p.log.Infof("获取ID参数错误")
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断是否有权限
	if !auth.CheckPermission(c, "posts", "read") {
		return nil, vo.Forbidden
	}

	// 4、调用service
	post := &po.Post{ID: id}
	if err := p.service.ISelectOne(c, post); err != nil {
		return nil, err
	}
	return vo.Success(post), nil
}

func (p *Post) List(c *gin.Context) (*vo.Response, error) {
	// 1、获取参数
	pager := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}
	pager.MustSearch(c)
	pager.MustSort(c)

	post := &po.Post{}

	if err := c.ShouldBindQuery(&post); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断是否有权限
	if !auth.CheckPermission(c, "posts", "list") {
		return nil, vo.Forbidden.SetMsg("没有该权限，请联系管理员. ")
	}

	// 4、调用service
	if err := p.service.ISelectList(c, &pager, post); err != nil {
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
	post := &po.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 4、调用service
	if err := p.service.ICreateOne(c, post); err != nil {
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	post := &po.Post{ID: id}

	// 4、调用service
	if err := p.service.IDeleteOne(c, post); err != nil {
		return nil, err
	}

	return vo.Success("deleted"), nil
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	post := &po.Post{}
	if err := c.ShouldBind(&post); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	// 4、调用service
	if err := p.service.IUpdateOne(c, &po.Post{ID: id}, post); err != nil {
		return nil, err
	}

	return vo.Success(post), nil
}

func (p *Post) ListTags(c *gin.Context) (*vo.Response, error) {
	// 获取post id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		p.log.Infof("获取ID参数错误")
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 判断是否有权限
	if !auth.CheckPermission(c, "posts", "read") {
		return nil, vo.Forbidden
	}

	tags := make([]*po.Tag, 0)
	if err := p.postTagService.ISelectTagsByPost(c, &tags, &po.Post{ID: id}); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(tags), nil
}
