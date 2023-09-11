package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/app"
	"server/config"
	"server/consts"
	_ "server/docs"
	"server/lib/logger"
	"server/lib/valid"
	"server/middleware"
	"server/utils"
	"strconv"
	"time"
)

// @title		男生自用 API 接口文档
// @version	1.0
func main() {
	fmt.Println("Version:", consts.Version)
	now := time.Now()

	// 配置
	config.New()
	var envName = utils.EnumLabel(consts.EnvDev)
	if config.Config.IsProd {
		envName = utils.EnumLabel(consts.EnvProd)
	}

	// 日志
	logger.New()

	if config.Config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin
	router := gin.New()
	// 全局中间件
	router.Use(middleware.RequestLogger(), gin.Recovery())

	// 启动
	app.New(router)

	// 验证器
	valid.New()

	go func() {
		fmt.Println("\nAPI", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port))
		fmt.Println("Swagger: ", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port)+"/swagger/index.html")
		//fmt.Printf("配置文件加载成功 %+v\n", config.Config)
		fmt.Println("当前环境", envName)
		fmt.Println("启动耗时: ", time.Now().Sub(now))
	}()

	// 监听端口
	err := router.Run(":" + strconv.Itoa(config.Config.Port))

	if err != nil {
		fmt.Println("启动失败!")
		panic(err)
	}
}
