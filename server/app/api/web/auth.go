package web

import (
	"blog/app/domain"
	"blog/app/encrypt"
	"blog/app/model/dto"
	"blog/app/response"
	"blog/app/service"
	"blog/core/global"
	"blog/core/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type Auth struct {
	Log         logger.Logger
	userService *service.UserService
}

func NewAuth() Auth {
	return Auth{Log: global.Logger, userService: service.NewUserService()}
}

func (a Auth) Login(c *gin.Context) (*response.Response, error) {
	var loginForm dto.Login

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		a.Log.Errorf("参数绑定错误: %s", err)
		return nil, response.InvalidParams.SetMsg("登录失败: 参数错误：%s", err)
	}

	if token, err := a.userService.Auth(&loginForm); err != nil {
		a.Log.Errorf("用户%s认证失败: %s", loginForm.Username, err)
		return nil, response.InternalServerError.SetMsg("%s",err)
	} else {
		a.Log.Infof("登录成功： %s", loginForm.Username)
		return response.Success(token), nil
	}
}

func (a Auth) Register(c *gin.Context) (*response.Response, error) {
	var register dto.Register

	if err := c.ShouldBindJSON(&register); err != nil {
		a.Log.Errorf("参数绑定校验错误：%v", err)
		return nil, response.InvalidParams.SetMsg("参数绑定错误：%s", err)
	}

	user := domain.User{
		Username:  register.Username,
		Nickname:  register.Nickname,
		Password:  encrypt.Sha256(register.Password),
		Active:    1,
		Email:     register.Email,
		Avatar:    1,
		CreatedAt: time.Now(),
	}


	if err := service.NewUserService().CreateOne(&user); err != nil {
		a.Log.Errorf("注册失败：%s", err)
		return nil, response.InternalServerError.SetMsg("%s",err)
	}
	a.Log.Infof("注册成功: %s ", user.String())
	return response.Success(user), nil
}

func (a Auth) RefreshToken(c *gin.Context) (*response.Response, error) {
	return nil, nil
}

func (a Auth) Captcha(c *gin.Context) (*response.Response, error) {
	return nil, nil
}
