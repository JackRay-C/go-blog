package dto

import "blog/pkg/model/po"

type PutUserRole struct {
	Roles []*po.Role `json:"roles"`
}


type PutRolePermission struct {
	Permissions []*po.Permissions `json:"permissions"`
}

