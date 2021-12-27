package v1

import (
	"blog/app/api"
	"blog/app/domain"
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

	if !api.CheckPermission(c, "roles", "read") {
		return nil, response.Forbidden.SetMsg("查询角色失败：没有权限. ")
	}

	role := &domain.Role{ID: id}

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

	if !api.CheckPermission(c, "roles", "list") {
		return nil, response.Forbidden.SetMsg("查询角色失败：没有权限. ")
	}

	if err := r.roleService.SelectAll(&p); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&p), nil
}

// 新建角色
func (r *Role) Post(c *gin.Context) (*response.Response, error) {
	role := &domain.Role{}
	if err := c.ShouldBindJSON(&role); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if !api.CheckPermission(c, "roles", "add") {
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

	if !api.CheckPermission(c, "roles", "delete") {
		return nil, response.Forbidden.SetMsg("删除角色失败：没有权限. ")
	}

	role := &domain.Role{ID: id}
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

	if !api.CheckPermission(c, "roles", "update") {
		return nil, response.Forbidden.SetMsg("更新角色失败：没有权限. ")
	}

	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}
	role.ID = id

	if err := r.roleService.UpdateOne(&role); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&role), nil
}


func (r *Role) PostMenus(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var requestMenus dto.AddRoleMenus
	if err := c.ShouldBindJSON(&requestMenus); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := r.roleService.UpdateMenus(&domain.Role{ID: id}, requestMenus.Menus); err != nil {
		return nil, response.InternalServerError.SetMsg("授权ID为【%d】的角色菜单失败：%s", id, err)
	}

	return response.Success(requestMenus),nil
}

func (r *Role) PutMenus(c *gin.Context) (*response.Response, error) {
	r.log.Infof("修改角色菜单")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var requestMenus dto.AddRoleMenus
	if err := c.ShouldBindJSON(&requestMenus); err != nil {
		r.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := r.roleService.UpdateMenus(&domain.Role{ID: id}, requestMenus.Menus); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(requestMenus), nil
}

