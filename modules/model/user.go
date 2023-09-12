package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Age     uint32 `json:"age"`
	Email   string `json:"email" gorm:"uniqueIndex"`
	Avatar  string `json:"avatar"`
	Account []Account
}

func (User) TableName() string {
	return "user"
}
