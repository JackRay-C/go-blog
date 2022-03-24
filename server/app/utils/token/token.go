package token

import (
	"blog/core/global"
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Interface interface {
	GenerateToken(claims *Claims) (string, error)
	ParseToken(token string) (string, error)
}

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

func (c *Claims) String() string {
	marshal, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func GenerateAccessToken(c *Claims) (string, error) {
	// 1、生成token
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	// 2、存储到cache中
	if err := global.Cache.Set("accessToken:"+newUUID.String(), c.String()); err != nil {
		return "", err
	}

	// 3、返回
	return newUUID.String(), nil
}

func GenerateRefreshToken(c *Claims) (string, error) {
	// 1、生成token
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	// 2、存储到cache中
	if err := global.Cache.Set("refreshToken:"+newUUID.String(), c.String()); err != nil {
		return "", err
	}

	// 3、返回
	return newUUID.String(), nil
}

func SetAccessTokenExpire(token string, expire time.Duration) error {
	if err := global.Cache.SetExpire("accessToken:"+token, expire); err != nil {
		return err
	}
	return nil
}

func SetRefreshTokenExpire(token string, expire time.Duration) error {
	if err := global.Cache.SetExpire("refreshToken:"+token, expire); err != nil {
		return err
	}
	return nil
}

func ParseAccessToken(token string) (*Claims, error) {
	get, err := global.Cache.Get("accessToken" + token)
	if err != nil {
		return nil, err
	}

	var claims *Claims
	err = json.Unmarshal([]byte(get), claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func ParseRefreshToken(token string) (*Claims, error) {
	get, err := global.Cache.Get("refreshToken:" + token)
	if err != nil {
		return nil, err
	}
	var claims *Claims
	err = json.Unmarshal([]byte(get), claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
