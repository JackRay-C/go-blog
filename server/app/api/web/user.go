package web

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

type User struct {
	log logger.Logger
	userService *service.UserService
	userRoleService *service.UsersRolesService
	roleMenuService *service.RolesMenusService
	rolePermissionService *service.RolesPermissionService
}

func NewUser() *User {
	return &User{
		log: global.Logger,
		userService: service.NewUserService(),
		userRoleService: service.NewUsersRolesService(),
		roleMenuService: service.NewRolesMenusService(),
		rolePermissionService: service.NewRolesPermissionService(),
	}
}

// Get 根据ID获取用户信息
func (u *User) Get(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if one, err := u.userService.SelectOne(&domain.User{ID: id}); err != nil {
		u.log.Errorf("根据ID查询用户失败 : %s", err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	} else {
		u.log.Infof("根据ID查询用户成功: %s", one)
		return response.Success(one), nil
	}
}

// GetRoles 获取当前用户的角色
func (u *User) GetRoles(c *gin.Context) (*response.Response, error)  {
	// 1、判断是否登录
	if !api.CheckLogin(c) {
		return nil, response.NotLogin
	}

	// 2、获取当前用户ID
	currentUserId, _ := c.Get("current_user_id")
	roles := make([]*domain.Role, 0)

	// 3、根据当前用户ID查询所有角色
	if err := u.userRoleService.SelectUserRoles(&domain.User{ID: currentUserId.(int)}, &roles); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&roles),nil
}

// GetMenus 获取当前用户的菜单
func (u *User) GetMenus(c *gin.Context) (*response.Response, error) {
	// 1、判断是否登录
	if !api.CheckLogin(c) {
		return nil, response.NotLogin
	}

	// 2、获取当前用户ID
	currentUserId, _ := c.Get("current_user_id")
	roles := make([]*domain.Role, 0)
	menus := make([]*domain.Menu, 0)

	// 3、获取当前用户的角色
	if err := u.userRoleService.SelectUserRoles(&domain.User{ID: currentUserId.(int)}, &roles); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	// 4、根据角色获取菜单
	if len(roles) != 0 {
		if err := u.roleMenuService.SelectMenusByRoles(&menus, roles...); err != nil {
			return nil, response.InternalServerError.SetMsg("%s", err)
		}
	}

	return response.Success(&menus), nil
}

func (u *User) GetPermissions(c *gin.Context) (*response.Response, error) {
	// 1、判断是否登录
	if !api.CheckLogin(c) {
		return nil, response.NotLogin
	}

	// 2、获取当前用户ID
	currentUserId, _ := c.Get("current_user_id")
	roles := make([]*domain.Role, 0)
	permissions := make([]*domain.Permissions, 0)


	// 3、获取当前用户的角色
	if err := u.userRoleService.SelectUserRoles(&domain.User{ID: currentUserId.(int)}, &roles); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	// 4、根据角色获取当前用户的权限列表
	if len(roles) != 0 {
		if err := u.rolePermissionService.SelectPermissionByRoles(&permissions, roles...); err != nil {
			return nil, response.InternalServerError.SetMsg("%s", err)
		}
	}

	return response.Success(&permissions), nil
}
