package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/config"
	"server/consts"
	"server/dal/dao"
	"server/dal/model"
	"server/lib/errs"
	"server/utils"
)

type AuthService struct {
	accountService *AccountService
	userService    *UserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		accountService: NewAccountService(),
		userService:    NewUserService(),
	}
}

type LoginOptions struct {
	Username string
}

type UsernameAndPasswordRegisterOptions struct {
	Username string
	Email    string
	Code     string
	Password string
}

// UsernameAndPasswordRegister 使用用户名邮箱和密码注册
func (s *AuthService) UsernameAndPasswordRegister(email, username, password string) (string, error) {
	var opt = map[string]string{
		"email":    email,
		"username": username,
		"password": password,
	}
	fmt.Println("accountService", s.accountService)

	// 验证用户名和邮箱是否已被使用
	if _, err := dao.Account.Where(dao.Account.Email.Eq(email)).Take(); err == nil {
		return "", &errs.ClientError{Msg: "邮箱已存在", Info: nil}
	}
	if _, err := dao.Account.Where(dao.Account.Username.Eq(username)).Take(); err == nil {
		return "", &errs.ClientError{Msg: "用户名已存在", Info: nil}
	}

	var hashPassword []byte
	// 密码加盐
	if result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", &errs.ServerError{Msg: "密码加盐失败", Err: err, Info: opt}
	} else {
		hashPassword = result
	}

	var userInfo model.User
	// 创建账号
	if result, err := s.accountService.CreateEmail(username, email, string(hashPassword)); err != nil {
		return "", &errs.ServerError{Msg: "账号创建失败", Err: err, Info: opt}
	} else {
		userInfo = result
	}

	token, err := utils.CreateUserToken(userInfo, consts.EmailAccountType, config.Config.Auth.TokenSecret)

	if err != nil {
		errInfo := map[string]any{"userInfo": userInfo, "opt": opt}
		return "", errs.CreateServerError("Token 生成失败", err, errInfo)
	}
	return token, nil
}

// UsernameAndPasswordLogin 使用用户名和密码登录
func (s *AuthService) UsernameAndPasswordLogin(username string, password string) (string, error) {
	info, _ := s.accountService.UseUsernameFindOne(username)
	if info.ID == 0 {
		return "", &errs.ClientError{Msg: "用户名未注册", Info: nil}
	}
	err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password))
	if err != nil {
		return "", errs.CreateClientError("密码错误", nil)
	}
	userinfo, _ := s.userService.FindByID(info.UserID)
	token, err := utils.CreateUserToken(*userinfo, info.Type, config.Config.Auth.TokenSecret)
	if err != nil {
		return "", errs.CreateServerError("Token 生成失败", err, nil)
	}
	return token, nil
}

func (s *AuthService) Register() {
}
