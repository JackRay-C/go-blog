package service

import (
	"blog/app/domain"
	"blog/app/encrypt"
	"blog/app/jwt"
	"blog/app/model/dto"
	"blog/app/model/vo"
	"blog/app/pager"
	"blog/app/response"
	"blog/core/global"
	"blog/core/logger"
	"errors"
	"gorm.io/gorm"
)

type UserService struct {
	Log    logger.Logger
	engine *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		Log:    global.Logger,
		engine: global.DB,
	}
}

func (a *UserService) Auth(login *dto.Login) (*vo.VToken, error) {
	var user domain.User

	if err := global.DB.Model(&domain.User{}).Where("username=? or email=?", login.Username, login.Username).First(&user).Error; err != nil {
		a.Log.Errorf("登录失败：用户【%s】不存在", login.Username)
		return nil, response.IncorrectUsernamePassword
	}

	if user.Password != encrypt.Sha256(login.Password) {
		a.Log.Errorf("登录失败：密码错误 【%s:%s】", login.Username, login.Password)
		return nil, response.IncorrectUsernamePassword.SetMsg("密码错误. ")
	}

	if user.Active == 2 {
		a.Log.Errorf("用户【%s】已被锁定，请联系管理员解锁该用户. ", login.Username)
		return nil, response.AccountLocked.SetMsg("用户【%s】已被锁定，请联系管理员解锁该用户. ", login.Username)
	}

	// 4、生成token
	if token, err := jwt.GenerateToken(user.ID, user.Username); err != nil {
		a.Log.Errorf("登录失败：用户【%s】生成token失败. ", login.Username)
		return nil, response.FailedGenerateToken.SetMsg("登录失败：用户【%s】生成token失败. ", login.Username)
	} else {
		a.Log.Infof("用户【%s】登录成功.", login.Username)
		return &vo.VToken{Token: token}, nil
	}
}

func (u *UserService) DeleteOne(user *domain.User) error {
	if err := u.SelectOne(user); err != nil {
		return err
	}

	if err := global.DB.Model(&domain.User{}).Where("id=?", user.ID).Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserService) SelectOne(user *domain.User) error {
	if err := global.DB.Model(&domain.User{}).Where("id=?", user.ID).First(&user).Error;err != nil {
		return err
	}

	return nil
}

func (u *UserService) CreateOne(user *domain.User) error {
	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		var u1 domain.User
		err := tx.Model(&domain.User{}).Where("username=?", user.Username).Or("email=?", user.Email).Or("nickname=?", user.Nickname).First(&u1).Error


		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该用户已存在. ")
		}

		if err := tx.Model(&domain.User{}).Create(user).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}


	//if err := u1.Select(); err == gorm.ErrRecordNotFound {
	//	// 2、根据昵称判断，查询不到则判断是否存在相同邮箱
	//	u2 := &domain.User{Nickname: user.Nickname}
	//	if err := u2.Select(); err == gorm.ErrRecordNotFound {
	//		// 3、判断重复邮箱
	//		u3 := &domain.User{Email: user.Email}
	//		if err := u3.Select(); err == gorm.ErrRecordNotFound {
	//			// 创建用户
	//			if err := user.Insert(); err != nil {
	//				return response.DatabaseInsertError.SetMsg("%s", err)
	//			}
	//			// 查询用户
	//			if err := user.Select(); err != nil {
	//				return response.DatabaseSelectError.SetMsg("%s", err)
	//			}
	//		} else if err != nil {
	//			return response.DatabaseSelectError.SetMsg("%s", err)
	//		} else {
	//			return response.RecoreExisted.SetMsg("该邮箱已注册: %s", user.Email)
	//		}
	//	} else if err != nil {
	//		return response.DatabaseSelectError.SetMsg("%s", err)
	//	} else {
	//		return response.RecoreExisted.SetMsg("该昵称已存在: %s", user.Nickname)
	//	}
	//} else if err != nil {
	//	return response.DatabaseSelectError.SetMsg("%s", err)
	//} else {
	//	return response.RecoreExisted.SetMsg("该用户名已存在: %s", user.Username)
	//}

	return nil
}

