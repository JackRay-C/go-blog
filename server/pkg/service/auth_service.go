package service

import (
	"blog/pkg/model/dto"
	"blog/pkg/model/po"
	"blog/pkg/service/impl"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	ILogin(body *dto.LoginBody) (token *dto.Token, err error)
	IRegister(body *dto.RegisterBody) error
	IRefreshToken(refreshToken string) (token *dto.Token, err error)
	ICaptcha() (string, error)
	IInfo(c *gin.Context, accessToken string, user *po.User) error
	IPermissions(c *gin.Context, accessToken string, permissions *[]*po.Permissions)error
	IRoles(c *gin.Context, accessToken string, roles *[]*po.Role) error
}

func NewAuthService() AuthService  {
	return &impl.AuthServiceImpl{}
}
