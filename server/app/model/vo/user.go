package vo

import (
	"blog/app/model/po"
	"encoding/json"
	"time"
)

type VUser struct {
	ID        int          `json:"id"`
	Username  string       `json:"username"`
	Nickname  string       `json:"nickname"`
	Active    int8         `json:"active"`
	Email     string       `json:"email"`
	Avatar    *po.File `json:"avatar"`
	CreatedAt time.Time    `json:"created_at"`
}

func (l *VUser) string() string {
	marshal, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type VToken struct {
	Token string `json:"token"`
}

type VUserInfo struct {
	ID          int                   `json:"id"`
	Username    string                `json:"username"`
	Nickname    string                `json:"nickname"`
	Active      int8                  `json:"active"`
	Email       string                `json:"email"`
	Avatar      *po.File          `json:"avatar"`
	Roles       []*po.Role        `json:"roles"`
	Permissions []*po.Permissions `json:"permissions"`
	CreatedAt   time.Time             `json:"created_at"`
}

func (v *VUserInfo) string() string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type UserInfo struct {
	ID          int            `json:"id"`
	Username    string         `json:"username"`
	Nickname    string         `json:"nickname"`
	Active      int8           `json:"active"`
	Email       string         `json:"email"`
	AvatarImage string         `json:"avatar_image"`
	Roles       []*po.Role `json:"roles"`
	Permissions []*po.Role `json:"permissions"`
	CreatedAt   time.Time      `json:"created_at"`
}
