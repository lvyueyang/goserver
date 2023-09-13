package model

import (
	"server/consts"
	"time"
)

type Captcha struct {
	ID          uint                 `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Expiration  time.Time            `json:"expiration"` // 过期时间
	Current     string               `json:"current"`    // 所属目标 手机号/邮箱
	CurrentType consts.CaptchaType   `json:"current_type"`
	Code        string               `json:"code"`
	Status      consts.CaptchaStatus `json:"status"`
	Scenes      consts.CaptchaScenes `json:"scenes"` // 使用场景
}

func (Captcha) TableName() string {
	return "captcha"
}
