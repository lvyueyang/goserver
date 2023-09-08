package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerStruct struct{}

func New(e *gin.Engine) {
	router := e.Group("/")
	controller := &ControllerStruct{}
	router.GET("/", controller.HomePage)
}

// HomePage 主页
func (c *ControllerStruct) HomePage(ctx *gin.Context) {
	ctx.String(http.StatusOK, "男生自用")
}
