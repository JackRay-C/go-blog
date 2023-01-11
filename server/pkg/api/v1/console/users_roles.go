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

type UserRole struct {
	log             logger.Logger
	service         common.BaseService
	userRoleService service.UsersRolesService
}

func NewUserRole() *UserRole {
	return &UserRole{
		log:             global.Log,
		service:         &common.BaseServiceImpl{},
		userRoleService: service.NewUsersRolesService(),
	}
}

// Get 根据用户ID获取角色
func (ur *UserRole) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 判断权限
	if !auth.CheckPermission(c, "users", "authorization") {
		return nil, vo.Forbidden.SetMsg("没有权限. ")
	}

	// 根据角色获取菜单
	var roles []*po.Role

	if err := ur.userRoleService.ISelectUserRoles(c, &po.User{ID: id}, &roles); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&roles), nil
}

// Put 修改用户角色
func (ur *UserRole) Put(c *gin.Context) (*vo.Response, error) {
	// 获取用户ID
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 判断权限
	if !auth.CheckPermission(c, "users", "authorization") {
		return nil, vo.Forbidden.SetMsg("没有权限. ")
	}

	// 绑定参数
	var putUserRoles *dto.PutUserRole
	if err := c.ShouldBindJSON(&putUserRoles); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	// 更新权限
	if err := ur.userRoleService.IUpdateUserRoles(c, &po.User{ID: id}, &putUserRoles.Roles); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success("success"), nil
}
