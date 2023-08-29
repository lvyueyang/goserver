package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Email string `gorm:"unique"`
	Sex   uint8
}
