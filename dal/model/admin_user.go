package model

import (
	"server/consts"
)

type AdminUser struct {
	BaseModel
	Name     string                 `json:"name"`
	Username string                 `json:"user_name" gorm:"unique"`
	Password string                 `json:"_"`
	Email    string                 `json:"email" gorm:"unique"`
	Avatar   string                 `json:"avatar"`
	IsRoot   bool                   `json:"is_root"` // 是否是超级管理员
	Status   consts.AdminUserStatus `json:"status"`
}

func (AdminUser) TableName() string {
	return "admin_user"
}
