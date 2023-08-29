package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"selfserver/utils/jsonutil"
)

type Controller struct {
	service Service
}

func Register(e *gin.Engine) {
	router := e.Group("/api/user")
	controller := Controller{
		service: ServiceInstance,
	}
	router.GET("/list", controller.FindList)
}

// FindList godoc
//
//	@Summary	用户列表
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	jsonutil.ResponseData{data=[]user.User}	"res"
//	@Router		/api/user/list [get]
func (c *Controller) FindList(ctx *gin.Context) {
	list := c.service.GetList()
	ctx.JSON(http.StatusOK, jsonutil.SuccessResponse(list, "success"))
}
