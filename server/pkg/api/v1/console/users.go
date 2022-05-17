package console

import (
	"blog/internal/logger"
	"blog/pkg/global"
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
	log            logger.Logger
	userService    service.UserService
	subjectService service.SubjectService
	postService    service.PostService
	fileService    service.FileService
}

func NewUser() *User {
	return &User{log: global.Log,
		userService:    service.NewUserService(),
		subjectService: service.NewSubjectService(),
		postService:    service.NewPostService(),
		fileService:    service.NewFileService(),
	}
}

// 更新用户信息
func (u *User) Put(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
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

	currentUserId, _ := c.Get("current_user_id")
	// 修改别人的信息
	if putUser.ID != currentUserId.(int) {
		if !auth.CheckPermission(c, "users", "update") {
			return nil, vo.Forbidden.SetMsg("修改ID为【%d】的用户信息失败：没有权限. ", id)
		}
	}

	// 修改自己的信息
	user := &po.User{ID: id}

	if err := u.userService.IUpdateOne(c, user, user);err != nil {
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

	active := c.DefaultQuery("active", "0")
	atoi, err := strconv.Atoi(active)
	if err != nil {
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if !auth.CheckPermission(c, "users", "list") {
		return nil, vo.Forbidden.SetMsg("查询用户列表失败：没有权限")
	}

	if err := u.userService.ISelectList(c, &p, &po.User{Active: int8(atoi)}); err != nil {
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(&p), nil
}

// 根据ID获取用户信息
func (u *User) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	if !auth.CheckPermission(c, "users", "read") {
		return nil, vo.Forbidden.SetMsg("查询用户信息失败：没有权限")
	}

	user := &po.User{ID: id}

	if err := u.userService.ISelectOne(c, user);err != nil {
		u.log.Error(err)
		return nil, vo.InternalServerError.SetMsg("%s", err)
	}

	return vo.Success(user), nil
}

// 创建用户
func (u *User) Post(c *gin.Context) (*vo.Response, error) {
	u.log.Infof("创建用户")
	createUser := &po.User{}
	if err := c.ShouldBindJSON(createUser); err != nil {
		u.log.Errorf("参数绑定错误: %s", err)
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := u.userService.ICreateOne(c, createUser); err != nil {
		u.log.Errorf("创建用户失败：%s", err)
		return nil, vo.InternalServerError.SetMsg("创建用户失败：%s", err)
	}

	u.log.Infof("创建用户成功: %s", createUser)
	return vo.Success(createUser), nil
}

// 根据ID删除用户
func (u *User) Delete(c *gin.Context) (*vo.Response, error) {
	u.log.Infof("删除用户")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.log.Errorf("参数绑定错误：%s", err)
		return nil, vo.InvalidParams.SetMsg("%s", err)
	}

	if err := u.userService.IDeleteOne(c, &po.User{ID: id}); err != nil {
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

