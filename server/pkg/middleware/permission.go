package middleware

import (
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoleService := service.NewUsersRolesService()
		rolePermissionService := service.NewRolesPermissionService()

		var roles []*po.Role
		var permissions []*po.Permissions

		isLogin, exists := c.Get(global.SessionIsLoginKey)
		if !exists || !isLogin.(bool) {
			c.AbortWithStatusJSON(http.StatusOK, vo.NotLogin)
			return
		}

		// 获取用户角色列表
		userId, _ := c.Get(global.SessionUserIDKey)
		if err := userRoleService.ISelectUserRoles(c, &po.User{ID: userId.(int)}, &roles); err != nil {
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
