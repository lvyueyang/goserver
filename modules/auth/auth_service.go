package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"server/modules/account"
	"server/modules/user"
)

type ServiceStruct struct{}

var Service *ServiceStruct

func init() {
	Service = &ServiceStruct{}
}

type LoginOptions struct {
	Username string
}

// UsernameAndPasswordRegister 使用用户名邮箱和密码注册
func (s *ServiceStruct) UsernameAndPasswordRegister(opt struct {
	Username string
	Email    string
	Code     string
	Password string
}) (token user.User, err error) {
	nilUser := user.User{}

	if list := account.Service.UseEmailFindList(opt.Email); len(list) > 0 {
		return nilUser, errors.New("邮箱已注册")
	}
	if list := account.Service.UseUsernameFindList(opt.Username); len(list) > 0 {
		return nilUser, errors.New("用户名已注册")
	}

	// 创建用户
	userinfo := user.Service.Create(user.CreateUser{
		Email: opt.Email,
	})

	userAccount := account.Service.CreateEmail()

	// 密码加盐
	hashPasword, err := bcrypt.GenerateFromPassword([]byte(opt.Password), bcrypt.DefaultCost)
	if err != nil {
		return user.User{}, err
	}

	user.Service.Create(user.CreateUser{
		Email: opt.Email,
	})

	return user.User{}, nil
}

// UsernameAndPasswordLogin 使用用户名和密码登录
func (s *ServiceStruct) UsernameAndPasswordLogin(username string, password string) (token string, err error) {
	return "", nil
}

func (s *ServiceStruct) Register() {
}
