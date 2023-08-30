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
