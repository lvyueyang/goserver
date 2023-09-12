package app

import (
	"github.com/gin-gonic/gin"
	"server/modules/api"
)

type Controller interface {
	New(gin *gin.Engine)
}

var modules = []Controller{
	new(api.SwaggerController),

	new(api.HomeController),
	new(api.UserController),
	new(api.AccountController),
	new(api.AuthController),
	new(api.CaptchaController),
}

func New(r *gin.Engine) {
	for _, module := range modules {
		module.New(r)
	}
}
