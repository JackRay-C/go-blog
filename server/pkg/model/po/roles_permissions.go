package po

import "encoding/json"

type RolesPermissions struct {
	RoleId       int64 `json:"role_id" gorm:"type:int;index:idx_roleId_permissionId,unique;"`
	PermissionId int64 `json:"permission_id" gorm:"type:int;index:idx_roleId_permissionId,unique;"`
}

func (p *RolesPermissions) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func (p *RolesPermissions) TableName() string {
	return "roles_permissions"
}