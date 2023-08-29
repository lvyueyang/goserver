package app

import (
	"github.com/gin-gonic/gin"
	"selfserver/modules/cli"
	"selfserver/modules/home"
	"selfserver/modules/swagger"
	"selfserver/modules/user"
)

type Application struct {
}

var App Application

func Run(r *gin.Engine) {
	// 数据库链接

	cli.Register(r)
	swagger.Register(r)

	home.Register(r)
	user.Register(r)
}
