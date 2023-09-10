package model

import (
	"gorm.io/gorm"
	"server/consts"
)

type Account struct {
	gorm.Model
	Type      consts.AccountType
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	WxOpenId  string
	WxUnionId string
	UserID    uint
}

func (*Account) TableName() string {
	return "account"
}
