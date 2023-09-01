package {{.Name}}

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "server/utils/jsonutil"
)

type Controller struct {
    service Service
}

func Register(e *gin.Engine) {
    router := e.Group("/api/{{.Name}}")
    controller := Controller{
        service: ServiceInstance,
    }
    router.GET("/list", controller.FindList)
}

// FindList godoc
//	@Summary		列表
//	@Description	获取列表
//	@Tags			{{.Name}}
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	jsonutil.ResponseData	"res"
//	@Router			/api/{{.Name}}/create [get]
func (c *Controller) FindList(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, jsonutil.SuccessResponse(nil, "success"))
}
