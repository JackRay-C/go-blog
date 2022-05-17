package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/dto"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Tag struct {
	Log         logger.Logger
	tagService  service.TagService
	postService service.PostService
}

func NewTag() *Tag {
	return &Tag{
		Log:         global.Log,
		tagService:  service.NewTagService(),
		postService: service.NewPostService(),
	}
}

func (t *Tag) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	if !auth.CheckPermission(c, "tags", "read") {
		return nil, vo.Forbidden.SetMsg("查询ID为【%d】的标签失败：没有权限", id)
	}

	currentUserId, _ := c.Get("current_user_id")

	tag := &po.Tag{ID: id, UserId: currentUserId.(int)}

	if err := t.tagService.ISelectOne(c, tag); err != nil {
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

	userID, _ := c.Get(global.SessionUserIDKey)
	if err := t.tagService.ISelectList(c, &p, &po.Tag{UserId: userID.(int)}); err != nil {
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

	userID, _ := c.Get(global.SessionUserIDKey)
	tag.UserId = userID.(int)
	if err := t.tagService.ICreateOne(c, tag); err != nil {
		t.Log.Error(err)
		return nil, vo.InternalServerError.SetMsg("创建标签【%s】失败：%s", tag.Name, err)
	}
	return vo.Success(tag), nil
}

func (t *Tag) Delete(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "tags", "delete") {
		return nil, vo.Forbidden.SetMsg("删除ID为【%d】的标签失败：没有权限", id)
	}

	currentUserId, _ := c.Get(global.SessionUserIDKey)
	tag := &po.Tag{ID: id, UserId: currentUserId.(int)}
	if err := t.tagService.IDeleteOne(c, tag); err != nil {
		t.Log.Error(err)
		return nil, vo.InternalServerError
	}

	return vo.Success(tag), nil
}

func (t *Tag) Put(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	var putTag *dto.PutTags
	if err := c.ShouldBindJSON(&putTag); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "tags", "update") {
		return nil, vo.Forbidden.SetMsg("更新ID为【%d】的标签失败：没有权限", id)
	}

	userID, _ := c.Get(global.SessionUserIDKey)
	tag := &po.Tag{ID: id, UserId: userID.(int)}
	if err := t.tagService.IUpdateOne(c, tag, putTag); err != nil {
		t.Log.Error(err)
		return nil, vo.InternalServerError
	}
	return vo.Success(tag), nil
}
