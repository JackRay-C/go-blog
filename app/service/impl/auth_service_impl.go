package impl

import (
	"blog/app/encrypt"
	"blog/pkg/global"
	"blog/pkg/model/dto"
	"blog/pkg/model/po"
	"blog/pkg/model/vo"

	"blog/app/utils/token"

	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AuthServiceImpl struct {
}

func (a *AuthServiceImpl) ILogin(body *dto.LoginBody) (*dto.Token, error) {
	var user po.User

	// 1、Check whether the user name or password is empty
	if strings.Trim(body.Username, " ") == "" || strings.Trim(body.Password, " ") == "" {
		return nil, vo.RequiredUsernamePassword
	}

	// 2、Check whether the user exists in the database according to the user name or mailbox
	if err := global.DB.Model(&po.User{}).Where("username=? or email=?", body.Username, body.Username).First(&user).Error; err != nil {
		global.Log.Errorf("username not exist: %v", body)
		return nil, vo.IncorrectUsernamePassword
	}

	// 3、Check password whether correct
	if user.Password != encrypt.Sha256(body.Password) {
		global.Log.Errorf("password incorrect：%v", body)
		return nil, vo.IncorrectUsernamePassword
	}

	// 4、Check user whether locked
	if user.Active == 2 {
		global.Log.Errorf("user has locked：%v", body)
		return nil, vo.Locked
	}

	// 4、generate token
	//t, err := jwt.GenerateToken(user.ID, user.Username)
	accessToken, err := token.GenerateAccessToken(&token.Claims{UserId: user.ID, Username: user.Username})
	if err != nil {
		global.Log.Errorf("generate access token failed: %s", err)
		return nil, vo.FailedGenerateToken
	}
	// 设置accessToken过期时间
	if err := token.SetAccessTokenExpire(accessToken, global.App.Server.AccessTokenExpire); err != nil {
		global.Log.Errorf("set access token expire failed: %s", err)
		return nil, vo.FailedGenerateToken
	}

	refreshToken, err := token.GenerateRefreshToken(&token.Claims{UserId: user.ID, Username: user.Username})
	if err != nil {
		global.Log.Errorf("generate refresh token failed: %s", err)
		return nil, vo.FailedGenerateToken
	}
	// 设置refreshToken过期时间
	if err := token.SetRefreshTokenExpire(refreshToken, global.App.Server.RefreshTokenExpire); err != nil {
		global.Log.Errorf("set refresh token expire failed: %s", err)
		return nil, vo.FailedGenerateToken
	}

	t2 := &dto.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expire:       1000,
	}
	global.Log.Infof("login success: %s", body)
	return t2, nil
}

func (a *AuthServiceImpl) IRegister(register *dto.RegisterBody) error {
	// 1、设置user信息
	user := po.User{
		Username:  register.Username,
		Nickname:  register.Nickname,
		Password:  encrypt.Sha256(register.Password),
		Active:    1,
		Email:     register.Email,
		Avatar:    1,
		CreatedAt: time.Now(),
	}

	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		var u1 *po.User
		err := tx.Model(&po.User{}).Where("username=?", user.Username).First(&u1).Error

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该用户名已存在. ")
		}

		u1 = nil
		err = tx.Model(&po.User{}).Where("email=?", user.Email).First(&u1).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该邮箱已经注册. ")
		}

		u1 = nil
		err = tx.Model(&po.User{}).Where("nickname=?", user.Nickname).First(&u1).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该昵称已经存在. ")
		}

		// 密码加密
		user.Password = encrypt.Sha256(user.Password)
		if err := tx.Model(&po.User{}).Create(user).Error; err != nil {
			return err
		}

		// 设置默认用户角色为Editor
		userRole := &po.UsersRoles{UserId: user.ID, RoleId: po.ROLE_EDITOR}
		if err := global.DB.Model(&po.UsersRoles{}).Create(userRole).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return vo.CreateUserFailed
	}
	return nil
}

// IRefreshToken 根据refreshToken刷新accessToken，并重置refreshToken过期时间
func (a *AuthServiceImpl) IRefreshToken(refreshToken string) (*dto.Token, error) {
	// 1、判断refreshToken是否过期
	claims, err := token.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, vo.TokenExpire
	}

	// 2、重新生成accessToken
	accessToken, err := token.GenerateAccessToken(claims)
	if err != nil {
		return nil, vo.FailedGenerateToken
	}

	// 3、返回token 并设置refreshToken过期时间
	if err := token.SetRefreshTokenExpire(refreshToken, global.App.Server.RefreshTokenExpire); err != nil {
		return nil, vo.FailedGenerateToken
	}
	return &dto.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expire:       global.App.Server.AccessTokenExpire,
	}, nil
}

func (a *AuthServiceImpl) ICaptcha() (string, error) {
	panic("implement me")
}
