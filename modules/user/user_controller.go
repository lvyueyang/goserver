package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/lib/validate"
	"server/utils/resp"
)

type Controller struct {
}

func New(e *gin.Engine) {
	router := e.Group("/api/user")
	controller := &Controller{}
	router.GET("/list", controller.FindList)
	router.POST("/create", controller.Create)
}

// FindList
//
//	@Summary	用户列表
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	string	"resp"
//	@Router		/api/user/list [get]
func (c *Controller) FindList(ctx *gin.Context) {
	list := Service.GetList()
	ctx.JSON(resp.Succ(list))
}

// Create
//
//	@Summary	创建用户
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateUserBodyDto	true	"body"
//	@Success	200	{object}	nil					"resp"
//	@Router		/api/user/create [post]
func (c *Controller) Create(ctx *gin.Context) {
	var body CreateUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(validate.ErrTransform(err)))
		return
	}
	ctx.JSON(resp.Succ(nil))
}
