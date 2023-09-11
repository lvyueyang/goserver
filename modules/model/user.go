package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Age     uint32
	Email   string `gorm:"uniqueIndex"`
	Avatar  string
	Account []Account
}

func (User) TableName() string {
	return "user"
}
