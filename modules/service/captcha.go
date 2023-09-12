package service

import (
	"gorm.io/gorm"
	"server/consts"
	"server/db"
	"server/lib/errs"
	"server/modules/model"
	"server/utils"
	"time"
)

type CaptchaStruct struct{}

var CaptchaService *CaptchaStruct

func init() {
	//db.InitTable(new(model.Captcha))
	CaptchaService = new(CaptchaStruct)
}

func (s *CaptchaStruct) FindByID(id uint) (model.Captcha, *gorm.DB) {
	var info = model.Captcha{}

	result := db.Database.First(&info, "id = ?", id)

	return info, result
}

// FindByOne 根据场景和值查找
func (s *CaptchaStruct) FindByOne(current string, scenes consts.CaptchaScenes) (model.Captcha, *gorm.DB) {
	var info = model.Captcha{}

	result := db.Database.Order("updated_at desc").First(&info, "current = ? AND scenes = ? ", current, scenes)

	return info, result
}

// ValidateCode 验证手机/邮箱验证码
func validateCode(info model.Captcha, code string) (bool, error) {
	if info.ID == 0 {
		return false, errs.CreateClientError("验证码不存在", info)
	}
	if time.Now().After(info.Expiration) {
		return false, errs.CreateClientError("验证码已过期", info)
	}
	if info.Code != code {
		return false, errs.CreateClientError("验证码错误", info)
	}
	return true, nil
}

// Validate 验证手机/邮箱验证码
func (s *CaptchaStruct) Validate(current, code string, scenes consts.CaptchaScenes) (bool, error) {
	info, err := s.FindByOne(current, scenes)
	if err != nil {
		return false, errs.CreateClientError("验证码不存在", info)
	}
	return validateCode(info, code)
}

func (s *CaptchaStruct) Create(currentType consts.CaptchaType, current string, scenes consts.CaptchaScenes) model.Captcha {
	code := utils.GenCaptcha()
	info := model.Captcha{
		CurrentType: currentType,
		Current:     current,
		Expiration:  time.Now().Add(5 * time.Minute),
		Status:      consts.CaptchaStatusUnused,
		Code:        code,
		Scenes:      scenes,
	}

	db.Database.Create(&info)

	return info
}

// 批量删除过期验证码
func (s *CaptchaStruct) multiDeleteExpiration() {
	var info = model.Captcha{}

	db.Database.Delete(info, "expiration < ?", time.Now().Add(1*time.Minute))
}
