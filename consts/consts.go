package consts

// Env 运行环境
type Env string

const (
	EnvDev  Env = "development"
	EnvProd Env = "production"
)

func (e Env) Label() string {
	switch e {
	case EnvDev:
		return "开发环境"
	case EnvProd:
		return "生产环境"
	}
	return "开发环境"
}

// Sex 性别
type Sex int

const (
	SexUnknown Sex = 0 // 未知
	SexMan     Sex = 1 // 男
	SexWoman   Sex = 2 // 女
)

func (g Sex) Label() string {
	switch g {
	case SexMan:
		return "男"
	case SexWoman:
		return "女"
	case SexUnknown:
		return "未知"
	}
	return "未知"
}
