package consts

type CaptchaType uint // 验证码类型
const (
	CaptchaTypePhone CaptchaType = 1 // 手机号
	CaptchaTypeEmail CaptchaType = 2 // 邮箱
	CaptchaTypeImage CaptchaType = 3 // 图片
)

type CaptchaStatus uint // 验证码状态
const (
	CaptchaStatusUnused CaptchaStatus = 0 // 未使用
	CaptchaStatusUsed   CaptchaStatus = 1 // 已使用
)
