package auth

import (
	"blog/app/model/po"
	"github.com/gin-gonic/gin"
)


// CheckPermission 检查当前用户是否有权限
func CheckPermission(c *gin.Context, objectType string, actionType string) bool {
	permissions, exists := c.Get("current_user_permissions")
	if !exists {
		return false
	}
	for _, permission := range permissions.([]*po.Permissions) {
		if permission.ObjectType == objectType && permission.ActionType == actionType {
			return true
		}
	}
	return false
}

// CheckLogin 判断是否登录
func CheckLogin(c *gin.Context) bool  {
	isLogin, exists := c.Get("is_login")
	if !exists || !isLogin.(bool) {
		return false
	}
	return true
}

// CheckAdmin 判断是否是管理员
func CheckAdmin(c *gin.Context) bool  {
	roles, exists := c.Get("current_user_roles")
	if exists {
		r := roles.([]*po.Role)
		for _, role := range r {
			if role.Name == "Admin" {
				return true
			}
		}
	}
	return false
}