package web

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/model/vo"
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
	rolePermissionService *service.RolesPermissionService
}

func NewUser() *User {
	return &User{
		log: global.Logger,
		userService: service.NewUserService(),
		userRoleService: service.NewUsersRolesService(),
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

// GetUserInfo 获取当前用户信息
func (u *User) GetUserInfo(c *gin.Context) (*response.Response, error)  {
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

	// 5、获取当前用户的信息
	user, err := u.userService.SelectOne(&domain.User{ID: currentUserId.(int)})
	if err != nil {
		return nil, err
	}

	userInfo := &vo.VUserInfo{
		ID:          currentUserId.(int),
		Username:    user.Username,
		Nickname:    user.Nickname,
		Active:      user.Active,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Roles:       roles,
		Permissions: permissions,
		CreatedAt:   user.CreatedAt,
	}
	return response.Success(userInfo), nil
}