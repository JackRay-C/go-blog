package v1

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RoleMenu struct {
	log           logger.Logger
	roleService *service.RoleService
	roleMenusService *service.RolesMenusService
}

func NewRoleMenu() *RoleMenu {
	return &RoleMenu{
		log: global.Logger,
		roleService: service.NewRoleService(),
		roleMenusService: service.NewRolesMenusService(),
	}
}

// 根据角色获取菜单
func (r *RoleMenu) Get(c *gin.Context) (*response.Response, error)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 判断权限
	if !api.CheckPermission(c, "roles", "assign") {
		return nil, response.Forbidden.SetMsg("没有权限. ")
	}

	// 根据角色获取菜单
	var menus []*domain.Menu

	if err := r.roleMenusService.SelectRoleMenus(&domain.Role{ID: id}, &menus); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&menus), nil
}

// 更新角色的菜单
func (r *RoleMenu) Put(c *gin.Context) (*response.Response, error) {
	// 1、获取角色ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断权限
	if !api.CheckPermission(c, "roles", "assign") {
		return nil, response.Forbidden.SetMsg("没有权限. ")
	}

	// 3、更新角色的菜单
	var menus dto.PutRoleMenus
	if err := c.ShouldBindJSON(&menus); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := r.roleMenusService.UpdateRoleMenus(&domain.Role{ID: id}, menus.Menus); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success("success"), nil
}