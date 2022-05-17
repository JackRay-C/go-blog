package service

import (
	"blog/pkg/model/po"
	"blog/pkg/service/impl"
	"github.com/gin-gonic/gin"
)

type UsersRolesService interface {
	ISelectUserRoles(c *gin.Context, user *po.User, roles *[]*po.Role) error
	IUpdateUserRoles(c *gin.Context, user *po.User, roles *[]*po.Role) error
}

func NewUsersRolesService() UsersRolesService {
	return &impl.UsersRolesServiceImpl{}
}

