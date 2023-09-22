package app

import (
	"github.com/gin-gonic/gin"
	"path"
	"server/config"
	"server/consts"
	"server/modules/api"
)

type Controller interface {
	New(gin *gin.Engine)
}

func New(r *gin.Engine) {
	r.Static(consts.UploadFilePathName, path.Join(config.Config.FileUploadDir))

	api.NewSwaggerController(r)

	api.NewHomeController(r)

	api.NewUserController(r)
	api.NewAccountController(r)
	api.NewAuthController(r)
	api.NewCaptchaController(r)

	api.NewAdminUserController(r)
	api.NewAdminAuthController(r)
	api.NewAdminRoleController(r)

	api.NewNewsController(r)

}
