package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/dto"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RolePermission struct {
	log                   logger.Logger
	service               common.BaseService
	rolePermissionService service.RolePermissionService
}

func NewRolePermission() *RolePermission {
	return &RolePermission{
		log:                   global.Log,
		service:               &common.BaseServiceImpl{},
		rolePermissionService: service.NewRolesPermissionService(),
	}
}

func (r *RolePermission) Get(c *gin.Context) (*vo.Response, error) {
	// 1、获取角色ID
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断权限
	if !auth.CheckPermission(c, "roles", "assign") {
		return nil, vo.Forbidden.SetMsg("没有权限. ")
	}

	// 3、根据角色获取权限列表
	var permissions []*po.Permissions

	if err := r.rolePermissionService.ISelectPermissionByRoles(c, &permissions, &po.Role{ID: id}); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&permissions), nil
}

func (r *RolePermission) Put(c *gin.Context) (*vo.Response, error) {
	// 1、获取角色ID
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 2、判断权限
	if !auth.CheckPermission(c, "roles", "assign") {
		return nil, vo.Forbidden.SetMsg("没有权限. ")
	}

	// 3、绑定数据
	var permissions dto.PutRolePermission
	if err := c.ShouldBindJSON(&permissions); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 4、修改角色权限列表
	if err := r.rolePermissionService.IUpdateRolePermissions(c, &po.Role{ID: id}, permissions.Permissions); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success("success"), nil
}
