package user

import (
	"gorm.io/gorm"
	"server/db"
)

type ServiceStruct struct {
	db *gorm.DB
}

var storage = db.Database.Model(&User{})
var Service *ServiceStruct

func init() {
	db.InitTable(User{})
	Service = &ServiceStruct{}
}

func (s *ServiceStruct) GetList() []User {
	var list []User
	storage.Find(&list)
	return list
}

type CreateUser struct {
	Name   string
	Age    uint32
	Email  string
	Avatar string
}

func (s *ServiceStruct) Create(u CreateUser) User {
	user := User{
		Name:   u.Name,
		Age:    u.Age,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	storage.Create(&user)
	return user
}
