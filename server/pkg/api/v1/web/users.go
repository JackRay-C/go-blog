package web

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/auth"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	log logger.Logger
	userService service.UserService
	userRoleService service.UsersRolesService
	rolePermissionService service.RolePermissionService
}

func NewUser() *User {
	return &User{
		log: global.Log,
		userService: service.NewUserService(),
		userRoleService: service.NewUsersRolesService(),
		rolePermissionService: service.NewRolesPermissionService(),
	}
}

// Get 根据ID获取用户信息
func (u *User) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	user := &po.User{ID: id}
	if err := u.userService.ISelectOne(c, user); err != nil {
		u.log.Errorf("根据ID查询用户失败 : %s", err)
		return nil, vo.InternalServerError.SetMsg("%s", err)
	} else {
		u.log.Infof("根据ID查询用户成功: %s", user)
		return vo.Success(user), nil
	}
}

// GetUserInfo 获取当前用户信息
func (u *User) GetUserInfo(c *gin.Context) (*vo.Response, error)  {
	// 1、判断是否登录
	if !auth.CheckLogin(c) {
		return nil, vo.NotLogin
	}

	// 2、获取当前用户ID
	currentUserId, _ := c.Get("current_user_id")
	roles := make([]*po.Role, 0)
	permissions := make([]*po.Permissions, 0)

	// 3、获取当前用户的角色
	if err := u.userRoleService.ISelectUserRoles(c, &po.User{ID: currentUserId.(int)}, &roles); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	// 4、根据角色获取当前用户的权限列表
	if len(roles) != 0 {
		if err := u.rolePermissionService.ISelectPermissionByRoles(c, &permissions, roles...); err != nil {
			return nil, vo.InternalServerError.SetMsg("%s", err)
		}
	}

	// 5、获取当前用户的信息
	user := &po.User{ID: currentUserId.(int)}
	err := u.userService.ISelectOne(c, user)
	if err != nil {
		return nil, err
	}

	userInfo := &vo.VUserInfo{
		ID:          currentUserId.(int),
		Username:    user.Username,
		Nickname:    user.Nickname,
		Active:      user.Active,
		Email:       user.Email,
		Avatar:      &po.File{ID: user.Avatar},
		Roles:       roles,
		Permissions: permissions,
		CreatedAt:   user.CreatedAt,
	}
	return vo.Success(userInfo), nil
}