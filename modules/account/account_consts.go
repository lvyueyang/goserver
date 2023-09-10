package account

type Type uint

const (
	NormalAccountType Type = 1 // 普通账户,使用用户名和密码登录
	EmailAccountType  Type = 2 // 邮箱账户
	WxMpAccountType   Type = 3 // 微信小程序账户
)
