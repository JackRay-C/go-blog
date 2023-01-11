package impl

import (
	"blog/pkg/global"
	"blog/pkg/model/po"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersRolesServiceImpl struct {

}

func (u *UsersRolesServiceImpl) ISelectUserRoles(c *gin.Context, user *po.User, roles *[]*po.Role) error {
	// 获取用户角色

	// 1、查询用户是否存在
	err := global.DB.Model(&po.User{}).Where("id=?", user.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该用户不存在. ")
	}

	return global.DB.Table("roles").
		Joins("left join users_roles as ur on roles.id=ur.role_id").
		Where("ur.user_id=?", user.ID).Find(&roles).Error
}

func (u *UsersRolesServiceImpl) IUpdateUserRoles(c *gin.Context, user *po.User, roles *[]*po.Role) error {
	panic("implement me")
}

