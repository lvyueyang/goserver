package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Type      uint
	Username  string `gorm:"unique"`
	Password  string
	Email     string `gorm:"unique"`
	WxOpenId  string
	WxUnionId string
	UserID    uint
}
