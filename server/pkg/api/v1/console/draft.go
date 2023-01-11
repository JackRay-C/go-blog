package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Draft struct {
	log     logger.Logger
	service common.BaseService
}

func NewDraft() *Draft {
	return &Draft{
		log:     global.Log,
		service: &common.BaseServiceImpl{},
	}
}

func (d *Draft) Get(c *gin.Context) (*vo.Response, error) {
	// 根据HeadID获取Draft
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams
	}

	if !auth.CheckPermission(c, "drafts", "read") {
		return nil, vo.Forbidden
	}

	if err := d.service.ISelectOne(c, &po.Draft{ID: id}); err != nil {
		return nil, err
	}

	return nil, nil
}

func (d *Draft) List(c *gin.Context) (*vo.Response, error) {
	// 获取所有draft
	pager := vo.Pager{PageNo: page.GetPageNo(c), PageSize: page.GetPageSize(c)}

	if !auth.CheckPermission(c, "drafts", "list") {
		return nil, vo.Forbidden
	}

	if err := d.service.ISelectList(c, &pager, &po.Draft{}); err != nil {
		return nil, err
	}

	return vo.Success(pager), nil
}

func (d *Draft) Post(c *gin.Context) (*vo.Response, error) {
	// 新建草稿
	draft := po.Draft{}
	if err := c.ShouldBind(&draft); err != nil {
		return nil, vo.InvalidParams
	}
	if draft.Title == "" {
		return nil, vo.InvalidParams.SetMsg("title isn't be nil. ")
	} else if draft.MarkdownContent == "" {
		return nil, vo.InvalidParams.SetMsg("markdown content isn't be nil ")
	}

	if !auth.CheckPermission(c, "drafts", "add") {
		return nil, vo.Forbidden
	}

	if err := d.service.ICreateOne(c, &draft); err != nil {
		return nil, vo.InternalServerError
	}

	return vo.Success(&draft), nil
}

func (d *Draft) Delete(c *gin.Context) (*vo.Response, error) {
	// 删除草稿
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams
	}

	if !auth.CheckPermission(c, "drafts", "delete") {
		return nil, vo.Forbidden
	}

	if err := d.service.IUnscopeDelete(c, &po.Draft{ID: id}); err != nil {
		return nil, vo.InternalServerError
	}

	return vo.Success("success"), nil
}

func (d *Draft) Put(c *gin.Context) (*vo.Response, error) {
	// 修改草稿
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams
	}

	draft := po.Draft{ID: id}
	if err := c.ShouldBind(&draft); err != nil {
		return nil, vo.InvalidParams
	}

	if !auth.CheckPermission(c, "drafts", "update") {
		return nil, vo.Forbidden
	}

	if err := d.service.IUpdateOne(c, &po.Draft{ID: id}, &draft); err != nil {
		return nil, vo.InternalServerError
	}

	return vo.Success("success"), nil
}
