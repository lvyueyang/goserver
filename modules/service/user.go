package service

import (
	"server/dal/dao"
	"server/dal/model"
)

type UserService struct {
}

func NewUserService() *UserService {
	return new(UserService)
}

func (s *UserService) FindList() ([]*model.User, error) {
	return dao.User.Find()
}

type CreateUser struct {
	Name   string
	Age    uint32
	Email  string
	Avatar string
}

func (s *UserService) FindByID(id uint) (*model.User, error) {
	return dao.User.FindByID(id)
}

func (s *UserService) Create(u CreateUser) *model.User {
	info := &model.User{
		Name:   u.Name,
		Age:    u.Age,
		Email:  u.Email,
		Avatar: u.Avatar,
	}
	dao.User.Create(info)
	return info
}
