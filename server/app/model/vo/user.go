package vo

import (
	"blog/app/domain"
	"encoding/json"
	"time"
)

type VUser struct {
	ID       int          `json:"id"`
	Username string       `json:"username"`
	Nickname string       `json:"nickname"`
	Active   int8         `json:"active"`
	Email    string       `json:"email"`
	Avatar   *domain.File `json:"avatar"`
	Created  time.Time    `json:"created"`
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
