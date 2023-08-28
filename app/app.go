package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"selfserver/api/v1"
	"selfserver/db"
)

type Application struct {
	DB *gorm.DB
}

var App Application

func Run(r *gin.Engine) {
	// 数据库链接
	App.DB = db.Connect()

	v1.CreateHome(r)
	v1.CreateUser(r)
}
