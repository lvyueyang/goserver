package service

import (
	"golang.org/x/crypto/bcrypt"
	"server/config"
	"server/consts"
	"server/dal/dao"
	"server/dal/model"
	"server/lib/errs"
	"server/types"
	"server/utils"
	"strconv"
	"strings"
)

type AdminUserService struct {
}

func NewAdminUserService() *AdminUserService {
	return new(AdminUserService)
}

type FindAdminUserListOption struct {
	types.Pagination
	types.Order
	Keyword string `json:"keyword" form:"keyword"`
}

func (s *AdminUserService) FindList(query FindAdminUserListOption) (utils.ListResult[[]*model.AdminUser], error) {
	result := utils.ListResult[[]*model.AdminUser]{}
	u := dao.AdminUser
	q := u.Where(
		u.Username.Like("%" + query.Keyword + "%"),
	).Or(
		u.Name.Like("%" + query.Keyword + "%"),
	)

	if id, err := strconv.ParseUint(query.Keyword, 10, 64); err == nil {
		q = q.Or(u.ID.Eq(uint(id)))
	}

	if query.OrderKey != "" {
		col, _ := u.GetFieldByName(query.OrderKey)
		if strings.ToLower(query.OrderType) == "desc" {
			q = q.Order(col.Desc())
		} else {
			q = q.Order(col)
		}
	}

	if list, total, err := q.Preload(u.Roles).FindByPage(utils.PageTrans(query.Pagination)); err != nil {
		return result, err
	} else {
		result.List = list
		result.Total = total
	}

	return result, nil
}

type CreateAdminUser struct {
	Name   string
	Age    uint32
	Email  string
	Avatar string
}

func (s *AdminUserService) Create(user model.AdminUser) (*model.AdminUser, error) {
	var nilUser = new(model.AdminUser)
	if _, err := dao.AdminUser.Where(dao.AdminUser.Username.Eq(user.Username)).Take(); err == nil {
		return nilUser, errs.CreateServerError("用户名已存在", err, nil)
	}
	if _, err := dao.AdminUser.Where(dao.AdminUser.Email.Eq(user.Email)).Take(); err == nil {
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
		Avatar:   user.Avatar,
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

func (s *AdminUserService) Update(id uint, user model.AdminUser) error {
	if _, err := dao.AdminUser.FindByID(id); err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	}
	data := model.AdminUser{
		Name:   user.Name,
		Avatar: user.Avatar,
	}

	if _, err := dao.AdminUser.Where(dao.AdminUser.ID.Eq(id)).Updates(data); err != nil {
		return err
	}
	return nil
}

func (s *AdminUserService) Delete(id uint) error {
	if user, err := dao.AdminUser.FindByID(id); err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	} else {
		if user.IsRoot {
			return errs.CreateClientError("禁止删除超级管理员", nil)
		}
	}

	if _, err := dao.AdminUser.Where(dao.AdminUser.ID.Eq(id)).Delete(); err != nil {
		return err
	}
	return nil
}

func (s *AdminUserService) UpdatePassword(id uint, password string) error {
	user, err := dao.AdminUser.FindByID(id)
	if err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	}

	return updateAdminUserPassword(user.ID, password)
}

func (s *AdminUserService) UpdateStatus(id uint, status consts.AdminUserStatus) error {
	user, err := dao.AdminUser.FindByID(id)
	if err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	}

	if _, err := dao.AdminUser.Where(dao.AdminUser.ID.Eq(user.ID)).Update(dao.AdminUser.Status, status); err != nil {
		return errs.CreateServerError("状态更新失败", err, status)
	}

	return nil
}

// CreateRootUser  创建超管账号
func (s *AdminUserService) CreateRootUser(username, name, password, email string) (string, error) {
	if _, err := dao.AdminUser.Where(dao.AdminUser.IsRoot).Take(); err == nil {
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
		return token, nil
	}
}

func (s *AdminUserService) ResetPassword(email, password string) error {
	user, err := dao.AdminUser.Where(dao.AdminUser.Email.Eq(email)).Take()
	if err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	}

	return updateAdminUserPassword(user.ID, password)
}

func (s *AdminUserService) AddRole(userID uint, roleIDs []uint) error {
	user, err := dao.AdminUser.FindByID(userID)
	if err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	}
	if user.IsRoot {
		return errs.CreateClientError("超管用户禁止更新角色", nil)
	}
	roles, err := dao.AdminRole.Where(dao.AdminRole.ID.In(roleIDs...)).Find()
	if err != nil {
		return errs.CreateServerError("角色不存在", err, nil)
	}

	return dao.AdminUser.Roles.Model(user).Append(roles...)
}

func (s *AdminUserService) DeleteRole(userID uint, roleIDs []uint) error {
	user, err := dao.AdminUser.FindByID(userID)
	if err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	}
	if user.IsRoot {
		return errs.CreateClientError("超管用户禁止更新角色", nil)
	}
	roles, err := dao.AdminRole.Where(dao.AdminRole.ID.In(roleIDs...)).Find()
	if err != nil {
		return errs.CreateServerError("角色不存在", err, nil)
	}

	return dao.AdminUser.Roles.Model(user).Delete(roles...)
}

func (s *AdminUserService) UpdateRole(userID uint, roleIDs []uint) error {
	user, err := dao.AdminUser.FindByID(userID)
	if err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	}
	if user.IsRoot {
		return errs.CreateClientError("超管用户禁止更新角色", nil)
	}
	roles, err := dao.AdminRole.Where(dao.AdminRole.ID.In(roleIDs...)).Find()
	if err != nil {
		return errs.CreateServerError("角色不存在", err, nil)
	}

	if err := dao.AdminUser.Roles.Model(user).Clear(); err != nil {
		return errs.CreateServerError("清除用户角色失败", err, nil)
	}

	return dao.AdminUser.Roles.Model(user).Append(roles...)
}

// 用于验证是否为超管用户在操作自己的账号
func (s *AdminUserService) OnlyRootAdminUser(userID, currentID uint) error {
	user, err := dao.AdminUser.FindByID(userID)
	if err != nil {
		return errs.CreateServerError("用户不存在", err, nil)
	}
	if user.IsRoot && userID != currentID {
		return errs.CreateClientError("超管用户信息仅超管本人可以修改", nil)
	}
	return nil
}

func updateAdminUserPassword(id uint, password string) error {
	// 密码加盐
	var hashPassword []byte
	if result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return errs.CreateServerError("密码加盐失败", err, password)
	} else {
		hashPassword = result
	}

	if _, err := dao.AdminUser.Where(dao.AdminUser.ID.Eq(id)).Update(dao.AdminUser.Password, hashPassword); err != nil {
		return errs.CreateServerError("密码更新失败", err, password)
	}

	return nil
}
