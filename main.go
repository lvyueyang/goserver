package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"selfserver/app"
	"selfserver/config"
	"strconv"
	"time"

	_ "selfserver/config"
)

func main() {
	// 配置
	config.Run()

	// 日志
	gin.DisableConsoleColor()
	now := time.Now()
	f, _ := os.Create(path.Join("logs/", now.Format("2006-01-02")+".log"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// gin
	router := gin.Default()

	// 启动
	app.Run(router)

	fmt.Println("http://127.0.0.1:" + strconv.Itoa(config.Config.Port))

	// 监听端口
	err := router.Run(":" + strconv.Itoa(config.Config.Port))

	if err != nil {
		fmt.Println("启动失败!")
		panic(err)
	}
}
