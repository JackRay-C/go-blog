package v1

import (
	"blog/app/api"
	"blog/app/domain"
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
}

func NewRoleMenu() *RoleMenu {
	return &RoleMenu{
		log: global.Logger,
		roleService: service.NewRoleService(),
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

	if err := r.roleService.SelectMenus(&menus, &domain.Role{ID: id}); err != nil {
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

	// 2、获取角色的菜单

	return response.Success("success"), nil
}