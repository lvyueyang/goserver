package service

import (
	"server/dal/model"
	"server/dal/query"
	"server/db"
)

type UserService struct {
}

var user = query.User

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

func (s *UserService) FindByID(id uint) (*model.User, error) {
	return user.Where(user.ID.Eq(id)).First()
}

func (s *UserService) Create(u CreateUser) *model.User {
	info := &model.User{
		Name:   u.Name,
		Age:    u.Age,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	user.Create(info)
	return info
}
