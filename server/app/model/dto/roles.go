package dto

import "blog/app/domain"


type PutUserRole struct {
	Roles []*domain.Role `json:"roles"`
}


type PutRolePermission struct {
	Permissions []*domain.Permissions `json:"permissions"`
}


type PutRoleMenus struct {
	Menus []*domain.Menu `json:"menus"`
}