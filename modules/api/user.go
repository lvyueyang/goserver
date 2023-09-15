package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/lib/valid"
	"server/modules/service"
	"server/utils/resp"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(e *gin.Engine) {
	c := &UserController{
		service: service.NewUserService(),
	}
	admin := e.Group("/api/admin/user")
	admin.GET("/list", c.FindList)
}

// FindList
//
//	@Summary	用户列表
//	@Tags		管理后台-用户
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	string	"resp"
//	@Router		/api/admin/user/list [get]
func (c *UserController) FindList(ctx *gin.Context) {
	list, _ := c.service.FindList()
	ctx.JSON(resp.Succ(list))
}

// Create 创建用户
func (c *UserController) Create(ctx *gin.Context) {
	var body CreateUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

type CreateUserBodyDto struct {
	Name string `json:"name" binding:"required" label:"姓名"` // 姓名
	Sex  string `json:"sex" binding:"required" label:"性别"`  // 性别
}
