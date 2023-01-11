package console

import (
	"blog/pkg/model/common"
	"blog/pkg/model/dto"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"
	"blog/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
)

type Auth struct {
	userService service.AuthService
	service     common.BaseService
}

func NewAuth() *Auth {
	return &Auth{userService: service.NewAuthService(), service: &common.BaseServiceImpl{}}
}

func (a *Auth) Login(c *gin.Context) (*vo.Response, error) {
	loginForm := &dto.LoginBody{}

	if err := c.ShouldBindJSON(&loginForm); err != nil {
		return nil, vo.InvalidParams.SetMsg("登录失败: 参数错误：%s", err)
	}

	token, err := a.userService.ILogin(loginForm)
	if err != nil {
		return nil, vo.IncorrectUsernamePassword
	}

	return vo.Success(token), nil
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
	if refreshToken, ok := c.GetQuery("refresh_token"); ok {
		token, err := a.userService.IRefreshToken(refreshToken)
		if err != nil {
			return nil, err
		}
		return vo.Success(token), nil
	}
	return nil, vo.InvalidParams.SetMsg("refresh_token params required. ")
}

func (a Auth) Captcha(c *gin.Context) (*vo.Response, error) {
	return nil, nil
}

func (a *Auth) Info(c *gin.Context) (*vo.Response, error) {
	// 获取当前用户信息
	user := &po.User{}
	if accessToken, ok := c.GetQuery("access_token"); ok {
		if err:= a.userService.IInfo(c, accessToken, user); err != nil {
			return nil, vo.InternalServerError.SetMsg("%s", err)
		}
		log.Println(user)
		return vo.Success(user), nil
	} else if  c.GetHeader("token") != "" {
		if err:= a.userService.IInfo(c, c.GetHeader("token"), user); err != nil {
			return nil, vo.InternalServerError.SetMsg("%s", err)
		}
		log.Println(user)
		return vo.Success(user), nil
	}
	return nil, vo.InvalidParams.SetMsg("access_token params required. ")
}

func (a *Auth) Permissions(c *gin.Context) (*vo.Response, error) {
	permissions := make([]*po.Permissions, 0)
	if accessToken, ok := c.GetQuery("access_token"); ok {
		if err:= a.userService.IPermissions(c, accessToken, &permissions); err != nil {
			return nil, vo.InternalServerError.SetMsg("%s", err)
		}
		return vo.Success(permissions), nil
	}
	return nil, vo.InvalidParams.SetMsg("access_token params required. ")
}

func (a *Auth) Roles(c *gin.Context) (*vo.Response, error) {
	roles := make([]*po.Role, 0)
	if accessToken, ok := c.GetQuery("access_token"); ok {
		if err:= a.userService.IRoles(c, accessToken, &roles); err != nil {
			return nil, vo.InternalServerError.SetMsg("%s", err)
		}
		return vo.Success(roles), nil
	}
	return nil, vo.InvalidParams.SetMsg("access_token params required. ")
}

