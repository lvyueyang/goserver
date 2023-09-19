package service

import (
	"server/consts"
	"server/dal/dao"
	"server/dal/model"
	"server/lib/errs"
)

type AccountService struct{}

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
	if err := dao.Account.Create(&info); err != nil {
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
	return dao.Account.Where(dao.Account.Email.Eq(email)).Take()
}

// UseUsernameFindOne 使用用户名查账号
func (s *AccountService) UseUsernameFindOne(username string) (info *model.Account, err error) {
	return dao.Account.Where(dao.Account.Username.Eq(username)).Take()
}

// UseWxMpOpenIDFindOne 使用微信 openid 查账号
func (s *AccountService) UseWxMpOpenIDFindOne(openid string) (info *model.Account, err error) {
	return dao.Account.Where(dao.Account.WxOpenId.Eq(openid)).Take()

}

func createUserAccount(userInfo *model.User, accountInfo *model.Account) error {
	err := dao.Q.Transaction(func(tx *dao.Query) error {
		if err := tx.User.Create(userInfo); err != nil {
			return errs.CreateServerError("创建用户失败", err, userInfo)
		}
		accountInfo.UserID = userInfo.ID
		if err := tx.Account.Create(accountInfo); err != nil {
			return errs.CreateServerError("创建用户账号失败", err, accountInfo)
		}
		return nil
	})
	if err != nil {
		return errs.CreateServerError("创建用户账号事务失败", err, nil)
	}
	return nil
}
