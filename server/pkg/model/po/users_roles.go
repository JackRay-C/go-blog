package po

import "encoding/json"

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

