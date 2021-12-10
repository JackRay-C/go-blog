package dto

import "blog/app/domain"

type AddRoleMenus struct {
	Menus []*domain.Menu `json:"menus"`
}

type AddUserRole struct {
	Roles []*domain.Role `json:"roles"`
}


type AddRolePermission struct {
	Permissions []*domain.Permissions `json:"permissions"`
}
