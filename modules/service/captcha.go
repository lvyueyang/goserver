package service

import (
	"server/consts"
	"server/db"
	"server/modules/model"
	"server/utils"
	"time"
)

type CaptchaStruct struct{}

var CaptchaService *CaptchaStruct

func init() {
	db.InitTable(new(model.Captcha))
	CaptchaService = new(CaptchaStruct)
}

func (s *CaptchaStruct) FindByID(id uint) (model.Captcha, error) {
	var info = model.Captcha{}

	db.Database.First(&info, "id = ?", id)

	return info, nil
}

func (s *CaptchaStruct) Create(currentType consts.CaptchaType, current string) (model.Captcha, error) {
	code := utils.GenCaptcha()
	info := model.Captcha{
		CurrentType: currentType,
		Current:     current,
		Expiration:  time.Now().Add(5 * time.Minute),
		Status:      consts.CaptchaStatusUnused,
		Code:        code,
	}

	db.Database.Create(&info)

	return info, nil
}
