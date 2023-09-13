package consts

type CaptchaType uint // 验证码类型
const (
	CaptchaTypePhone CaptchaType = 1 // 手机号
	CaptchaTypeEmail CaptchaType = 2 // 邮箱
)

var CaptchaTypeMap = map[CaptchaType]string{
	CaptchaTypePhone: "手机号",
	CaptchaTypeEmail: "邮箱",
}

type CaptchaStatus uint // 验证码状态
const (
	CaptchaStatusUnused CaptchaStatus = 0 // 未使用
	CaptchaStatusUsed   CaptchaStatus = 1 // 已使用
)

type CaptchaScenes uint // 验证码使用场景
const (
	CaptchaScenesRegister CaptchaScenes = 1 // 注册
)

const (
	CaptchaWidth  = 200
	CaptchaHeight = 72
)
