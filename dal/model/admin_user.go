package model

import (
	"server/consts"
)

type AdminUser struct {
	BaseModel
	Name     string                 `json:"name"`
	Username string                 `json:"username" gorm:"unique"`
	Password string                 `json:"-"`
	Email    string                 `json:"email" gorm:"unique"`
	Avatar   string                 `json:"avatar"`
	IsRoot   bool                   `json:"is_root"` // 是否是超级管理员
	Status   consts.AdminUserStatus `json:"status"`
	Roles    []*AdminRole           `json:"roles" gorm:"many2many:admin_user_roles;"`
	News     []*News                `json:"news" gorm:"foreignKey:AuthorID"`
}

func (AdminUser) TableName() string {
	return "admin_user"
}
