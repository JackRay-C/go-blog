package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Permission struct {
	log logger.Logger
	permissionService service.PermissionService
}

func NewPermission() *Permission {
	return &Permission{
		log: global.Log,
		permissionService: service.NewPermissionService(),
	}
}

// 获取一条权限记录
func (p *Permission) Get(c *gin.Context) (*vo.Response, error) {
	// 1、获取ID参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	// 2、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, vo.Forbidden.SetMsg("获取ID为【%d】的权限记录失败：非管理员. ", id)
	}

	// 3、检查是否有权限
	if !auth.CheckPermission(c, "permission", "read") {
		return nil, vo.Forbidden.SetMsg("获取ID为【%d】的权限记录失败：没有权限. ", id)
	}

	// 3、查询记录
	permission := &po.Permissions{ID: id}
	if err := p.permissionService.ISelectOne(c, permission); err != nil {
		return nil, err
	}

	return vo.Success(permission), nil
}


// 获取所有权限列表
func (p *Permission) List(c *gin.Context) (*vo.Response, error) {
	// 1、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, vo.Forbidden
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "permission", "list") {
		return nil, vo.Forbidden
	}

	pager := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	// 3、查询所有权限
	if err := p.permissionService.ISelectList(c, &pager, &po.Permissions{}); err != nil {
		return nil, err
	}
	return vo.Success(&pager),nil
}

// 添加一条权限列表
func (p *Permission) Post(c *gin.Context) (*vo.Response, error) {
	// 1、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, vo.Forbidden
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "permission", "add") {
		return nil, vo.Forbidden
	}

	// 3、绑定参数
	var permission *po.Permissions
	if err := c.ShouldBindJSON(&permission); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 4、创建
	if err := p.permissionService.ICreateOne(c, permission); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(permission),nil
}

// 删除一条权限记录
func (p *Permission) Delete(c *gin.Context) (*vo.Response, error) {
	// 1、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, vo.Forbidden
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "permission", "delete") {
		return nil, vo.Forbidden
	}

	// 3、获取ID参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	// 4、删除记录
	if err := p.permissionService.IDeleteOne(c, &po.Permissions{ID: id}); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success("success"), nil
}

// 更新一条权限记录
func (p *Permission) Put(c *gin.Context) (*vo.Response, error) {
	// 1、检查是否是管理员
	if !auth.CheckAdmin(c) {
		return nil, vo.Forbidden
	}

	// 2、检查是否有权限
	if !auth.CheckPermission(c, "permission", "update") {
		return nil, vo.Forbidden
	}

	// 3、绑定json
	var permission *po.Permissions

	if err := c.ShouldBindJSON(&permission); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := p.permissionService.IUpdateOne(c, permission, permission); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}
	return vo.Success(permission), nil
}
