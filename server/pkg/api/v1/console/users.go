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
	"blog/pkg/utils/page"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	log                   logger.Logger
	service               common.BaseService
	postService           service.PostService
	fileService           service.FileService
	userRoleService       service.UsersRolesService
	rolePermissionService service.RolePermissionService
}

func NewUser() *User {
	return &User{log: global.Log,
		service:               &common.BaseServiceImpl{},
		postService:           service.NewPostService(),
		fileService:           service.NewFileService(),
		userRoleService:       service.NewUsersRolesService(),
		rolePermissionService: service.NewRolesPermissionService(),
	}
}

// GetUserInfo 获取当前用户信息
func (u *User) GetUserInfo(c *gin.Context) (*vo.Response, error) {
	// 1、获取当前用户ID
	currentUserId, _ := c.Get(global.SessionUserIDKey)
	roles := make([]*po.Role, 0)
	permissions := make([]*po.Permissions, 0)

	// 2、获取当前用户的角色
	if err := u.userRoleService.ISelectUserRoles(c, &po.User{ID: currentUserId.(int64)}, &roles); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	// 3、根据角色获取当前用户的权限列表
	if len(roles) != 0 {
		if err := u.rolePermissionService.ISelectPermissionByRoles(c, &permissions, roles...); err != nil {
			return nil, vo.InternalServerError.SetMsg("%s", err)
		}
	}

	// 4、获取当前用户的信息
	user := &po.User{ID: currentUserId.(int64)}
	err := u.service.ISelectOne(c, user)
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

// 更新用户信息
func (u *User) Put(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	var putUser *dto.PutUser
	if err := c.ShouldBindJSON(&putUser); err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if putUser.ID == 0 {
		putUser.ID = id
	}

	currentUserId, _ := c.Get(global.SessionUserIDKey)
	// 修改别人的信息
	if putUser.ID != currentUserId.(int64) {
		if !auth.CheckPermission(c, "users", "update") {
			return nil, vo.Forbidden.SetMsg("修改ID为【%d】的用户信息失败：没有权限. ", id)
		}
	}

	// 修改自己的信息
	user := &po.User{ID: id}

	if err := u.service.IUpdateOne(c, user, user); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(user), nil
}

// 分页获取用户列表
func (u *User) List(c *gin.Context) (*vo.Response, error) {
	p := vo.Pager{
		PageNo:   page.GetPageNo(c),
		PageSize: page.GetPageSize(c),
	}

	active, err := strconv.ParseInt(c.DefaultQuery("active", "0"), 10, 8)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "users", "list") {
		return nil, vo.Forbidden
	}

	if err := u.service.ISelectList(c, &p, &po.User{Active: int8(active)}); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&p), nil
}

// Get 根据ID获取用户信息
func (u *User) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	if !auth.CheckPermission(c, "users", "read") {
		return nil, vo.Forbidden.SetMsg("查询用户信息失败：没有权限")
	}

	user := &po.User{ID: id}

	if err := u.service.ISelectOne(c, user); err != nil {
		u.log.Error(err)
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(user), nil
}

// Post 创建用户
func (u *User) Post(c *gin.Context) (*vo.Response, error) {
	u.log.Infof("创建用户")
	createUser := &po.User{}
	if err := c.ShouldBindJSON(createUser); err != nil {
		u.log.Errorf("参数绑定错误: %s", err)
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := u.service.ICreateOne(c, createUser); err != nil {
		u.log.Errorf("创建用户失败：%s", err)
		return nil, vo.InternalServerError.SetMsg("创建用户失败：%s", err)
	}

	u.log.Infof("创建用户成功: %s", createUser)
	return vo.Success(createUser), nil
}

// Delete 根据ID删除用户
func (u *User) Delete(c *gin.Context) (*vo.Response, error) {
	u.log.Infof("删除用户")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		u.log.Errorf("参数绑定错误：%s", err)
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := u.service.IDeleteOne(c, &po.User{ID: id}); err != nil {
		u.log.Info("查询删除用户失败： ", err)
		return nil, vo.InternalServerError.SetMsg("删除用户失败：%s", err)
	}

	u.log.Info("删除用户成功 ", id)
	return vo.Success("删除成功"), nil
}

//
//func (u *User) ListRole(c *gin.Context) (*vo.Response, error) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		return nil, vo.InvalidParams.SetMsg("ID is required. ")
//	}
//	u.logs.Infof("根据ID查询用户角色: ID[%d]", id)
//
//	var roles []*domain.Role
//	if err := u.userService.SelectRoles(&domain.User{ID: id}, &roles); err != nil {
//		return nil, err
//	}
//	return vo.Success(roles), nil
//}
//
//func (u *User) PostRole(c *gin.Context) (*vo.Response, error) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		u.logs.Errorf("参数绑定错误： %s", err)
//		return nil, vo.InvalidParams.SetMsg("%s", err)
//	}
//
//	var requestRoles dto.AddUserRole
//	if err := c.ShouldBindJSON(&requestRoles); err != nil {
//		u.logs.Errorf("参数绑定错误： %s", err)
//		return nil, vo.InvalidParams.SetMsg("%s", err)
//	}
//
//	u.logs.Infof("给用户: %s 添加角色：%s", id, requestRoles.Roles)
//	if err := u.userService.InsertUserRoles(&domain.User{ID: id}, requestRoles.Roles); err != nil {
//		return nil, err
//	}
//
//	return vo.Success(requestRoles), nil
//}
//
//func (u *User) PutRole(c *gin.Context) (*vo.Response, error) {
//	// 修改用户角色
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id == 0 {
//		u.logs.Errorf("参数绑定错误： %s", err)
//		return nil, vo.InvalidParams.SetMsg("%s", err)
//	}
//
//	var requestRoles dto.AddUserRole
//	if err := c.ShouldBindJSON(&requestRoles); err != nil {
//		u.logs.Errorf("参数绑定错误： %s", err)
//		return nil, vo.InvalidParams.SetMsg("%s", err)
//	}
//
//	u.logs.Infof("修改用户: %s 角色：%s", id, requestRoles.Roles)
//	if err := u.userService.UpdateUserRoles(&domain.User{ID: id}, requestRoles.Roles); err != nil {
//		return nil, err
//	}
//	return vo.Success(requestRoles), nil
//}
