package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Meta struct {
	Title       string `json:"title"`
	Layout      string `json:"layout"`
	RequireAuth bool   `json:"require_auth"`
}

func (m *Meta) String() string {
	marshal, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (m *Meta) Value() (driver.Value, error) {
	marshal, err := json.Marshal(m)
	return string(marshal), err
}

func (m *Meta) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := Meta{}
	err := json.Unmarshal(bytes, &result)
	*m = result
	return err
}

type Menu struct {
	ID        int            `json:"id" gorm:"type:int;primary_key;auto_increment;comment:主键ID;"`
	Name      string         `json:"name" gorm:"type:varchar(100);index:idx_menu,unique;"`
	Path      string         `json:"path" gorm:"type:varchar(100);index:idx_menu,unique;"`
	Component string         `json:"component" gorm:"type:varchar(255);"`
	Meta      *Meta          `json:"meta" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (m *Menu) String() string {
	marshal, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (*Menu) TableName() string {
	return "menus"
}