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

type Permission struct {
	log logger.Logger
	permissionService *service.PermissionService
}

func NewPermission() *Permission {
	return &Permission{
		log: global.Logger,
		permissionService: service.NewPermissionService(),
	}
}

// 获取一条权限记录
func (p *Permission) Get(c *gin.Context) (*response.Response, error) {
	// 1、获取ID参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, response.Forbidden.SetMsg("获取ID为【%d】的权限记录失败：非管理员. ", id)
	}

	// 3、检查是否有权限
	if !auth.CheckPermission(c, "permission", "read") {
		return nil, response.Forbidden.SetMsg("获取ID为【%d】的权限记录失败：没有权限. ", id)
	}

	// 3、查询记录
	permission := &po.Permissions{ID: id}
	if err := p.permissionService.SelectOne(permission); err != nil {
		return nil, err
	}

	return response.Success(permission), nil
}


// 获取所有权限列表
func (p *Permission) List(c *gin.Context) (*response.Response, error) {
	// 1、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, response.Forbidden
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "permission", "list") {
		return nil, response.Forbidden
	}

	page := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	// 3、查询所有权限
	if err := p.permissionService.SelectAll(&page); err != nil {
		return nil, err
	}
	return response.Success(&page),nil
}

// 添加一条权限列表
func (p *Permission) Post(c *gin.Context) (*response.Response, error) {
	// 1、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, response.Forbidden
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "permission", "add") {
		return nil, response.Forbidden
	}

	// 3、绑定参数
	var permission *po.Permissions
	if err := c.ShouldBindJSON(&permission); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	// 4、创建
	if err := p.permissionService.CreateOne(permission); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success(permission),nil
}

// 删除一条权限记录
func (p *Permission) Delete(c *gin.Context) (*response.Response, error) {
	// 1、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, response.Forbidden
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "permission", "delete") {
		return nil, response.Forbidden
	}

	// 3、获取ID参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	// 4、删除记录
	if err := p.permissionService.DeleteOne(&po.Permissions{ID: id}); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	return response.Success("success"), nil
}

// 更新一条权限记录
func (p *Permission) Put(c *gin.Context) (*response.Response, error) {
	// 1、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, response.Forbidden
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "permission", "update") {
		return nil, response.Forbidden
	}

	// 3、绑定json
	var permission *po.Permissions

	if err := c.ShouldBindJSON(&permission); err != nil {
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := p.permissionService.UpdateOne(permission); err != nil {
		return nil, response.InternalServerError.SetMsg("%s", err)
	}
	return response.Success(permission), nil
}
