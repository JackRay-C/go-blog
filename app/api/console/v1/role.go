package v1

import (
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

type Role struct {
	log           logger.Logger
	roleService *service.RoleService
}

func NewRole() *Role {
	return &Role{
		log: global.Logger,
		roleService: service.NewRoleService(),
	}
}

// 获取角色信息
func (r *Role) Get(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if !auth.CheckPermission(c, "roles", "read") {
		return nil, response.Forbidden.SetMsg("查询角色失败：没有权限. ")
	}

	role := &po.Role{ID: id}

	if err := r.roleService.SelectOne(role); err != nil {
		return nil, response.InternalServerError.SetMsg("查询ID为【%d】的角色信息失败：%s",id, err)
	}

	return response.Success(role), nil
}

// 获取角色列表
func (r *Role) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	if !auth.CheckPermission(c, "roles", "list") {
		return nil, response.Forbidden.SetMsg("查询角色失败：没有权限. ")
	}

	if err := r.roleService.SelectAll(&p); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&p), nil
}

// 新建角色
func (r *Role) Post(c *gin.Context) (*response.Response, error) {
	role := &po.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "roles", "add") {
		return nil, response.Forbidden.SetMsg("新建角色失败：没有权限. ")
	}

	if err := r.roleService.CreateOne(role); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(role), nil
}

// 删除角色
func (r *Role) Delete(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if !auth.CheckPermission(c, "roles", "delete") {
		return nil, response.Forbidden.SetMsg("删除角色失败：没有权限. ")
	}

	role := &po.Role{ID: id}
	if err := r.roleService.DeleteOne(role); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success("删除成功"), nil
}

// 修改角色
func (r *Role) Put(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "roles", "update") {
		return nil, response.Forbidden.SetMsg("更新角色失败：没有权限. ")
	}

	var role po.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	role.ID = id

	if err := r.roleService.UpdateOne(&role); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&role), nil
}

