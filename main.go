package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"selfserver/app"
	"selfserver/config"
	_ "selfserver/config"
	_ "selfserver/docs"
	"selfserver/lib/logger"
	"selfserver/middleware"
	"selfserver/modules/swagger"
	"strconv"
)

// @title		男生自用 API 接口文档
// @version	1.0
func main() {
	// 日志
	logger.InitLogger()
	defer logger.Logger.Sync()

	// swagger
	swagger.RunCmd()

	// 配置
	config.Run()

	if config.Config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin
	router := gin.New()
	// 全局中间件
	router.Use(middleware.Logger(), gin.Recovery())

	// 启动
	app.Run(router)

	fmt.Println("http://127.0.0.1:" + strconv.Itoa(config.Config.Port))
	fmt.Println("swagger: ", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port)+"/swagger/index.html")

	// 监听端口
	err := router.Run(":" + strconv.Itoa(config.Config.Port))

	if err != nil {
		fmt.Println("启动失败!")
		panic(err)
	}
}
