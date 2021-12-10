package dto

import (
	"blog/app/domain"
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

type RequestPutUser struct {
	ID       int    `json:"username" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Avatar   string `json:"avatar"`
}

type PostUserRoles struct {
	Roles []*domain.Role `json:"roles"`
}