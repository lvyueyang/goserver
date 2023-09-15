package model

import (
	"server/consts"
)

type User struct {
	BaseModel
	Name    string            `json:"name"`
	Age     uint32            `json:"age"`
	Email   string            `json:"email" gorm:"unique"`
	Avatar  string            `json:"avatar"`
	Status  consts.UserStatus `json:"status" gorm:"default=1"`
	Account []Account
}

func (User) TableName() string {
	return "user"
}
