package impl

import (
	"blog/pkg/model/common"
)

type UserServiceImpl struct {
	common.BaseServiceImpl
}

//func (u *UserServiceImpl) ISelectOne(c *gin.Context, user *po.User) error {
//	userID, _ := c.Get(common.SessionUserIDKey)
//
//
//	global.DB.Model(&po.User{}).Where("user_id=?", userID)
//
//	return nil
//
//}
//
//func (u *UserServiceImpl) ISelectList(c *gin.Context, pager *vo.Pager, user *po.User) error {
//	panic("implement me")
//}
//
//func (u *UserServiceImpl) ICreateOne(c *gin.Context, user *po.User) error {
//	panic("implement me")
//}
//
//func (u *UserServiceImpl) IDeleteOne(c *gin.Context, user *po.User) error {
//	panic("implement me")
//}
//
//func (u *UserServiceImpl) IUpdateOne(c *gin.Context, user *po.User) error {
//	panic("implement me")
//}
