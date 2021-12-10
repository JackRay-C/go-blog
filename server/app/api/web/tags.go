package web

import (
	"blog/app/model/dto"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Tag struct {
	log         logger.Logger
	tagService *service.TagService
	postService *service.PostService
}


func NewTag() *Tag {
	return &Tag{
		log: global.Logger,
		tagService: service.NewTagService(),
		postService: service.NewPostService(),
	}
}

func (t *Tag) Get(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if tag, err := t.tagService.SelectOne(id); err != nil {
		t.log.Errorf("查询ID为【%d】的标签失败 : %s", id, err)
		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的标签失败 : %s", id, err)
	} else {
		t.log.Infof("根据ID查询专题成功: %s", tag)
		return response.Success(tag), nil
	}
}

func (t *Tag) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	if err := t.tagService.SelectAll(&p); err != nil {
		t.log.Errorf("分页查询标签失败： %s", err)
		return nil, err
	}

	t.log.Infof("分页查询标签成功：%s", &p)
	return response.PagerResponse(&p), nil
}

func (t *Tag) ListPosts(c *gin.Context) (*response.Response, error)  {
	post := dto.ListPosts{}
	if err := c.ShouldBind(&post); err != nil {
		t.log.Errorf("绑定参数错误: error: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0{
		t.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	post.TagId = id

	page := pager.Pager{}
	if err := t.postService.SelectAll1(c, &page, &post); err != nil {
		t.log.Errorf("根据TagId【%id】查询博客: %s", err)
		return nil, response.InternalServerError.SetMsg("根据TagId【%id】查询博客失败：%s", err)
	}

	// 3、返回查询结果
	t.log.Infof("根据TagId【%id】查询博客成功: [第 %d 页，总页数：%d, 总行数：%d]", id, page.PageNo, page.PageCount, page.TotalRows)
	return response.PagerResponse(&page), nil
}