package service

import (
	"server/db"
	"server/internal/model"
)

type ServiceStruct struct {
}

var UserService *ServiceStruct

func init() {
	db.InitTable(model.User{})
	UserService = &ServiceStruct{}
}

func (s *ServiceStruct) GetList() []model.User {
	var list []model.User
	db.Database.Find(&list)
	return list
}

type CreateUser struct {
	Name   string
	Age    uint32
	Email  string
	Avatar string
}

func (s *ServiceStruct) FindByID(id uint) model.User {
	user := model.User{}
	db.Database.First(&user, "id = ?", id)
	return user
}

func (s *ServiceStruct) Create(u CreateUser) model.User {
	user := model.User{
		Name:   u.Name,
		Age:    u.Age,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	db.Database.Create(&user)
	return user
}
