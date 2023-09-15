package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"server/config"
	"server/consts"
	"server/dal/dao"
	"server/dal/model"
	"server/lib/errs"
	"server/types"
	"server/utils"
)

type AdminUserService struct {
}

func NewAdminUserService() *AdminUserService {
	return new(AdminUserService)
}

type FindUserListOption struct {
	types.PaginationQuery
	types.OrderQuery
	Keyword string `json:"keyword"`
}

func (s *AdminUserService) FindList(query FindUserListOption) ([]*model.AdminUser, error) {
	return dao.AdminUser.Where(dao.AdminUser.Username.Like(query.Keyword)).Find()
}

type CreateAdminUser struct {
	Name   string
	Age    uint32
	Email  string
	Avatar string
}

func (s *AdminUserService) Create(user model.AdminUser) (*model.AdminUser, error) {
	var nilUser = new(model.AdminUser)
	if _, err := dao.AdminUser.Where(dao.AdminUser.Username.Eq(user.Username)).First(); err == nil {
		return nilUser, errs.CreateServerError("用户名已存在", err, nil)
	}
	if _, err := dao.AdminUser.Where(dao.AdminUser.Email.Eq(user.Email)).First(); err == nil {
		return nilUser, errs.CreateServerError("邮箱已存在", err, nil)
	}

	// 密码加盐
	var hashPassword []byte
	if result, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		return nilUser, &errs.ServerError{Msg: "密码加盐失败", Err: err, Info: user}
	} else {
		hashPassword = result
	}

	var data = &model.AdminUser{
		Name:     user.Name,
		Username: user.Username,
		Password: string(hashPassword),
		Email:    user.Email,
		IsRoot:   false,
		Status:   consts.AdminUserStatusNormal,
	}
	if err := dao.AdminUser.Create(data); err != nil {
		return nilUser, err
	}

	return data, nil
}

// CreateRootUser  创建超管账号
func (s *AdminUserService) CreateRootUser(username, name, password, email string) (string, error) {
	if _, err := dao.AdminUser.Where(dao.AdminUser.IsRoot).First(); err == nil {
		return "", errs.CreateServerError("超管账户已存在禁止重复创建", err, nil)
	}

	var hashPassword []byte
	// 密码加盐
	if result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", &errs.ServerError{Msg: "密码加盐失败", Err: err, Info: password}
	} else {
		hashPassword = result
	}

	var user = &model.AdminUser{
		Name:     name,
		Username: username,
		Password: string(hashPassword),
		Email:    email,
		IsRoot:   true,
		Status:   consts.AdminUserStatusNormal,
	}
	if err := dao.AdminUser.Create(user); err != nil {
		return "", err
	}

	if token, err := utils.CreateAdminUserToken(user.ID, config.Config.Auth.AdminTokenSecret); err != nil {
		return "", nil
	} else {
		fmt.Println("token", token)
		return token, nil
	}
}

func (s *AdminUserService) ResetPassword(email, password string) error {
	if _, err := dao.AdminUser.Where(dao.AdminUser.Email.Eq(email)).First(); err != nil {
		return errs.CreateServerError("用户名不存在", err, nil)
	}

	// 密码加盐
	var hashPassword []byte
	if result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return &errs.ServerError{Msg: "密码加盐失败", Err: err, Info: password}
	} else {
		hashPassword = result
	}

	if _, err := dao.AdminUser.Where(dao.AdminUser.Email.Eq(email)).Update(dao.AdminUser.Password, hashPassword); err != nil {
		return err
	}

	return nil
}
