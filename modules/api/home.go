package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct{}

func NewHomeController(e *gin.Engine) {
	c := &HomeController{}
	router := e.Group("/")
	router.GET("/", c.HomePage)
}

// HomePage 主页
func (c *HomeController) HomePage(ctx *gin.Context) {
	ctx.String(http.StatusOK, "男生自用")
}
