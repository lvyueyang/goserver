package service

import (
	"context"
	"server/consts"
	"server/dal/model"
	"server/dal/query"
	"server/lib/errs"
	"server/utils"
	"time"
)

var captcha = query.Captcha

type CaptchaService struct{}

func NewCaptchaService() *CaptchaService {
	return new(CaptchaService)
}

func (s *CaptchaService) FindByID(id uint) model.Captcha {
	info, _ := captcha.WithContext(context.Background()).Where(captcha.ID.Eq(id)).First()
	return *info
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
func (s *CaptchaService) Validate(current, code string, scenes consts.CaptchaScenes) (bool, error) {
	info, err := captcha.WithContext(context.Background()).Order(captcha.CreatedAt.Desc()).First()

	if err != nil {
		return false, errs.CreateClientError("验证码不存在", info)
	}
	return validateCode(*info, code)
}

func (s *CaptchaService) Create(currentType consts.CaptchaType, current string, scenes consts.CaptchaScenes) model.Captcha {
	code := utils.GenCaptcha()
	info := model.Captcha{
		CurrentType: currentType,
		Current:     current,
		Expiration:  time.Now().Add(5 * time.Minute),
		Status:      consts.CaptchaStatusUnused,
		Code:        code,
		Scenes:      scenes,
	}

	captcha.WithContext(context.Background()).Create(&info)

	return info
}

// 批量删除过期验证码
func (s *CaptchaService) multiDeleteExpiration() {
	captcha.WithContext(context.Background()).Where(captcha.Expiration.Lt(time.Now().Add(1 * time.Minute))).Delete()
}
