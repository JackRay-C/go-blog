package domain

import (
	"blog/core/global"
	"encoding/json"
	"gorm.io/gorm/clause"
)

type RoleMenu struct {
	RoleId  int `json:"role_id" gorm:"type:int;index:idx_roleId_menuId,unique;"`
	MenuId int `json:"menu_id" gorm:"type:int;index:idx_roleId_menuId,unique;"`
}

func (r *RoleMenu) TableName() string {
	return "roles_menus"
}

func (r *RoleMenu) String() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (r *RoleMenu) Insert(routes ...*RoleMenu) error {
	return global.DB.Model(r).Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(routes, 1000).Error
}

func (r *RoleMenu) Delete() error {
	return global.DB.Model(r).Where("role_id=? and menu_id=?", r.RoleId, r.MenuId).Delete(r).Error
}