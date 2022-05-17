package web

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/bo"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Post struct {
	log            logger.Logger
	postService    service.PostService
	commentService service.CommentService
}

func NewPost() *Post {
	return &Post{
		log:            global.Log,
		postService:    service.NewPostService(),
		commentService: service.NewCommentService(),
	}
}

// Get 获取单个博客信息
func (p *Post) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	post := bo.Post{Head: &po.Head{ID: id}}
	if err := p.postService.ISelectOneWeb(c, &post); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(post), err
}

// List 分页获取所有博客信息
func (p *Post) List(c *gin.Context) (*vo.Response, error) {

	p2 := &bo.Post{
		Head: &po.Head{
			Visibility: global.Public,
			Status:     global.Publish,
			SubjectID:  0,
			UserID: 0,
		},
		Repositories: []*po.Repository{},
		Histories:    []*po.History{},
	}

	if err := c.ShouldBindQuery(&p2); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	pages := &vo.Pager{}
	pages.MustPageNo(c)
	pages.MustPageSize(c)
	pages.MustSort(c)
	pages.MustSearch(c)

	if err := p.postService.ISelectListWeb(c, pages, p2); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(pages), nil
}
