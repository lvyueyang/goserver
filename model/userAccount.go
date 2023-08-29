package model

import "gorm.io/gorm"

// UserAccount 用户账号
type UserAccount struct {
	gorm.Model
	Name string `gorm:"unique"`
	Type string
}
