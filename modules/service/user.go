package service

import (
	"server/dal/model"
	"server/db"
)

type UserService struct {
}

func NewUserService() *UserService {
	return new(UserService)
}

func (s *UserService) GetList() []model.User {
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

func (s *UserService) FindByID(id uint) model.User {
	user := model.User{}
	db.Database.First(&user, "id = ?", id)
	return user
}

func (s *UserService) Create(u CreateUser) model.User {
	user := model.User{
		Name:   u.Name,
		Age:    u.Age,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	db.Database.Create(&user)
	return user
}
