package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHome(r *gin.Engine) {
	r.GET("/", get)
}

func get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
