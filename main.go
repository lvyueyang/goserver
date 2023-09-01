package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/app"
	"server/config"
	_ "server/config"
	_ "server/docs"
	"server/lib/logs"
	"server/lib/validate"
	"server/middleware"
	"strconv"
)

// @title		男生自用 API 接口文档
// @version	1.0
func main() {
	// 配置
	config.Run()

	// 日志
	logs.InitLogger()

	// swagger
	//swagger.RunCmd()

	if config.Config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin
	router := gin.New()
	// 全局中间件
	router.Use(middleware.RequestLogger(), gin.Recovery())

	// 启动
	app.Run(router)

	// 验证器
	validate.InitValidate()

	fmt.Println("http://127.0.0.1:" + strconv.Itoa(config.Config.Port))
	fmt.Println("swagger: ", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port)+"/swagger/index.html")

	// 监听端口
	err := router.Run(":" + strconv.Itoa(config.Config.Port))

	if err != nil {
		fmt.Println("启动失败!")
		panic(err)
	}
}
