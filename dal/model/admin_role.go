package model

import (
	"server/dal/dbtypes"
)

type AdminRole struct {
	BaseModel
	Name            string              `json:"name" gorm:"unique"`
	Code            string              `json:"code" gorm:"unique"` // 角色编号
	Desc            string              `json:"desc"`
	PermissionCodes dbtypes.StringArray `json:"permission_codes" gorm:"type:longtext"` // 权限码
	Users           []*AdminUser        `json:"users" gorm:"many2many:admin_user_roles;"`
}

func (AdminRole) TableName() string {
	return "admin_role"
}
