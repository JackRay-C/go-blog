package dto

import (
	"blog/pkg/model/po"
	"encoding/json"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *Login) String() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type Register struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (l *Register) String() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type RequestCreateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   string `json:"avatar"`
}

type PutUser struct {
	ID       int64    `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Active   int8   `json:"active"`
	Avatar   int64    `json:"avatar"`
}

type PostUserRoles struct {
	Roles []*po.Role `json:"roles"`
}
