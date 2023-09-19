package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/config"
	"server/dal/dao"
	"server/lib/errs"
	"server/utils"
)

type AdminAuthService struct {
	adminUserService *AdminUserService
}

func NewAdminAuthService() *AdminAuthService {
	return new(AdminAuthService)
}

func (s *AdminAuthService) UsernameAndPasswordLogin(username string, password string) (string, error) {
	fmt.Println("username", username)
	info, err := dao.AdminUser.Where(dao.AdminUser.Username.Eq(username)).Take()
	if err != nil {
		return "", &errs.ClientError{Msg: "用户未注册", Info: nil}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password)); err != nil {
		return "", errs.CreateClientError("密码错误", nil)
	}
	token, err := utils.CreateAdminUserToken(info.ID, config.Config.Auth.AdminTokenSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}
