package service

import (
	"fmt"
	"server/consts"
	"server/dal/dao"
	"server/dal/model"
	"server/lib/errs"
	"server/utils"
	"time"
)

type CaptchaService struct{}

func NewCaptchaService() *CaptchaService {
	return new(CaptchaService)
}

func (s *CaptchaService) FindByID(id uint) (*model.Captcha, error) {
	return dao.Captcha.Where(dao.Captcha.ID.Eq(id)).Take()
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
func (s *CaptchaService) Validate(current string, currentType consts.CaptchaType, code string, scenes consts.CaptchaScenes) (bool, error) {
	info, err := dao.Captcha.Where(
		dao.Captcha.Current.Eq(current),
		dao.Captcha.CurrentType.Eq(uint(currentType)),
		dao.Captcha.Code.Eq(code),
		dao.Captcha.Scenes.Eq(uint(scenes)),
	).Last()

	if err != nil {
		return false, errs.CreateClientError("验证码错误", info)
	}
	return validateCode(*info, code)
}

func (s *CaptchaService) Create(currentType consts.CaptchaType, current string, scenes consts.CaptchaScenes) (*model.Captcha, error) {
	code := utils.GenCaptcha()
	info := model.Captcha{
		CurrentType: currentType,
		Current:     current,
		Expiration:  time.Now().Add(5 * time.Minute),
		Status:      consts.CaptchaStatusUnused,
		Code:        code,
		Scenes:      scenes,
	}

	if err := dao.Captcha.Create(&info); err != nil {
		return new(model.Captcha), err
	}

	return &info, nil
}

// ClearExpiration 清理过期验证码
func (s *CaptchaService) ClearExpiration() {
	dao.Captcha.Where(dao.Captcha.Expiration.Lt(time.Now().Add(1 * time.Minute))).Delete()
	fmt.Println("已清除过期验证码")
}
