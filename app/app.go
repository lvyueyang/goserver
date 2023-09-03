package app

import (
	"github.com/gin-gonic/gin"
	"server/modules/cli"
	"server/modules/home"
	"server/modules/swagger"
	"server/modules/user"
)

var modules = []func(e *gin.Engine){
	cli.New,
	swagger.New,
	user.New,
	home.New,
}

func Run(r *gin.Engine) {
	for _, module := range modules {
		module(r)
	}
}
