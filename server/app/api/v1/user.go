package v1

import (
	"blog/app/api"
	"blog/app/domain"
	"blog/app/model/dto"
	"blog/app/pager"
	"blog/app/request"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	log            logger.Logger
	userService    *service.UserService
	subjectService *service.SubjectService
	postService    *service.PostService
	fileService    *service.FileService
}

func NewUser() *User {
	return &User{log: global.Logger,
		userService:    service.NewUserService(),
		subjectService: service.NewSubjectService(),
		postService:    service.NewPostService(),
		fileService:    service.NewFileService(),
	}
}

func (u *User) Put(c *gin.Context) (*response.Response, error) {
	panic("implement me")
}

// 分页获取用户列表
func (u *User) List(c *gin.Context) (*response.Response, error) {
	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}

	if !api.CheckPermission(c, "users", "list") {
		return nil, response.Forbidden.SetMsg("查询用户列表失败：没有权限")
	}

	if err := u.userService.SelectAll(&p, &domain.User{}); err != nil {
		u.log.Errorf("查询用户列表失败： %s", err)
		return nil, err
	}

	u.log.Infof("查询用户列表成功：%s", &p)
	return response.PagerResponse(&p), nil
}

// 根据ID获取用户信息
func (u *User) Get(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}

	if !api.CheckPermission(c, "users", "read") {
		return nil, response.Forbidden.SetMsg("查询用户信息失败：没有权限")
	}

	if one, err := u.userService.SelectOneById(id); err != nil {
		u.log.Errorf("根据ID查询用户失败 : %s", err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	} else {
		u.log.Infof("根据ID查询用户成功: %s", one)
		return response.Success(one), nil
	}
}

// 创建用户
func (u *User) Post(c *gin.Context) (*response.Response, error) {
	u.log.Infof("创建用户")
	createUser := &domain.User{}
	if err := c.ShouldBindJSON(createUser); err != nil {
		u.log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := u.userService.CreateOne(createUser); err != nil {
		u.log.Errorf("创建用户失败：%s", err)
		return nil, response.InternalServerError.SetMsg("创建用户失败：%s", err)
	}

	u.log.Infof("创建用户成功: %s", createUser)
	return response.Success(createUser), nil
}

// 根据ID删除用户
func (u *User) Delete(c *gin.Context) (*response.Response, error) {
	u.log.Infof("删除用户")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.log.Errorf("参数绑定错误：%s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	if err := u.userService.DeleteOne(&domain.User{ID: id}); err != nil {
		u.log.Info("查询删除用户失败： ", err)
		return nil, response.InternalServerError.SetMsg("删除用户失败：%s", err)
	}

	u.log.Info("删除用户成功 ", id)
	return response.Success("删除成功"), nil
}

// 部分更新
func (u *User) Patch(c *gin.Context) (*response.Response, error) {
	u.log.Infof("更新用户信息")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.log.Errorf("参数错误： %d", id)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	user := &domain.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		u.log.Errorf("参数错误： %s", err)
		return nil, response.InvalidParams.SetMsg("参数错误: %s", err)
	}
	user.ID = id
	if err := u.userService.UpdateOne(user); err != nil {
		u.log.Errorf("更新失败：error： %s", err)
		return nil, response.InternalServerError.SetMsg("%s", err)
	}

	u.log.Infof("更新成功")
	return response.Success(user), nil
}

func (u *User) ListRole(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		return nil, response.InvalidParams.SetMsg("ID is required. ")
	}
	u.log.Infof("根据ID查询用户角色: ID[%d]", id)

	var roles []*domain.Role
	if err := u.userService.SelectRoles(&domain.User{ID: id}, &roles); err != nil {
		return nil, err
	}
	return response.Success(roles), nil
}

func (u *User) PostRole(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		u.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var requestRoles dto.AddUserRole
	if err := c.ShouldBindJSON(&requestRoles); err != nil {
		u.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	u.log.Infof("给用户: %s 添加角色：%s", id, requestRoles.Roles)
	if err := u.userService.InsertUserRoles(&domain.User{ID: id}, requestRoles.Roles); err != nil {
		return nil, err
	}

	return response.Success(requestRoles), nil
}


func (u *User) PutRole(c *gin.Context) (*response.Response, error) {
	// 修改用户角色
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		u.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	var requestRoles dto.AddUserRole
	if err := c.ShouldBindJSON(&requestRoles); err != nil {
		u.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	u.log.Infof("修改用户: %s 角色：%s", id, requestRoles.Roles)
	if err := u.userService.UpdateUserRoles(&domain.User{ID: id}, requestRoles.Roles); err != nil {
		return nil, err
	}
	return response.Success(requestRoles), nil
}

// 获取用户的菜单
func (u *User) ListMenus(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		u.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	u.log.Infof("获取用户：%d 所有菜单", id)

	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}
	if err := u.userService.SelectMenus(&p, &domain.User{ID: id}); err != nil {
		return nil, err
	}
	return response.Success(&p), nil
}

// 获取用户的博客
func (u *User) ListPosts(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		u.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	p := pager.Pager{}
	u.log.Infof("获取用户：%d 所有博客", id)
	posts := dto.ListPosts{}
	if err := c.ShouldBind(&posts); err != nil {
		return nil, err
	}

	posts.Visibility = 2
	posts.Status = 2

	if err := u.postService.SelectAll(&p, &posts); err != nil {
		return nil, err
	}
	return response.Success(&p), nil
}

// 获取用户的文件列表
func (u *User) ListFiles(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		u.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}
	u.log.Infof("获取用户：%d 所有文件", id)

	if err := u.fileService.SelectAll(&p, &domain.File{UserID: id}); err != nil {
		return nil, err
	}

	return response.Success(&p), nil
}

// 获取用户的专题
func (u *User) ListSubjects(c *gin.Context) (*response.Response, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		u.log.Errorf("参数绑定错误： %s", err)
		return nil, response.InvalidParams.SetMsg("%s", err)
	}

	p := pager.Pager{
		PageNo:   request.GetPageNo(c),
		PageSize: request.GetPageSize(c),
	}
	u.log.Infof("获取用户：%d 所有博客", id)

	if err := u.subjectService.SelectAll(c, &p, &dto.ListSubjects{UserId: id}); err != nil {
		return nil, err
	}

	return response.Success(&p), nil
}
