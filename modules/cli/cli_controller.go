package cli

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"selfserver/config"
	"selfserver/utils/jsonutil"
	"text/template"
)

type Controller struct {
	service Service
}

func Register(e *gin.Engine) {
	router := e.Group("/api/cli")
	controller := Controller{
		service: ServiceInstance,
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
//	@Param			name	query		string								true	"模块名称"
//	@Success		200		{object}	jsonutil.ResponseData{data=bool}	"res"
//	@Router			/api/cli/module/create [get]
func (c *Controller) CreateModule(ctx *gin.Context) {
	name := ctx.Query("name")
	dir := path.Join("modules", name)
	// 创建文件夹
	creatDirErr := os.Mkdir(dir, 0755)
	if creatDirErr != nil {
		ctx.JSON(http.StatusBadRequest, jsonutil.ErrorResponse(nil, "创建文件夹失败", 500))
		return
	}
	modNames := []string{"controller", "service", "model"}

	for _, modName := range modNames {
		filePath := path.Join(dir, name+"_"+modName+".go")
		tempFile, _ := template.ParseFiles("modules/cli/template/" + modName + ".tpl")
		os.WriteFile(filePath, []byte(""), 0755)
		file, _ := os.OpenFile(filePath, os.O_RDWR, 0755)
		err := tempFile.Execute(file, struct {
			Name string
		}{Name: name})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, jsonutil.ErrorResponse(nil, "创建"+filePath+"失败", 500))
			return
		}
	}

	ctx.JSON(http.StatusOK, jsonutil.SuccessResponse(nil, "success"))
}
