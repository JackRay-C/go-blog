package domain

import (
	"blog/core/global"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type Role struct {
	ID          int          `json:"id" gorm:"type:int;primary_key;auto_increment;comment:'主键ID'"`
	Name        string         `json:"name" gorm:"type:varchar(100);index:idx_name,unique;"`
	Description string         `json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeleteAt    gorm.DeletedAt `json:"delete_at"`
}


func (i *Role) TableName() string {
	return "roles"
}

func (i *Role) String() string {
	marshal, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(marshal)
}


func (i *Role) Insert() error {
	return global.DB.Model(i).Create(i).Error
}

func (i *Role) InsertAll(roles ...*Role) error {
	return global.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(roles, 1000).Error
}

func (i *Role) Save() error {
	return global.DB.Model(i).Omit("created_at").Save(i).Error
}

func (i *Role) Update() error {
	return global.DB.Model(i).Updates(i).Error
}

func (i *Role) Delete() error {
	return global.DB.Where("ID=?", i.ID).Delete(i).Error
}

func (i *Role) DeleteIds(ids []int) error {
	return global.DB.Delete(i, ids).Error
}

func (i *Role) Count(count *int64) error {
	return global.DB.Unscoped().Model(i).Where(i).Count(count).Error
}


func (i *Role) MenusCount(count *int64) error {
	return global.DB.Unscoped().Table("menus").Joins("left join roles_menus on menus.id=roles_menus.menu_id").Joins("left join roles on roles_menus.role_id=roles.id").Where("roles.id=?", i.ID).Count(count).Error
}

func (i *Role) ListMenus(list *[]Menu, offset int, limit int) error {
	return global.DB.Unscoped().Table("menus").Joins("left join roles_menus on menus.id=roles_menus.menu_id ").Joins("left join roles on roles_menus.role_id=roles.id").Where("roles.id=?", i.ID).Offset(offset).Limit(limit).Find(list).Error
}



