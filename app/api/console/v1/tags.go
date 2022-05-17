package v1

import (
	"blog/app/model/dto"
	"blog/app/model/po"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/app/utils/auth"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Tag struct {
	log         logger.Logger
	tagService  *service.TagService
	postService service.PostService
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

	if !auth.CheckPermission(c, "tags", "read") {
		return nil, response.Forbidden.SetMsg("查询ID为【%d】的标签失败：没有权限", id)
	}

	currentUserId, _ := c.Get("current_user_id")

	tag := &po.Tag{ID: id, UserId: currentUserId.(int)}

	if err := t.tagService.SelectOne(tag); err != nil {
		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的标签失败 : %s", id, err)
	}
	return response.Success(tag), nil
}

func (t *Tag) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}
	if !auth.CheckPermission(c, "tags", "list") {
		return nil, response.Forbidden.SetMsg("查询标签列表失败：没有权限")
	}

	if err := t.tagService.SelectAll(c, &p); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&p), nil
}

func (t *Tag) Post(c *gin.Context) (*response.Response, error) {
	tag := &dto.AddTags{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "tags", "add") {
		return nil, response.Forbidden.SetMsg("新建标签失败：没有权限")
	}

	if one, err := t.tagService.CreateOne(c, tag); err != nil {
		return nil, response.InternalServerError.SetMsg("创建标签【%s】失败：%s", tag.Name, err)
	} else {
		return response.Success(one), nil
	}
}

func (t *Tag) Delete(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "tags", "delete") {
		return nil, response.Forbidden.SetMsg("删除ID为【%d】的标签失败：没有权限", id)
	}

	currentUserId, _ := c.Get("current_user_id")

	if err := t.tagService.DeleteOne(currentUserId.(int), id); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(id), nil
}

func (t *Tag) Put(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var putTag dto.PutTags
	if err := c.ShouldBindJSON(&putTag); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "tags", "update") {
		return nil, response.Forbidden.SetMsg("更新ID为【%d】的标签失败：没有权限", id)
	}

	if tag, err := t.tagService.UpdateOne(c, &putTag); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	} else {
		return response.Success(tag), nil
	}
}
