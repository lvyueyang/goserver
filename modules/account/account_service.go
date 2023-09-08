package account

import (
	"errors"
	"gorm.io/gorm"
	"server/db"
)

type ServiceStruct struct {
	db *gorm.DB
}

var Service *ServiceStruct

var storage = db.Database.Model(&Account{})

func init() {
	db.InitTable(Account{})
	Service = &ServiceStruct{}
}

// CreateNormal 创建普通账号，用户名和密码
func (s *ServiceStruct) CreateNormal(userID uint, username, password string) (Account, error) {
	info := Account{
		Type:     NormalAccountType,
		Username: username,
		Password: password,
		UserID:   userID,
	}
	storage.Create(&info)
	return info, nil
}

// CreateEmail 创建邮箱账号，邮箱用户名和密码
func (s *ServiceStruct) CreateEmail(userID uint, username, email, password string) (Account, error) {
	nilAccount := Account{}

	if list := s.UseEmailFindList(email); len(list) > 0 {
		return nilAccount, errors.New("邮箱已注册")
	}
	if list := s.UseUsernameFindList(username); len(list) > 0 {
		return nilAccount, errors.New("用户名已注册")
	}

	info := Account{
		Type:     EmailAccountType,
		Username: username,
		Password: password,
		Email:    email,
		UserID:   userID,
	}
	storage.Create(&info)
	return info, nil
}

// CreateWxMp 创建微信小程序账号，openid
func (s *ServiceStruct) CreateWxMp(userID uint, openid string) (Account, error) {
	info := Account{
		Type:     WxMpAccountType,
		UserID:   userID,
		WxOpenId: openid,
	}
	storage.Create(&info)
	return info, nil
}

// UseEmailFindList 使用邮箱查账号
func (s *ServiceStruct) UseEmailFindList(email string) []Account {
	var list []Account
	storage.Find(&list, "email = ?", email)
	return list
}

// UseUsernameFindList 使用用户名查账号
func (s *ServiceStruct) UseUsernameFindList(username string) []Account {
	var list []Account
	storage.Find(&list, "username = ?", username)
	return list
}
