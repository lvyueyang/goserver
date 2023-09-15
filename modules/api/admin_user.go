package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"server/consts"
	"server/dal/model"
	"server/lib/valid"
	"server/middleware"
	"server/modules/service"
	"server/utils/resp"
)

type AdminUserController struct {
	service *service.AdminUserService
}

func NewAdminUserController(e *gin.Engine) {
	c := &AdminUserController{
		service: service.NewAdminUserService(),
	}
	admin := e.Group("/api/admin/user")
	admin.GET("", c.FindList)
	admin.POST("", c.Create)
	admin.POST("/:id", c.Detail)
	admin.PUT("/:id", c.Update)
	admin.PUT("/reset-password/:id", c.ResetPassword)
	admin.GET("/current", middleware.AdminAuth(), c.CurrentInfo)
}

// FindList
//
//	@Summary	管理员列表
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	string	"resp"
//	@Router		/api/admin/user [get]
func (c *AdminUserController) FindList(ctx *gin.Context) {
	list, _ := c.service.FindList(service.FindUserListOption{})
	ctx.JSON(http.StatusOK, list)
}

// Create
//
//	@Summary	管理员新增
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	string					"resp.Result"
//	@Router		/api/admin/user [post]
func (c *AdminUserController) Create(ctx *gin.Context) {
	var body CreateAdminUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Detail
//
//	@Summary	管理员详情
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateAdminUserBodyDto				true	"管理员信息"
//	@Success	200	{object}	resp.Result{data=model.AdminUser}	"用户详情"
//	@Router		/api/admin/user/{id} [get]
func (c *AdminUserController) Detail(ctx *gin.Context) {
	ctx.JSON(resp.Succ(nil))
}

// Update
//
//	@Summary	管理员修改
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		UpdateAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	string					"resp.Result"
//	@Router		/api/admin/user/{id} [put]
func (c *AdminUserController) Update(ctx *gin.Context) {
	ctx.JSON(resp.Succ(nil))
}

// ResetPassword
//
//	@Summary	重置管理员密码
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		ResetPasswordAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	string							"resp.Result"
//	@Router		/api/admin/user/reset-password/{id} [put]
func (c *AdminUserController) ResetPassword(ctx *gin.Context) {
	ctx.JSON(resp.Succ(nil))
}

// CurrentInfo
//
//	@Summary	当前登陆者信息
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=model.AdminUser}	"用户详情"
//	@Router		/api/admin/user/current [get]
func (c *AdminUserController) CurrentInfo(ctx *gin.Context) {
	user := ctx.MustGet(consts.ContextUserInfoKey).(*model.AdminUser)
	ctx.JSON(resp.Succ(user))
}

type CreateAdminUserBodyDto struct {
	Name     string `json:"name" binding:"required" label:"姓名"`      // 姓名
	Username string `json:"username" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required" label:"密码"`  // 密码
	Email    string `json:"email" binding:"required" label:"邮箱"`     // 邮箱
}

type UpdateAdminUserBodyDto struct {
	Name string `json:"name" binding:"required" label:"姓名"` // 姓名
}

type ResetPasswordAdminUserBodyDto struct {
	Password string `json:"password" binding:"required" label:"密码"` // 密码
}
