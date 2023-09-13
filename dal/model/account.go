package model

import (
	"gorm.io/gorm"
	"server/consts"
)

type Account struct {
	gorm.Model
	Type      consts.AccountType `json:"type"`
	Username  string             `json:"username" gorm:"unique"`
	Password  string             `json:"password"`
	Email     string             `json:"email" gorm:"unique"`
	WxOpenId  string             `json:"wx_open_id"`
	WxUnionId string             `json:"wx_union_id"`
	UserID    uint               `json:"user_id"`
}

func (*Account) TableName() string {
	return "account"
}
