package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	service Service
}

func New(e *gin.Engine) {
	router := e.Group("/")
	controller := Controller{
		service: ServiceInstance,
	}
	router.GET("/", controller.HomePage)
}

// HomePage 主页
func (c *Controller) HomePage(ctx *gin.Context) {
	ctx.String(http.StatusOK, "男生自用")
}
