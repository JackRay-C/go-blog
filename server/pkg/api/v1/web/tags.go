package web

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Tag struct {
	log         logger.Logger
	tagService  service.TagService
	postService service.PostService
}

func NewTag() *Tag {
	return &Tag{
		log:         global.Log,
		tagService:  service.NewTagService(),
		postService: service.NewPostService(),
	}
}

func (t *Tag) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	tag := &po.Tag{ID: id}
	if err := t.tagService.ISelectOne(c, tag); err != nil {
		return nil, vo.InternalServerError.SetMsg("查询ID为【%d】的标签失败 : %s", id, err)
	}
	return vo.Success(tag), nil
}

func (t *Tag) List(c *gin.Context) (*vo.Response, error) {
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	if err := t.tagService.ISelectList(c, &p, &po.Tag{}); err != nil {
		t.log.Errorf("分页查询标签失败： %s", err)
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	t.log.Infof("分页查询标签成功：%s", &p)
	return vo.Success(&p), nil
}