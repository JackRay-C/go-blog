package middleware

import (
	"blog/pkg/global"
	"blog/pkg/model/vo"
	token2 "blog/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	log := global.Log

	return func(c *gin.Context) {
		var token string
		if s, exist := c.GetQuery(global.RequestAccessTokenKey); exist {
			token = s
		} else {
			token = c.GetHeader(global.RequestAccessTokenKey)
		}

		if token == "" {
			c.Set(global.SessionIsLoginKey, false)
			c.Next()
			return
		} else {
			claim, err := token2.ParseAccessToken(token)
			if err != nil {
				log.Infof("failed to valied token: %s", token)
				c.AbortWithStatusJSON(http.StatusOK, vo.TokenError)
				return
				//switch err.(*gojwt.ValidationError).Errors {
				//case gojwt.ValidationErrorExpired:
				//	log.Error("Token超时! ")
				//	c.AbortWithStatusJSON(http.StatusOK, vo.TokenExpire)
				//	return
				//default:
				//	log.Error("Token验证错误! ")
				//	c.AbortWithStatusJSON(http.StatusOK, vo.TokenError)
				//	return
				//}
			}

			c.Set(global.SessionUserNameKey, claim.Username)
			c.Set(global.SessionUserIDKey, claim.UserId)
			c.Set(global.SessionIsLoginKey, true)
			c.Next()
		}
	}
}

