package auth

import (
	"github.com/gin-gonic/gin"
	"server/utils/resp"
)

type Controller struct {
	service *Service
}

func New(e *gin.Engine) {
	router := e.Group("/api/auth")
	controller := &Controller{
		service: ServiceInstance,
	}
	router.GET("/list", controller.FindList)
}

// FindList godoc
//
//	@Summary		列表
//	@Description	获取列表
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	jsonutil.ResponseData	"res"
//	@Router			/api/auth/create [get]
func (c *Controller) FindList(ctx *gin.Context) {
	ctx.JSON(resp.Succ(nil))
}
