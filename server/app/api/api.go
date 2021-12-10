package api

import (
	"blog/app/domain"
	"blog/app/response"
	"github.com/gin-gonic/gin"
)

type API interface {
	Get(c *gin.Context) (*response.Response, error)    //根据ID获取
	List(c *gin.Context) (*response.Response, error)   //分页获取所有
	Post(c *gin.Context) (*response.Response, error)   // 创建
	Delete(c *gin.Context) (*response.Response, error) // 根据id删除
	Patch(c *gin.Context) (*response.Response,error)  // 根据id更新
	Put(c *gin.Context) (*response.Response, error)    // 根据id全量更新
}

func CheckPermission(c *gin.Context, objectType string, actionType string) bool {
	permissions, exists := c.Get("current_user_permissions")
	if !exists {
		return false
	}
	for _, permission := range permissions.([]*domain.Permissions) {
		if permission.ObjectType == objectType && permission.ActionType == actionType {
			return true
		}
	}
	return false
}

func CheckLogin(c *gin.Context) bool  {
	isLogin, exists := c.Get("is_login")
	if !exists || !isLogin.(bool) {
		return false
	}
	return true
}

// 判断是否是管理员
func CheckAdmin(c *gin.Context) bool  {
	roles, exists := c.Get("current_user_roles")
	if exists {
		r := roles.([]*domain.Role)
		for _, role := range r {
			if role.Name == "Admin" {
				return true
			}
		}
	}
	return false
}