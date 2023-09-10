package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Type      Type
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
