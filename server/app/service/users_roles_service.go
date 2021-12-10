package service

import (
	"blog/app/domain"
	"blog/core/global"
	"blog/core/logger"
)

type UsersRolesService struct {
	log logger.Logger
}

func NewUsersRolesService() *UsersRolesService {
	return &UsersRolesService{
		log: global.Logger,
	}
}

// 根据用户id获取所有角色
func (urs UsersRolesService) SelectRolesByUserId(userId int, roles *[]*domain.Role) error {
	urs.log.Infof("根据用户ID【%d】获取所有角色", userId)
	if err := global.DB.Table("roles").Joins("left join users_roles as ur on roles.id=ur.role_id").Where("ur.user_id=?", userId).Find(&roles).Error;err!=nil {
		urs.log.Errorf("获取用户角色失败：%s", err)
		return err
	}

	return nil
}

// 根据角色ID获取所有用户
func (urs UsersRolesService) SelectUsersByRoleId(roleId int) (users []*domain.User, err error)  {
	return nil, err
}