func (u *UserService) SelectAll(p *pager.Pager, user *domain.User) error {
	var users []domain.User

	if err := user.Count(&p.TotalRows); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	p.PageCount = int((p.TotalRows + int64(p.PageSize) - 1) / int64(p.PageSize))
	p.List = &users
	if err := user.List(&users, (p.PageNo-1)*p.PageSize, p.PageSize); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	return nil
}

func (u *UserService) UpdateOne(user *domain.User) error {
	if err := user.Update(); err != nil {
		return response.DatabaseUpdateError.SetMsg("%s", err)
	}

	if err := user.Select(); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}
	return nil
}

func (u *UserService) SelectRoles(user *domain.User, roles *[]*domain.Role) error {
	if err := user.ListRoles(roles); err != nil {
		return err
	}
	return nil
}

func (u *UserService) InsertUserRoles(user *domain.User, roles []*domain.Role) error {
	if err := global.DB.Model(&domain.User{}).Where("id=?", user.ID).First(&user).Error; err != nil || err == gorm.ErrRecordNotFound {
		return err
	}

	var usersRoles []*domain.UsersRoles
	for _, role := range roles {
		usersRoles = append(usersRoles, &domain.UsersRoles{UserId: user.ID, RoleId: role.ID})
	}
	if err := global.DB.Model(&domain.UsersRoles{}).CreateInBatches(usersRoles, 1000).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserService) UpdateUserRoles(user *domain.User, roles []*domain.Role) error {
	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.User{}).Where("id=?", user.ID).First(&user).Error; err != nil || err == gorm.ErrRecordNotFound {
			return err
		}

		// 删除用户的角色
		if err := tx.Model(&domain.UsersRoles{}).Where("user_id=?", user.ID).Delete(&domain.UsersRoles{UserId: user.ID}).Error; err != nil {
			return err
		}

		var usersRoles []*domain.UsersRoles
		for _, role := range roles {
			usersRoles = append(usersRoles, &domain.UsersRoles{UserId: user.ID, RoleId: role.ID})
		}

		// 重新添加
		if err := tx.Model(&domain.UsersRoles{}).Create(usersRoles).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (u *UserService) SelectMenus(p *pager.Pager, user *domain.User) error {
	var menus []*domain.Menu

	if err := user.CountMenus(&p.TotalRows); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	p.PageCount = int((p.TotalRows + int64(p.PageSize) - 1) / int64(p.PageSize))
	p.List = &menus
	if err := user.ListMenus(&menus, (p.PageNo-1)*p.PageSize, p.PageSize); err != nil {
		return response.DatabaseSelectError.SetMsg("%s", err)
	}

	return nil
}

func (u *UserService) SelectPosts(p *pager.Pager, user *domain.User) error {
	return nil
}

func (u *UserService) SelectFiles(p *pager.Pager, user *domain.User) error {
	return nil
}

func (u *UserService) SelectOneById(id int) (vUser *vo.VUser,err  error) {
	var user *domain.User
	if err := global.DB.Model(&domain.User{}).Where("id=?", id).First(&user).Error; err != nil || err == gorm.ErrRecordNotFound {
		return nil, err
	}
	var file *domain.File
	if err:=global.DB.Model(&domain.File{}).Where("id=?", user.Avatar).First(&file).Error; err != nil {
		return nil, err
	}

	return &vo.VUser{
		ID:       user.ID,
		Username: user.Username,
		Nickname: user.Nickname,
		Active:   user.Active,
		Email:    user.Email,
		Avatar:   file,
		Created:  user.CreatedAt,
	}, nil
}

func sendActiveEmail(done chan int) {
	//token, err := jwt.GenerateToken(user.ID, user.Username)
	//if err != nil {
	//	return
	//}
	//subject := "无名万物博客激活邮件！"
	//body := `
	//<h1>请于一周内点击以下链接进行账号激活，否则将删除该账号</h1>
	//<a href="https://blog.renhj.org/register/active?token=%s">https://blog.renhj.org/register/active?token=%s</a>
	//`
	//to := []string{user.Email}
	//err = global.Email.SendMail(to, subject, fmt.Sprintf(body, token, token))
	//if err != nil {
	//	return
	//}
}
