package jwt

import (
	"blog/core/global"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type Claims struct {
	UserId   int `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (c *Claims) String() string {
	marshal, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func GetJWTSecret() []byte {
	return []byte(global.Setting.Jwt.Secret)
}

func GenerateToken(userId int, username string) (string, error) {
	id, err := global.Snowflake.NextID()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	now := time.Now()
	expire := now.Add(global.Setting.Jwt.Expire)
	claims := Claims{
		UserId:      userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Id:        strconv.FormatInt(id, 10),
			Issuer:    global.Setting.Jwt.Issuer,
		},
	}
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return withClaims.SignedString(GetJWTSecret())
}

func ParseToken(token string) (*Claims, error) {
	withClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if withClaims != nil {
		if claims, ok := withClaims.Claims.(*Claims); ok && withClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
