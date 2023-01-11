package middleware

import (
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"blog/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken string
		if s, exist := c.GetQuery(global.RequestQueryTokenKey); exist {
			accessToken = s
		} else {
			accessToken = c.GetHeader(global.RequestQueryTokenKey)
		}

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusOK, vo.NotLogin)
			//c.Next()
			return
		} else {
			// 验证token
			claim, err := token.ParseAccessToken(accessToken)
			if err != nil {
				global.Log.Infof("failed to valied token: %s", err)
				c.AbortWithStatusJSON(http.StatusOK, vo.TokenError.SetMsg("%s", err))
				return
			}

			// 验证通过之后延长token的时长
			if err := token.SetAccessTokenExpire(accessToken, global.App.Server.AccessTokenExpire); err != nil {
				c.AbortWithStatusJSON(http.StatusOK, vo.TokenError.SetMsg("%s", err))
				return
			}

			c.Set(global.SessionUserNameKey, claim.Username)
			c.Set(global.SessionUserIDKey, claim.UserId)
			c.Set(global.SessionIsLoginKey, true)
		}

		userRoleService := service.NewUsersRolesService()
		rolePermissionService := service.NewRolesPermissionService()

		var roles []*po.Role
		var permissions []*po.Permissions

		// 获取用户角色列表
		userId, _ := c.Get(global.SessionUserIDKey)
		if err := userRoleService.ISelectUserRoles(c, &po.User{ID: userId.(int64)}, &roles); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, vo.Forbidden)
			return
		}

		// 获取用户权限列表
		if err := rolePermissionService.ISelectPermissionByRoles(c, &permissions, roles...); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, vo.InternalServerError)
			return
		}

		c.Set(global.SessionRoleKey, roles)
		c.Set(global.SessionPermissionKey, permissions)

		c.Next()
	}
}
