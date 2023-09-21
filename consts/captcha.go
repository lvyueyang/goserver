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
	CaptchaScenesRegister       CaptchaScenes = 1 // 注册
	CaptchaScenesForgetPassword CaptchaScenes = 2 // 密码重置
	CaptchaScenesUpdatePhone    CaptchaScenes = 3 // 更新手机号
	CaptchaScenesUpdateEmail    CaptchaScenes = 4 // 更新邮箱
)

type CaptchaScenesValue struct {
	Label      string
	EmailTitle string
}

var CaptchaScenesMap = map[CaptchaScenes]CaptchaScenesValue{
	CaptchaScenesRegister: {
		Label:      "注册",
		EmailTitle: "您正在注册账号",
	},
	CaptchaScenesForgetPassword: {
		Label:      "密码重置",
		EmailTitle: "您正在对您的账号进行密码重置",
	},
	CaptchaScenesUpdatePhone: {
		Label:      "更新手机号",
		EmailTitle: "您正在更新您的手机号",
	},
	CaptchaScenesUpdateEmail: {
		Label:      "更新邮箱",
		EmailTitle: "您正在更新邮箱账号",
	},
}

const (
	CaptchaWidth  = 200
	CaptchaHeight = 72
)
