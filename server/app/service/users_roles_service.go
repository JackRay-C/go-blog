package service

import (
	"blog/app/domain"
	"blog/core/global"
	"blog/core/logger"
	"errors"
	"gorm.io/gorm"
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
func (urs *UsersRolesService) SelectUserRoles(user *domain.User, roles *[]*domain.Role) error {
	// 1、查询用户是否存在
	err := global.DB.Model(&domain.User{}).Where("id=?", user.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该用户不存在. ")
	}

	return global.DB.Table("roles").
		Joins("left join users_roles as ur on roles.id=ur.role_id").
		Where("ur.user_id=?", user.ID).Find(&roles).Error
}

// 更新用户角色
func (urs *UsersRolesService) UpdateUserRoles(user *domain.User, roles []*domain.Role) error {
	// 1、查询用户是否存在
	err := global.DB.Model(&domain.User{}).Where("id=?", user.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该用户不存在. ")
	}

	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 1、删除用户的所有角色
		tx = tx.Model(&domain.UsersRoles{}).Where("user_id=?", user.ID)
		if err := tx.Delete(user).Error; err != nil {
			return err
		}

		// 2、添加用户角色
		var userRoles []*domain.UsersRoles
		for _, role := range roles {
			userRoles = append(userRoles, &domain.UsersRoles{UserId: user.ID, RoleId: role.ID})
		}

		return tx.Create(userRoles).Error
	})
}
