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
		if s, exist := c.GetQuery(global.RequestQueryTokenKey); exist {
			token = s
		} else {
			token = c.GetHeader(global.RequestHeaderTokenKey)
		}

		if token == "" || c.Request.RequestURI=="/api/v1/auth/login"{
			c.Set(global.SessionIsLoginKey, false)
			c.Next()
			return
		} else {
			claim, err := token2.ParseAccessToken(token)
			if err != nil {
				log.Infof("failed to valied token: %s", token)
				c.AbortWithStatusJSON(http.StatusOK, vo.TokenError)
				return
			}

			c.Set(global.SessionUserNameKey, claim.Username)
			c.Set(global.SessionUserIDKey, claim.UserId)
			c.Set(global.SessionIsLoginKey, true)
			c.Next()
		}
	}
}

