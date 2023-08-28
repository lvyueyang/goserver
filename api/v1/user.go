package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfserver/utils"
)

/*
CreateUser 用户相关
*/
func CreateUser(r *gin.Engine) {
	user := r.Group("/user")

	user.GET("", getList)
}

func getList(c *gin.Context) {
	c.JSON(http.StatusOK, utils.SuccessResponse([]string{"1", "2", "3"}, "ok"))
}
