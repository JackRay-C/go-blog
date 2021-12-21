package domain

import (
	"blog/core/global"
	_ "crypto/sha256"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id" gorm:"type:int;primary_key;auto_increment"`
	Username  string         `json:"username"  gorm:"type:varchar(255);uniqueIndex"`
	Nickname  string         `json:"nickname"  gorm:"type:varchar(255);uniqueIndex"`
	Password  string         `json:"password"  gorm:"type:varchar(255)"`
	Active    int8           `json:"active" gorm:"type:tinyint;default:1;comment:'1激活，2锁定'"`
	Email     string         `json:"email" gorm:"type:varchar(255);uniqueIndex"`
	Avatar    int            `json:"avatar" gorm:"type:int"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
//
//func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
//	if u.Password != "" {
//		u.Password = encrypt.Sha256(u.Password)
//	}
//	return
//}
//
//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	u.Password = encrypt.Sha256(u.Password)
//	return
//}

func (*User) TableName() string {
	return "users"
}

func (u *User) String() string {
	marshal, _ := json.Marshal(u)
	return string(marshal)
}

//func (u *User) Select() error {
//	return global.DB.Model(u).Where(u).First(&u).Error
//}
//
//func (u *User) List(list *[]User, offset int, limit int) error {
//	return global.DB.Model(u).Where(u).Offset(offset).Limit(limit).Find(list).Error
//}
//
//func (u *User) Insert() error {
//	return global.DB.Model(u).Clauses(clause.OnConflict{DoNothing: true}).Create(u).Error
//}
//
////func (u *User) Save(db *gorm.DB) error {
////	return db.Save(u).Error
////}
//
//func (u *User) Update() error {
//	return global.DB.Model(u).Omit("Username").Updates(u).Error
//}
//
//func (u *User) Delete() error {
//	return global.DB.Model(u).Delete(u).Error
//}
//
//func (u *User) DeleteIds(ids []int) error {
//	return global.DB.Model(u).Delete(u, ids).Error
//}
//
//func (u *User) Count(count *int64) error {
//
//	return global.DB.Model(u).Where(u).Count(count).Error
//
//}
//
//func (u *User) CountByCondition(m map[string]interface{}) (count int64) {
//	global.DB.Model(u).Where(m).Limit(1).Count(&count)
//	return count
//}

func (u *User) CountRole(count *int64) error {
	return global.DB.Table("roles").Joins("left join users_roles as ur on ur.role_id=roles.id").Joins("left join users as u on ur.user_id=u.id").Where("u.id=?", u.ID).Count(count).Error
}

func (u *User) ListRoles(roles *[]*Role) error {
	return global.DB.Table("roles").Joins("left join users_roles as ur on ur.role_id=roles.id").Joins("left join users as u on ur.user_id=u.id").Where("u.id=?", u.ID).Find(roles).Error
}


func (u *User) CountMenus(count *int64) interface{} {
	return global.DB.Table("menus").Joins("left join roles_menus rr on menus.id=rr.menu_id").Where("rr.role_id in (?)", global.DB.Table("roles").Select("ID").Joins("left join users_roles ur on roles.id=ur.role_id").Where("user_id=?", u.ID)).Count(count).Error
}

func (u *User) ListMenus(menus *[]*Menu, offset int, limit int) interface{} {
	return global.DB.Table("menus").Joins("left join roles_menus rr on menus.id=rr.menu_id").Where("rr.role_id in (?)", global.DB.Table("roles").Select("ID").Joins("left join users_roles ur on roles.id=ur.role_id").Where("user_id=?", u.ID)).Offset(offset).Limit(limit).Find(menus).Error
}
