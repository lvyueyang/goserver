package app

import (
	"github.com/gin-gonic/gin"
	"server/internal/controller"
)

type Controller interface {
	New(gin *gin.Engine)
}

var modules = []Controller{
	&controller.HomeController{},
}

func New(r *gin.Engine) {
	for _, module := range modules {
		module.New(r)
	}
}
