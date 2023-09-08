package {{.Name}}

import (
    "github.com/gin-gonic/gin"
    "server/utils/resp"
)

type Controller struct {
    service *Service
}

func New(e *gin.Engine) {
    router := e.Group("/api/{{.Name}}")
    controller := &Controller{
        service: ServiceInstance,
    }
    router.GET("", controller.FindList)
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
    ctx.JSON(resp.Succ(nil))
}
