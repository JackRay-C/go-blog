package impl

import (
	"blog/pkg/model/po"
	"github.com/gin-gonic/gin"
)

type UsersRolesServiceImpl struct {

}

func (u *UsersRolesServiceImpl) ISelectUserRoles(c *gin.Context, user *po.User, roles *[]*po.Role) error {
	panic("implement me")
}

func (u *UsersRolesServiceImpl) IUpdateUserRoles(c *gin.Context, user *po.User, roles *[]*po.Role) error {
	panic("implement me")
}

