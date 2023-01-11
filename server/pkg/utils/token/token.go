package token

import (
	"blog/pkg/global"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Interface interface {
	GenerateToken(claims *Claims) (string, error)
	ParseToken(token string) (string, error)
}

type Claims struct {
	UserId   int64    `json:"user_id"`
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

	uid := strings.ReplaceAll(newUUID.String(), "-", "")
	// 2、存储到cache中
	if err := global.Cache.Set(context.Background(), "accessToken:"+uid, c.String(), global.App.Server.AccessTokenExpire).Err(); err != nil {
		return "", err
	}

	// 3、返回
	return uid, nil
}

func GenerateRefreshToken(c *Claims) (string, error) {
	// 1、生成token
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	uid := strings.ReplaceAll(newUUID.String(), "-", "")
	// 2、存储到cache中
	if err := global.Cache.Set(context.Background(),"refreshToken:"+uid, c.String(), global.App.Server.RefreshTokenExpire).Err(); err != nil {
		return "", err
	}

	// 3、返回
	return uid, nil
}

func SetAccessTokenExpire(token string, expire time.Duration) error {
	return global.Cache.Expire(context.Background(), "accessToken:"+token, expire).Err()
}

func SetRefreshTokenExpire(token string, expire time.Duration) error {
	return global.Cache.Expire(context.Background(), "refreshToken:"+token, expire).Err()
}

func ParseAccessToken(token string) (*Claims, error) {
	get, err := global.Cache.Get(context.Background(), "accessToken:" + token).Result()

	if err == redis.Nil {
		return nil, errors.New("token expire. ")
	}
	if err != nil {
		return nil, err
	}

	claims :=  &Claims{}
	err = json.Unmarshal([]byte(get), claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func ParseRefreshToken(token string) (*Claims, error) {
	get, err := global.Cache.Get(context.Background(), "refreshToken:" + token).Result()
	if err != nil {
		return nil, err
	}
	claims:= &Claims{}

	if err := json.Unmarshal([]byte(get), claims);err != nil {
		return nil, err
	}

	return claims, nil
}
