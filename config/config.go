package config

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/spf13/viper"
	"path"
	"selfserver/consts"
	"selfserver/lib/console"
	"selfserver/utils"
)

type LogConfig struct {
	Output string // 日志路径
}

type Result struct {
	Env    string    //  环境
	IsDev  bool      //  是否是开发环境
	IsProd bool      //  是否是生产环境
	Port   int       //  端口
	Log    LogConfig // 日志
}

var Config Result

func Run() {
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

	var envName = utils.EnumLabel(consts.EnvDev)
	if Config.IsProd {
		envName = utils.EnumLabel(consts.EnvProd)
	}

	console.Success("配置文件加载成功")
	fmt.Printf("%+v\n", Config)
	console.Success("当前环境：", envName)
}

// GetLoggerOutPutPath 获取日志输出路径
func GetLoggerOutPutPath() string {
	if Config.IsProd {
		return Config.Log.Output
	}
	cPath := fileutil.CurrentPath()
	return path.Join(cPath, "../logs")
}
