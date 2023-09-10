package account

import (
	"gorm.io/gorm"
	"server/db"
	"server/lib/errs"
	"server/modules/user"
)

type ServiceStruct struct{}

var Service *ServiceStruct

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
	db.Database.Create(&info)
	return info, nil
}

// CreateEmail 创建邮箱账号，邮箱用户名和密码
func (s *ServiceStruct) CreateEmail(username, email, password string) (user.User, error) {
	userInfo := user.User{
		Email: email,
	}
	accountInfo := Account{
		Type:     EmailAccountType,
		Username: username,
		Password: password,
		Email:    email,
	}

	if err := createUserAccount(&userInfo, &accountInfo); err != nil {
		return user.User{}, err
	}
	return userInfo, nil
}

// CreateWxMp 创建微信小程序账号，openid
func (s *ServiceStruct) CreateWxMp(openid string) (user.User, error) {
	userInfo := user.User{}
	accountInfo := Account{
		Type:     WxMpAccountType,
		WxOpenId: openid,
	}
	if err := createUserAccount(&userInfo, &accountInfo); err != nil {
		return user.User{}, err
	}
	return userInfo, nil
}

// UseEmailFindOne 使用邮箱查账号
func (s *ServiceStruct) UseEmailFindOne(email string) (Account, int64) {
	account := Account{}
	result := db.Database.First(&account, "email = ?", email)
	return account, result.RowsAffected
}

// UseUsernameFindOne 使用用户名查账号
func (s *ServiceStruct) UseUsernameFindOne(username string) (Account, int64) {
	account := Account{}
	result := db.Database.First(&account, "username = ?", username)
	return account, result.RowsAffected
}

// UseWxMpOpenIDFindOne 使用微信 openid 查账号
func (s *ServiceStruct) UseWxMpOpenIDFindOne(openid string) (Account, int64) {
	account := Account{}
	result := db.Database.First(&account, "openid = ?", openid)
	return account, result.RowsAffected
}

func createUserAccount(userInfo *user.User, accountInfo *Account) error {
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
