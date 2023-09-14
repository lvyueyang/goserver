package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"os/signal"
	"server/app"
	"server/config"
	"server/consts"
	"server/db"
	_ "server/docs"
	"server/lib/logger"
	"server/lib/valid"
	"server/middleware"
	"server/utils"
	"strconv"
	"time"
)

//	@title		男生自用 API 接口文档
//	@version	1.0
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

	// 数据库
	db.New()
	defer db.Close()

	// gin
	router := gin.New()
	// 全局中间件
	router.Use(middleware.RequestLogger(), gin.Recovery())

	// 验证器
	valid.New()

	// 启动
	app.New(router)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Config.Port),
		Handler: router,
	}

	fmt.Println("\nAPI: ", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port))
	fmt.Println("Swagger: ", "http://127.0.0.1:"+strconv.Itoa(config.Config.Port)+"/swagger/index.html")
	//fmt.Printf("配置文件加载成功 %+v\n", config.Config)
	fmt.Println("当前环境: ", envName)
	fmt.Println("启动耗时: ", time.Now().Sub(now))

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen:", "err", err)
			panic(err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("服务关闭中...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("服务关闭:", "err", err)
		panic(err)
	}
	fmt.Println("服务退出")
}
