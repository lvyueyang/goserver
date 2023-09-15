package model

import (
	"server/dal/dbtypes"
)

type AdminRole struct {
	BaseModel
	Name           string              `json:"name" gorm:"unique"`
	Desc           string              `json:"desc"`
	PermissionCode dbtypes.StringArray `gorm:"type:longtext"` // 权限码
}

func (AdminRole) TableName() string {
	return "admin_role"
}
