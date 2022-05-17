package console

import (

	"blog/pkg/model/dto"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	userService service.AuthService
}

func NewAuth() *Auth {
	return &Auth{userService: service.NewAuthService()}
}

func (a *Auth) Login(c *gin.Context) (*vo.Response, error) {
	var loginForm *dto.LoginBody

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		return nil, vo.InvalidParams.SetMsg("登录失败: 参数错误：%s", err)
	}

	token, err := a.userService.ILogin(loginForm)
	if err != nil {
		return nil, err
	}

	return vo.Success(token),nil
}

func (a *Auth) Register(c *gin.Context) (*vo.Response, error) {
	var register dto.RegisterBody

	if err := c.ShouldBindJSON(&register); err != nil {
		return nil, vo.InvalidParams.SetMsg("参数绑定错误：%s", err)
	}


	if err := a.userService.IRegister(&register); err != nil {
		return nil, err
	}

	return vo.Success("register success"), nil
}

func (a *Auth) RefreshToken(c *gin.Context) (*vo.Response, error) {
	refreshToken := c.Request.Header.Get("refresh_token")
	token, err := a.userService.IRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return vo.Success(token), nil
}

func (a Auth) Captcha(c *gin.Context) (*vo.Response, error) {
	return nil, nil
}
