package swagger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
	"os/exec"
	"server/config"
	"server/lib/console"
)

func Register(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RunCmd() {
	if config.Config.IsDev {
		console.Success("swagger 文档生成")
		cmd := exec.Command("swag", "init")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("err:", err)
			console.Err("swagger 文档生成失败")
			return
		}
		console.Success("swagger 文档生成成功")
	}
}
