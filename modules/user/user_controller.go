package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	router.POST("/create", controller.Create)
}

// FindList godoc
//
//	@Summary	用户列表
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	string	"res"
//	@Router		/api/user/list [get]
func (c *Controller) FindList(ctx *gin.Context) {
	list := c.service.GetList()
	ctx.JSON(http.StatusOK, jsonutil.SuccessResponse(list, "success"))
}

type CreateUserBody struct {
	Name string `json:"name"`
}

// Create godoc
//
//	@Summary	创建用户
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		name	body		CreateUserBody	true	"ID"
//	@Success	200		{object}	nil	"res"
//	@Router		/api/user/create [post]
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateUserBody
	ctx.ShouldBindBodyWith(&body, binding.JSON)
	fmt.Printf("Create: %+v\n", body)
	ctx.JSON(http.StatusOK, jsonutil.SuccessResponse(nil, "success"))
}
