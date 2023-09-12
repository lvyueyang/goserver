package service

import (
	"server/db"
	"server/modules/model"
)

type UserServiceStruct struct {
}

var UserService *UserServiceStruct

func init() {
	//db.InitTable(new(model.User))
	UserService = new(UserServiceStruct)
}

func (s *UserServiceStruct) GetList() []model.User {
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

func (s *UserServiceStruct) FindByID(id uint) model.User {
	user := model.User{}
	db.Database.First(&user, "id = ?", id)
	return user
}

func (s *UserServiceStruct) Create(u CreateUser) model.User {
	user := model.User{
		Name:   u.Name,
		Age:    u.Age,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	db.Database.Create(&user)
	return user
}
