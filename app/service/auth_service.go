package service

import (
	"blog/app/service/impl"
	"blog/pkg/model/dto"
)

type AuthService interface {
	ILogin(body *dto.LoginBody) (token *dto.Token, err error)
	IRegister(body *dto.RegisterBody) error
	IRefreshToken(refreshToken string) (token *dto.Token, err error)
	ICaptcha() (string, error)
}

func NewAuthService() AuthService  {
	return &impl.AuthServiceImpl{}
}
