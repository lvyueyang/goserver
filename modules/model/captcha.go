package model

import (
	"gorm.io/gorm"
	"server/consts"
	"time"
)

type Captcha struct {
	gorm.Model
	Expiration  time.Time // 过期时间
	Current     string    // 所属目标 手机号/邮箱
	CurrentType consts.CaptchaType
	Code        string
	Status      consts.CaptchaStatus
}

func (Captcha) TableName() string {
	return "captcha"
}
