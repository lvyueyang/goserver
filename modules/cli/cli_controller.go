package cli

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/config"
	"server/lib/types"
	"server/lib/valid"
	"server/utils/resp"
)

type Controller struct {
}

var New types.Controller = func(e *gin.Engine) {
	router := e.Group("/api/cli")
	controller := &Controller{}
	fmt.Println("DEV", config.Config.IsDev)
	// 非开发环境禁止使用
	if !config.Config.IsDev {
		router.GET("", controller.NotUse)
		return
	}

	router.POST("/module/create", controller.CreateModule)
}

// NotUse
//
//	@Summary	禁止使用
func (c *Controller) NotUse(ctx *gin.Context) {
	ctx.JSON(resp.Success(nil, "生产环境禁止使用 cli 工具"))
}

type CreateModuleBody struct {
	Name string `json:"name"` // 模块名称
}

// CreateModule godoc
//
//	@Summary		创建模块
//	@Description	get string by ID
//	@Tags			cli
//	@Accept			json
//	@Produce		json
//	@Param			data	body		CreateModuleBody			true	"模块名称"
//	@Success		200		{object}	resp.Result{data=bool}		"resp"
//	@Failure		500		{object}	resp.Result{data=string}	"resp"
//	@Router			/api/cli/module/create [post]
func (c *Controller) CreateModule(ctx *gin.Context) {
	var body CreateModuleBody
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	err := Service.CreateModule(body.Name)

	if err != nil {
		ctx.JSON(resp.ServerErr(err.Error(), "创建失败", 500))
		return
	}

	ctx.JSON(resp.SuccNil())
}
