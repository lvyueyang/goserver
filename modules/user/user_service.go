package user

import (
	"server/db"
)

type ServiceStruct struct {
}

var Service *ServiceStruct

func init() {
	db.InitTable(User{})
	Service = &ServiceStruct{}
}

func (s *ServiceStruct) GetList() []User {
	var list []User
	db.Database.Find(&list)
	return list
}

type CreateUser struct {
	Name   string
	Age    uint32
	Email  string
	Avatar string
}

func (s *ServiceStruct) FindByID(id uint) User {
	user := User{}
	db.Database.First(&user, "id = ?", id)
	return user
}

func (s *ServiceStruct) Create(u CreateUser) User {
	user := User{
		Name:   u.Name,
		Age:    u.Age,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	db.Database.Create(&user)
	return user
}
