package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"selfserver/db"
	"selfserver/modules/cli"
	"selfserver/modules/swagger"
)

type Application struct {
	DB *gorm.DB
}

var App Application

func Run(r *gin.Engine) {
	// 数据库链接
	App.DB = db.Connect()

	cli.Register(r)

	swagger.Register(r)
}
