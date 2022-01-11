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
	"time"
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
		return nil, response.AccountLocked.SetMsg("用户【%s】已被锁定，请联系管理员解锁该用户. ", login.Username)
	}

	// 4、生成token
	if token, err := jwt.GenerateToken(user.ID, user.Username); err != nil {
		return nil, response.FailedGenerateToken.SetMsg("登录失败：用户【%s】生成token失败. ", login.Username)
	} else {
		a.Log.Infof("用户【%s】登录成功.", login.Username)
		return &vo.VToken{Token: token}, nil
	}
}

func (u *UserService) DeleteOne(user *domain.User) error {
	if err := global.DB.Model(&domain.User{}).Where("id=?", user.ID).First(&user).Error; err == gorm.ErrRecordNotFound {
		return errors.New("该用户不存在. ")
	} else if err != nil {
		return err
	}

	if err := global.DB.Model(&domain.User{}).Where("id=?", user.ID).Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserService) SelectOne(user *domain.User) (*vo.VUser, error) {
	db := global.DB.Model(&domain.User{})

	if user.Active != 0 {
		db.Where("active=?", user.Active)
	}

	if err := db.Where("id=?", user.ID).First(&user).Error; err != nil {
		return nil, err
	}

	var file *domain.File
	if err := global.DB.Model(&domain.File{}).Where("id=?", user.Avatar).First(&file).Error; err != nil {
		return nil, err
	}

	return &vo.VUser{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Active:    user.Active,
		Email:     user.Email,
		Avatar:    file,
		CreatedAt: user.CreatedAt,
	}, nil

}

func (u *UserService) CreateOne(user *domain.User) error {
	if err := global.DB.Transaction(func(tx *gorm.DB) error {
		var u1 *domain.User
		err := tx.Model(&domain.User{}).Where("username=?", user.Username).First(&u1).Error

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该用户名已存在. ")
		}

		u1 = nil
		err = tx.Model(&domain.User{}).Where("email=?", user.Email).First(&u1).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该邮箱已经注册. ")
		}

		u1 = nil
		err = tx.Model(&domain.User{}).Where("nickname=?", user.Nickname).First(&u1).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该昵称已经存在. ")
		}

		// 密码加密
		user.Password = encrypt.Sha256(user.Password)
		if err := tx.Model(&domain.User{}).Create(user).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (u *UserService) SelectAll(p *pager.Pager, user *domain.User) error {
	var users []*domain.User
	var voUsers []*vo.VUser
	offset := (p.PageNo - 1) * p.PageSize
	limit := p.PageSize

	db := global.DB.Model(&domain.User{})
	if user.Active != 0 {
		db.Where("active=?", user.Active)
	}

	if err := db.Count(&p.TotalRows).Error; err != nil {
		return err
	}

	if err := db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		var userAvatar *domain.File
		if err := global.DB.Model(&domain.File{ID: user.Avatar}).First(&userAvatar).Error; err != nil {
			return err
		}
		voUsers = append(voUsers, &vo.VUser{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Active:    user.Active,
			Email:     user.Email,
			Avatar:    userAvatar,
			CreatedAt: user.CreatedAt,
		})
	}

	if p.TotalRows == 0 {
		p.PageCount = 0
		p.List = make([]string, 0)
	} else {
		p.PageCount = int((p.TotalRows + int64(p.PageSize) - 1) / int64(p.PageSize))
		p.List = &voUsers
	}

	return nil
}

func (u *UserService) UpdateOne(param *dto.PutUser) (*vo.VUser, error) {
	var user *domain.User
	err := global.DB.Model(&domain.User{}).Where("id=? and username=?", param.ID, param.Username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该用户不存在. ")
	}
	if err != nil {
		return nil, err
	}

	user = &domain.User{
		ID:        param.ID,
		Username:  param.Username,
		Nickname:  param.Nickname,
		Password:  encrypt.Sha256(param.Password),
		Active:    param.Active,
		Email:     param.Email,
		Avatar:    param.Avatar,
		UpdatedAt: time.Now(),
	}

	if err := global.DB.Model(&domain.User{}).Where("id=? and username=?", param.ID, param.Username).Omit("id", "username").Updates(user).Error; err != nil {
		return nil, err
	}

	return u.SelectOne(&domain.User{ID: user.ID})
}

//func (u *UserService) SelectUserRoles(user *domain.User, roles *[]*domain.Role) error {
//	return global.DB.Table("roles").Joins("left join users_roles as ur on ur.role_id=roles.id").Joins("left join users as u on ur.user_id=u.id").Where("u.id=?", user.ID).Find(roles).Error
//}
//
//func (u *UserService) InsertUserRoles(user *domain.User, roles []*domain.Role) error {
//	if err := global.DB.Model(&domain.User{}).Where("id=?", user.ID).First(&user).Error; err != nil || err == gorm.ErrRecordNotFound {
//		return err
//	}
//
//	var usersRoles []*domain.UsersRoles
//	for _, role := range roles {
//		usersRoles = append(usersRoles, &domain.UsersRoles{UserId: user.ID, RoleId: role.ID})
//	}
//	if err := global.DB.Model(&domain.UsersRoles{}).CreateInBatches(usersRoles, 1000).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (u *UserService) UpdateUserRoles(user *domain.User, roles []*domain.Role) error {
//	if err := global.DB.Transaction(func(tx *gorm.DB) error {
//		if err := tx.Model(&domain.User{}).Where("id=?", user.ID).First(&user).Error; err != nil || err == gorm.ErrRecordNotFound {
//			return err
//		}
//
//		// 删除用户的角色
//		if err := tx.Model(&domain.UsersRoles{}).Where("user_id=?", user.ID).Delete(&domain.UsersRoles{UserId: user.ID}).Error; err != nil {
//			return err
//		}
//
//		var usersRoles []*domain.UsersRoles
//		for _, role := range roles {
//			usersRoles = append(usersRoles, &domain.UsersRoles{UserId: user.ID, RoleId: role.ID})
//		}
//
//		// 重新添加
//		if err := tx.Model(&domain.UsersRoles{}).Create(usersRoles).Error; err != nil {
//			return err
//		}
//
//		return nil
//	}); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (u *UserService) SelectMenus(p *pager.Pager, user *domain.User) error {
//	var menus []*domain.Menu
//
//	if err := user.CountMenus(&p.TotalRows); err != nil {
//		return response.DatabaseSelectError.SetMsg("%s", err)
//	}
//
//	p.PageCount = int((p.TotalRows + int64(p.PageSize) - 1) / int64(p.PageSize))
//	p.List = &menus
//	if err := user.ListMenus(&menus, (p.PageNo-1)*p.PageSize, p.PageSize); err != nil {
//		return response.DatabaseSelectError.SetMsg("%s", err)
//	}
//
//	return nil
//}
//
//func (u *UserService) SelectPosts(p *pager.Pager, user *domain.User) error {
//	return nil
//}
//
//func (u *UserService) SelectFiles(p *pager.Pager, user *domain.User) error {
//	return nil
//}

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
