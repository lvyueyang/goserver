package app

import (
	"github.com/gin-gonic/gin"
	"server/modules/api"
)

type Controller interface {
	New(gin *gin.Engine)
}

func New(r *gin.Engine) {
	api.NewSwaggerController(r)

	api.NewHomeController(r)

	api.NewUserController(r)
	api.NewAccountController(r)
	api.NewAuthController(r)
	api.NewCaptchaController(r)
}
