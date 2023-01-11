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

type Tag struct {
	Log     logger.Logger
	service common.BaseService
	postTagService service.PostsTagsService
}

func NewTag() *Tag {
	return &Tag{
		Log:     global.Log,
		service: &common.BaseServiceImpl{},
		postTagService: service.NewPostsTagsService(),
	}
}

func (t *Tag) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	if !auth.CheckPermission(c, "tags", "read") {
		return nil, vo.Forbidden.SetMsg("查询ID为【%d】的标签失败：没有权限", id)
	}

	tag := &po.Tag{ID: id, UserID: auth.GetCurrentUserId(c)}

	if err := t.service.ISelectOne(c, tag); err != nil {
		return nil, err
	}
	return vo.Success(tag), nil
}

func (t *Tag) List(c *gin.Context) (*vo.Response, error) {
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	if !auth.CheckPermission(c, "tags", "list") {
		return nil, vo.Forbidden.SetMsg("查询标签列表失败：没有权限")
	}


	if err := t.service.ISelectList(c, &p, &po.Tag{UserID: auth.GetCurrentUserId(c)}); err != nil {
		t.Log.Errorf("%s", err)
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&p), nil
}

func (t *Tag) Post(c *gin.Context) (*vo.Response, error) {

	tag := &po.Tag{}
	if err := c.ShouldBindJSON(&tag); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "tags", "add") {
		return nil, vo.Forbidden.SetMsg("新建标签失败：没有权限")
	}

	tag.UserID = auth.GetCurrentUserId(c)
	if err := t.service.ICreateOne(c, tag); err != nil {
		t.Log.Error(err)
		return nil, vo.InternalServerError.SetMsg("创建标签【%s】失败：%s", tag.Name, err)
	}

	return vo.Success(tag), nil
}

func (t *Tag) Delete(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "tags", "delete") {
		return nil, vo.Forbidden.SetMsg("删除ID为【%d】的标签失败：没有权限", id)
	}

	tag := &po.Tag{ID: id, UserID: auth.GetCurrentUserId(c)}
	if err := t.service.IDeleteOne(c, tag); err != nil {
		t.Log.Error(err)
		return nil, vo.InternalServerError
	}

	return vo.Success(tag), nil
}

func (t *Tag) Put(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	putTag := &po.Tag{}
	if err := c.ShouldBindJSON(&putTag); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "tags", "update") {
		return nil, vo.Forbidden.SetMsg("更新ID为【%d】的标签失败：没有权限", id)
	}

	tag := &po.Tag{ID: id, UserID: auth.GetCurrentUserId(c)}
	if err := t.service.IUpdateOne(c, tag, putTag); err != nil {
		t.Log.Error(err)
		return nil, vo.InternalServerError
	}

	return vo.Success(tag), nil
}

func (t *Tag) ListPosts(c *gin.Context) (*vo.Response, error) {
	// 获取tagid
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	post := po.Post{}
	if err := c.ShouldBind(&post); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	if !auth.CheckPermission(c, "posts", "list") {
		return nil, vo.Forbidden.SetMsg("查询标签列表失败：没有权限")
	}

	if err := t.postTagService.ISelectPostsByTag(c, &p, &po.Tag{ID: id}); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&p), err
}