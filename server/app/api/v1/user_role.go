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

type UserRole struct {
	log           logger.Logger
	roleService *service.RoleService
	userRoleService *service.UsersRolesService
}

func NewUserRole() *UserRole {
	return &UserRole{
		log: global.Logger,
		roleService: service.NewRoleService(),
		userRoleService: service.NewUsersRolesService(),
	}
}

// Get 根据用户ID获取角色
func (ur *UserRole) Get(c *gin.Context) (*response.Response, error)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 判断权限
	if !api.CheckPermission(c, "users", "authorization") {
		return nil, response.Forbidden.SetMsg("没有权限. ")
	}

	// 根据角色获取菜单
	var roles []*domain.Role

	if err := ur.userRoleService.SelectUserRoles(&domain.User{ID: id}, &roles); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&roles), nil
}

// Put 修改用户角色
func (ur *UserRole) Put(c *gin.Context) (*response.Response, error)  {
	// 获取用户ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 判断权限
	if !api.CheckPermission(c, "users", "authorization") {
		return nil, response.Forbidden.SetMsg("没有权限. ")
	}

	// 绑定参数
	var putUserRoles *dto.PutUserRole
	if err := c.ShouldBindJSON(&putUserRoles); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 更新权限
	if err := ur.userRoleService.UpdateUserRoles(&domain.User{ID: id}, putUserRoles.Roles); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success("success"), nil
}