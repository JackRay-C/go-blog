package web

import (
	"blog/internal/logger"
	"blog/pkg/global"
	"blog/pkg/model/common"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
	log logger.Logger
	service common.BaseService
	userRoleService service.UsersRolesService
	rolePermissionService service.RolePermissionService
}

func NewUser() *User {
	return &User{
		log: global.Log,
		service: &common.BaseServiceImpl{},
		userRoleService: service.NewUsersRolesService(),
		rolePermissionService: service.NewRolesPermissionService(),
	}
}

// Get 根据ID获取用户信息
func (u *User) Get(c *gin.Context) (*vo.Response, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return nil, vo.InvalidParams.SetMsg("ID is required. ")
	}

	user := &po.User{ID: id}
	if err := u.service.ISelectOneWeb(c, user); err != nil {
		u.log.Errorf("根据ID查询用户失败 : %s", err)
		return nil, vo.InternalServerError.SetMsg("%s", err)
	} else {
		u.log.Infof("根据ID查询用户成功: %s", user)
		return vo.Success(user), nil
	}
}
