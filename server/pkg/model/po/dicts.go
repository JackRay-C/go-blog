package po

import "encoding/json"

type Dict struct {
	ID          int    `json:"id" gorm:"type:int;primary_key;auto_increment;common:主键ID"`
	Name        string `json:"name" gorm:"type:varchar(255);index:idx_type_code,unique;"`
	Code        int    `json:"code" gorm:"type:int;index:idx_type_code,unique;"`
	Value       string `json:"value" gorm:"type:varchar(255);index:idx_type_code,unique;"`
	Description string `json:"description" gorm:"type:varchar(255)"`
}

func (d *Dict) TableName() string  {
	return "dicts"
}

func (d *Dict) String() string {
	marshal, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(marshal)
}
