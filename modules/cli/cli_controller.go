package cli

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"selfserver/config"
	"selfserver/utils/jsonutil"
)

type Controller struct {
	service Service
}

func Register(e *gin.Engine) {
	router := e.Group("/api/cli")
	controller := Controller{
		service: CreateService(),
	}
	fmt.Println("DEV", config.Config.IsDev)
	// 非开发环境禁止使用
	if !config.Config.IsDev {
		router.GET("", controller.NotUse)
		return
	}

	router.GET("/module/create", controller.CreateModule)
}

// NotUse
//
//	@Summary	禁止使用
func (c *Controller) NotUse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsonutil.SuccessResponse(nil, "生产环境禁止使用 cli 工具"))
}

// CreateModule godoc
//
//	@Summary		创建模块
//	@Description	get string by ID
//	@Tags			cli
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"Account ID"
//	@Success		200	{string}	string
//	@Router			/api/cli/module/create [get]
func (c *Controller) CreateModule(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsonutil.SuccessResponse([]string{"1", "2", "3"}, "success"))
}
