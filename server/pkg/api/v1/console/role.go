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

type Role struct {
	log           logger.Logger
	service common.BaseService
}

func NewRole() *Role {
	return &Role{
		log: global.Log,
		service: &common.BaseServiceImpl{},
	}
}

// 获取角色信息
func (r *Role) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	if !auth.CheckPermission(c, "roles", "read") {
		return nil, vo.Forbidden.SetMsg("查询角色失败：没有权限. ")
	}

	role := &po.Role{ID: id}

	if err := r.service.ISelectOne(c, role); err != nil {
		return nil, vo.InternalServerError.SetMsg("查询ID为【%d】的角色信息失败：%s",id, err)
	}

	return vo.Success(role), nil
}

// 获取角色列表
func (r *Role) List(c *gin.Context) (*vo.Response, error) {
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	if !auth.CheckPermission(c, "roles", "list") {
		return nil, vo.Forbidden.SetMsg("查询角色失败：没有权限. ")
	}

	if err := r.service.ISelectList(c, &p, &po.Role{}); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&p), nil
}

// 新建角色
func (r *Role) Post(c *gin.Context) (*vo.Response, error) {
	role := &po.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "roles", "add") {
		return nil, vo.Forbidden.SetMsg("新建角色失败：没有权限. ")
	}

	if err := r.service.ICreateOne(c, role); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(role), nil
}

// 删除角色
func (r *Role) Delete(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	if !auth.CheckPermission(c, "roles", "delete") {
		return nil, vo.Forbidden.SetMsg("删除角色失败：没有权限. ")
	}

	role := &po.Role{ID: id}
	if err := r.service.IDeleteOne(c,role); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}
	return vo.Success("删除成功"), nil
}

// 修改角色
func (r *Role) Put(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "roles", "update") {
		return nil, vo.Forbidden.SetMsg("更新角色失败：没有权限. ")
	}

	var role po.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}
	role.ID = id

	if err := r.service.IUpdateOne(c, &role, &role); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&role), nil
}

