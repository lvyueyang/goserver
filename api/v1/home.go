package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfserver/service"
)

func CreateHome(r *gin.Engine) {
	controller := useHomeController()
	r.GET("/", controller.get)
}

type HomeController struct {
	service service.HomeService
}

func useHomeController() HomeController {
	return HomeController{
		service: service.HomeServiceInstance,
	}
}

func (c *HomeController) get(ctx *gin.Context) {
	ctx.String(http.StatusOK, c.service.GetInfo())
}
