package service

import (
	"blog/pkg/model/common"
	"blog/pkg/service/impl"
)

type UserService interface {
	//ISelectOne(c *gin.Context, user *po.User) error
	//ISelectList(c *gin.Context, pager *vo.Pager, user *po.User) error
	//ICreateOne(c *gin.Context, user *po.User) error
	//IDeleteOne(c *gin.Context, user *po.User) error
	//IUpdateOne(c *gin.Context, user *po.User) error
	common.BaseService
}

func NewUserService() UserService {
	return &impl.UserServiceImpl{}
}
