package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/app"
	"server/config"
	_ "server/docs"
	"server/lib/logger"
	"server/lib/validate"
	"server/middleware"
	"strconv"
	"time"
)

// @title		男生自用 API 接口文档
// @version	1.0
func main() {
	now := time.Now()
	// 配置
	config.New()

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
	validate.InitValidate()

	fmt.Println("http://127.0.0.1:" + strconv.Itoa(config.Config.Port))
	fmt.Println("swagger: ", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port)+"/swagger/index.html")

	fmt.Println("启动耗时: ", time.Now().Sub(now))

	// 监听端口
	err := router.Run(":" + strconv.Itoa(config.Config.Port))

	if err != nil {
		fmt.Println("启动失败!")
		panic(err)
	}
}
