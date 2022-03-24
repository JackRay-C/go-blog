package v1

import (
	"blog/app/model/dto"
	"blog/app/model/po"
	"blog/app/response"
	"blog/app/service"
	"blog/app/utils/auth"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RolePermission struct {
	log           logger.Logger
	roleService *service.RoleService
	rolePermissionService *service.RolesPermissionService
}

func NewRolePermission() *RolePermission {
	return &RolePermission{
		log: global.Logger,
		roleService: service.NewRoleService(),
		rolePermissionService: service.NewRolesPermissionService(),
	}
}

func (r *RolePermission) Get(c *gin.Context) (*response.Response, error)  {
	// 1、获取角色ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断权限
	if !auth.CheckPermission(c, "roles", "assign") {
		return nil, response.Forbidden.SetMsg("没有权限. ")
	}

	// 3、根据角色获取权限列表
	var permissions []*po.Permissions

	if err := r.rolePermissionService.SelectRolePermissions(&po.Role{ID: id}, &permissions); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(&permissions), nil
}

func (r *RolePermission) Put(c *gin.Context) (*response.Response, error)  {
	// 1、获取角色ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断权限
	if !auth.CheckPermission(c, "roles", "assign") {
		return nil, response.Forbidden.SetMsg("没有权限. ")
	}

	// 3、绑定数据
	var permissions dto.PutRolePermission
	if err := c.ShouldBindJSON(&permissions); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 4、修改角色权限列表
	if err := r.rolePermissionService.UpdateRolePermissions(&po.Role{ID: id}, permissions.Permissions); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success("success"), nil
}
