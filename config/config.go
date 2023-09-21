package config

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/spf13/viper"
	"path"
	"server/consts"
)

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
	// 判断文件夹是否存在
	if !fileutil.IsExist(Config.FileUploadDir) {
		// 创建文件夹
		fileutil.CreateDir(Config.FileUploadDir)
	}

	//fmt.Printf("Config: %+v\n", Config)
}

// GetLoggerOutPutPath 获取日志输出路径
func GetLoggerOutPutPath() string {
	if Config.IsProd {
		return Config.Log.Output
	}
	cPath := fileutil.CurrentPath()
	return path.Join(cPath, "../logs")
}
