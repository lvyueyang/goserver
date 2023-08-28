package config

import (
	"fmt"
	"github.com/spf13/viper"
	"selfserver/consts"
	"selfserver/utils"
)

type Result struct {
	IsDev  bool //  是否是开发环境
	IsProd bool //  是否是生产环境
	Port   int  // 	端口
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

	Config.IsDev = viper.Get("ENV") == consts.EnvDev
	Config.IsProd = viper.Get("ENV") == consts.EnvDev
	Config.Port = viper.GetInt("PORT")

	fmt.Println("配置文件加载成功", Config)

	var envName = utils.EnumLabel(consts.EnvDev)
	if Config.IsProd {
		envName = utils.EnumLabel(consts.EnvProd)
	}
	fmt.Println("当前环境：", envName)
}
