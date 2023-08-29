package user

import (
	"gorm.io/gorm"
	"selfserver/db"
)

type Service struct {
	db *gorm.DB
}

func init() {
	db.InitTable(User{})
}

var storage = db.Database.Model(&User{})
var ServiceInstance Service

func init() {
	ServiceInstance = Service{}
}

func (s *Service) GetList() []User {
	var list []User
	storage.Find(&list)
	return list
}
