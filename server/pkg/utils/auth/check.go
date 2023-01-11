package auth

import (
	"blog/pkg/global"
	"blog/pkg/model/po"
	"blog/pkg/utils/token"
	"github.com/gin-gonic/gin"
)

// CheckPermission 检查当前用户是否有权限
func CheckPermission(c *gin.Context, objectType string, actionType string) bool {
	permissions, exists := c.Get(global.SessionPermissionKey)
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
func CheckLogin(c *gin.Context) bool {
	// 1、获取header的access_token
	accessToken := c.GetHeader(global.RequestQueryTokenKey)

	if accessToken == "" {
		return false
	}

	if _, err := token.ParseAccessToken(accessToken); err != nil {
		return false
	}

	return true
}

// CheckAdmin 判断是否是管理员
func CheckAdmin(c *gin.Context) bool {
	roles, exists := c.Get(global.SessionRoleKey)
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

func GetCurrentUserId(c *gin.Context) int64 {
	userId, _ := c.Get(global.SessionUserIDKey)
	return userId.(int64)
}
