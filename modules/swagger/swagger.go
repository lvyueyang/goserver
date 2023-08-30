package swagger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
	"os/exec"
	"selfserver/config"
)

func Register(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RunCmd() {
	if config.Config.IsDev {
		fmt.Println("---------swagger 文档生成-------------------")
		cmd := exec.Command("swag", "init")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println("---------END -------------------")
	}
}
