package config

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/spf13/viper"
	"path"
	"server/consts"
)

type LogConfig struct {
	Output string // 日志路径
}

type AuthConfig struct {
	TokenSecret string // token 秘钥
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Port     uint32
	Dbname   string
}

type Result struct {
	Env    string     //  环境
	IsDev  bool       //  是否是开发环境
	IsProd bool       //  是否是生产环境
	Port   int        //  端口
	Log    LogConfig  //  日志
	Auth   AuthConfig //  用户认证
	Db     DBConfig   //  数据库配置
}

var Config Result

func New() {
	viper.SetConfigFile("config/config.prod.toml")
	viper.SetConfigType("toml")

	//判断文件是否存在
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("config/config.dev.toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic("配置文件不存在")
	}

	viper.Unmarshal(&Config)
	viper.WatchConfig()

	Config.IsDev = Config.Env == string(consts.EnvDev)
	Config.IsProd = Config.Env == string(consts.EnvProd)
}

// GetLoggerOutPutPath 获取日志输出路径
func GetLoggerOutPutPath() string {
	if Config.IsProd {
		return Config.Log.Output
	}
	cPath := fileutil.CurrentPath()
	return path.Join(cPath, "../logs")
}
