package user

import (
	"gorm.io/gorm"
	"server/modules/account"
)

type User struct {
	gorm.Model
	Name    string
	Age     uint32
	Email   string `gorm:"uniqueIndex"`
	Avatar  string
	Account []account.Account
}

func (User) TableName() string {
	return "user"
}
