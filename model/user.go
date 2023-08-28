package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	username string `gorm:"unique"`
	email    string `gorm:"unique"`
	sex      string
}
