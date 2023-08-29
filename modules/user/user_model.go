package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string
	Age    uint8
	Email  string `gorm:"uniqueIndex"`
	Avatar string
}

func (User) TableName() string {
	return "user"
}
