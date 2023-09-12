package service

import (
	"gorm.io/gorm"
	"server/consts"
	"server/db"
	"server/lib/errs"
	"server/modules/model"
)

type AccountServiceStruct struct{}

var AccountService *AccountServiceStruct

func init() {
	//db.InitTable(new(model.Account))
	AccountService = new(AccountServiceStruct)
}

// CreateNormal 创建普通账号，用户名和密码
func (s *AccountServiceStruct) CreateNormal(userID uint, username, password string) (model.Account, error) {
	info := model.Account{
		Type:     consts.NormalAccountType,
		Username: username,
		Password: password,
		UserID:   userID,
	}
	db.Database.Create(&info)
	return info, nil
}

// CreateEmail 创建邮箱账号，邮箱用户名和密码
func (s *AccountServiceStruct) CreateEmail(username, email, password string) (model.User, error) {
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
func (s *AccountServiceStruct) CreateWxMp(openid string) (model.User, error) {
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
func (s *AccountServiceStruct) UseEmailFindOne(email string) (model.Account, int64) {
	account := model.Account{}
	result := db.Database.First(&account, "email = ?", email)
	return account, result.RowsAffected
}

// UseUsernameFindOne 使用用户名查账号
func (s *AccountServiceStruct) UseUsernameFindOne(username string) (model.Account, int64) {
	account := model.Account{}
	result := db.Database.First(&account, "username = ?", username)
	return account, result.RowsAffected
}

// UseWxMpOpenIDFindOne 使用微信 openid 查账号
func (s *AccountServiceStruct) UseWxMpOpenIDFindOne(openid string) (model.Account, int64) {
	account := model.Account{}
	result := db.Database.First(&account, "openid = ?", openid)
	return account, result.RowsAffected
}

func createUserAccount(userInfo *model.User, accountInfo *model.Account) error {
	err := db.Database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userInfo).Error; err != nil {
			return errs.CreateServerError("创建用户失败", err, userInfo)
		}
		accountInfo.UserID = userInfo.ID
		if err := tx.Create(&accountInfo).Error; err != nil {
			return errs.CreateServerError("创建用户账号失败", err, accountInfo)
		}

		return nil
	})
	return errs.CreateServerError("创建用户账号事务失败", err, nil)
}
