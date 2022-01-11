package web

import (
	"blog/app/domain"
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
	postService *service.PostService
	subjectService *service.SubjectService
}

func NewUser() *User {
	return &User{
		log: global.Logger,
		userService: service.NewUserService(),
		postService: service.NewPostService(),
		subjectService: service.NewSubjectService(),
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

func (u *User) GetRoles(c *gin.Context) (*response.Response, error)  {
	return response.Success("success"),nil
}
