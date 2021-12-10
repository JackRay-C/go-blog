package v1

import (
	"blog/app/api"
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
	tagService  *service.TagService
	postService *service.PostService
}

func NewTag() *Tag {
	return &Tag{
		log:         global.Logger,
		tagService:  service.NewTagService(),
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

	t.log.Infof("分页查询标签成功. ")
	return response.PagerResponse(&p), nil
}

func (t *Tag) ListPosts(c *gin.Context) (*response.Response, error) {
	post := dto.ListPosts{}
	if err := c.ShouldBind(&post); err != nil {
		t.log.Errorf("绑定参数错误: error: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
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

func (t *Tag) Post(c *gin.Context) (*response.Response, error) {
	tag := &dto.AddTags{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		t.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !api.CheckPermission(c, "tags", "add") {
		t.log.Errorf("创建标签【%s】失败：没有权限！ ", tag.Name)
		return nil, response.Forbidden.SetMsg("创建标签失败，没有权限！")
	}

	if one, err := t.tagService.CreateOne(c, tag); err != nil {
		t.log.Errorf("创建标签【%s】失败：%s", tag.Name, err)
		return nil, response.InternalServerError.SetMsg("创建标签【%s】失败：%s", tag.Name, err)
	} else {
		t.log.Infof("创建标签【%s】成功. ", tag.Name)
		return response.Success(one), nil
	}
}

func (t *Tag) Delete(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		t.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !api.CheckPermission(c, "tags", "delete") {
		t.log.Errorf("删除标签【%d】失败： 没有权限", id)
		return nil, response.Forbidden.SetMsg("删除标签【%d】失败： 没有权限", id)
	}

	if err := t.tagService.DeleteOne(id); err != nil {
		t.log.Info("删除标签失败： ", err)
		return nil, err
	}

	t.log.Info("删除标签【%d】成功 ", id)
	return response.Success(id), nil
}

func (t *Tag) Put(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		t.log.Errorf("更新ID为【%d】的标签失败：参数绑定错误：%s", id, err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var putTag dto.PutTags
	if err := c.ShouldBindJSON(&putTag); err != nil {
		t.log.Errorf("更新ID为【%d】的标签失败：参数绑定错误 %s", id, err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !api.CheckPermission(c, "tags", "update") {
		t.log.Errorf("更新ID为【%d】的标签失败：没有权限", id)
		return nil, response.Forbidden.SetMsg("更新ID为【%d】的标签失败：没有权限", id)
	}

	if tag, err := t.tagService.UpdateOne(c, &putTag); err != nil {
		t.log.Errorf("更新ID为【%d】的标签失败：%s", id, err)
		return nil, err
	} else {
		t.log.Infof("更新Tag成功. ")
		return response.Success(tag), nil
	}
}
