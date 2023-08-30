package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"selfserver/app"
	"selfserver/config"
	_ "selfserver/config"
	_ "selfserver/docs"
	"selfserver/middleware/logger"
	"selfserver/modules/swagger"
	"strconv"
)

// @title		男生自用 API 接口文档
// @version	1.0
func main() {
	// swagger
	swagger.RunCmd()

	// 配置
	config.Run()

	if config.Config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// 日志
	//gin.DisableConsoleColor()
	//now := time.Now()
	//f, _ := os.Create(path.Join("logs/request", now.Format("2006-01-02")+".log"))
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// gin
	router := gin.New()
	// 全局中间件
	router.Use(logger.Logger(), gin.Recovery())

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
