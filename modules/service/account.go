package service

import (
	"context"
	"server/consts"
	"server/dal/model"
	"server/dal/query"
	"server/lib/errs"
)

type AccountService struct{}

var account = query.Account

func NewAccountService() *AccountService {
	return new(AccountService)
}

// CreateNormal 创建普通账号，用户名和密码
func (s *AccountService) CreateNormal(userID uint, username, password string) (model.Account, error) {
	info := model.Account{
		Type:     consts.NormalAccountType,
		Username: username,
		Password: password,
		UserID:   userID,
	}
	if err := account.WithContext(context.Background()).Create(&info); err != nil {
		return info, err
	}

	return info, nil
}

// CreateEmail 创建邮箱账号，邮箱用户名和密码
func (s *AccountService) CreateEmail(username, email, password string) (model.User, error) {
	userInfo := model.User{
		Email: email,
	}
	accountInfo := model.Account{
		Type:     consts.EmailAccountType,
		Username: username,
		Password: password,
		Email:    email,
	}

	if err := createUserAccount(&userInfo, &accountInfo); err != nil {
		return model.User{}, err
	}
	return userInfo, nil
}

// CreateWxMp 创建微信小程序账号，openid
func (s *AccountService) CreateWxMp(openid string) (model.User, error) {
	userInfo := model.User{}
	accountInfo := model.Account{
		Type:     consts.WxMpAccountType,
		WxOpenId: openid,
	}
	if err := createUserAccount(&userInfo, &accountInfo); err != nil {
		return model.User{}, err
	}
	return userInfo, nil
}

// UseEmailFindOne 使用邮箱查账号
func (s *AccountService) UseEmailFindOne(email string) (info *model.Account, err error) {
	return account.WithContext(context.Background()).Where(account.Email.Eq(email)).First()
}

// UseUsernameFindOne 使用用户名查账号
func (s *AccountService) UseUsernameFindOne(username string) (info *model.Account, err error) {
	return account.WithContext(context.Background()).Where(account.Username.Eq(username)).First()
}

// UseWxMpOpenIDFindOne 使用微信 openid 查账号
func (s *AccountService) UseWxMpOpenIDFindOne(openid string) (info *model.Account, err error) {
	return account.WithContext(context.Background()).Where(account.WxOpenId.Eq(openid)).First()

}

func createUserAccount(userInfo *model.User, accountInfo *model.Account) error {
	err := query.Q.Transaction(func(tx *query.Query) error {

		if err := tx.User.WithContext(context.Background()).Create(userInfo); err != nil {
			return errs.CreateServerError("创建用户失败", err, userInfo)
		}
		accountInfo.UserID = userInfo.ID
		if err := tx.Account.WithContext(context.Background()).Create(accountInfo); err != nil {
			return errs.CreateServerError("创建用户账号失败", err, accountInfo)
		}

		return nil
	})
	return errs.CreateServerError("创建用户账号事务失败", err, nil)
}
