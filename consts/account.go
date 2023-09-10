package consts

type AccountType uint // 账号类型

const (
	NormalAccountType AccountType = 1 // 普通账户,使用用户名和密码登录
	EmailAccountType  AccountType = 2 // 邮箱账户
	WxMpAccountType   AccountType = 3 // 微信小程序账户
)
