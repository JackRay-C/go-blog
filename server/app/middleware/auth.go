package middleware

import (
	"blog/app/jwt"
	"blog/app/response"
	"blog/core/global"
	gojwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	log := global.Logger

	return func(c *gin.Context) {
		var token string
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			c.Set("is_login", false)
			c.Next()
			return
		} else {
			claim, err := jwt.ParseToken(token)
			if err != nil {
				switch err.(*gojwt.ValidationError).Errors {
				case gojwt.ValidationErrorExpired:
					log.Error("Token超时! ")
					c.AbortWithStatusJSON(http.StatusOK, response.TokenExpire)
					return
				default:
					log.Error("Token验证错误! ")
					c.AbortWithStatusJSON(http.StatusOK, response.TokenError)
					return
				}
			}

			c.Set("current_user_name", claim.Username)
			c.Set("current_user_id", claim.UserId)
			c.Set("is_login", true)
			c.Next()
		}
	}
}

