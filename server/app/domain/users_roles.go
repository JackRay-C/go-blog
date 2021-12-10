package domain

import (
	"blog/core/global"
	"encoding/json"
)

type UsersRoles struct {
	UserId int `json:"user_id" gorm:"type:int"`
	RoleId int `json:"role_id" gorm:"type:int"`
}

func (*UsersRoles) TableName() string {
	return "users_roles"
}

func (u *UsersRoles) String() string {
	marshal, _ := json.Marshal(u)
	return string(marshal)
}

func (ur *UsersRoles) Insert(urs []*UsersRoles) error  {
	return global.DB.Model(ur).CreateInBatches(urs, 1000).Error
}

func (ur *UsersRoles) DeleteByUser() error  {
	return global.DB.Model(ur).Where("user_id=?", ur.UserId).Delete(ur).Error
}

func (ur *UsersRoles) Delete() error  {
	return global.DB.Model(ur).Where("user_id=? and role_id=?", ur.UserId, ur.RoleId).Delete(ur).Error
}