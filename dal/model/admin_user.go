package model

import (
	"gorm.io/gorm"
)

type AdminUser struct {
	gorm.Model
	Name     string `json:"name"`
	UserName string `json:"user_name" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Avatar   string `json:"avatar"`
	IsRoot   bool   `json:"is_root"` // 是否是超级管理员
}

func (AdminUser) TableName() string {
	return "admin_user"
}
