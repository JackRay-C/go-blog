package middleware

import (
	"blog/app/response"
	"blog/app/service"
	"blog/pkg/model/po"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoleService := service.NewUsersRolesService()
		rolePermissionService := service.NewRolesPermissionService()

		var roles []*po.Role
		var permissions []*po.Permissions

		isLogin, exists := c.Get("is_login")
		if !exists || !isLogin.(bool) {
			c.AbortWithStatusJSON(http.StatusOK, response.NotLogin)
			return
		}

		// 获取用户角色列表
		userId, _ := c.Get("current_user_id")
		if err := userRoleService.SelectUserRoles(&po.User{ID: userId.(int)}, &roles); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, response.Forbidden)
			return
		}

		// 获取用户权限列表
		var roleIds []int
		for _, role := range roles {
			roleIds = append(roleIds, role.ID)
		}

		if err := rolePermissionService.SelectPermissionByRoleId(&permissions, roleIds...); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, response.InternalServerError)
			return
		}

		c.Set("current_user_roles", roles)
		c.Set("current_user_permissions", permissions)

		c.Next()
	}
}
