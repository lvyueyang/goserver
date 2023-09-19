package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/consts"
	"server/dal/model"
	"server/lib/valid"
	"server/middleware"
	"server/modules/service"
	"server/utils/resp"
	"strconv"
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
	admin.GET("/current", middleware.AdminAuth(), c.CurrentInfo)
	admin.POST("", c.Create)
	admin.POST("/:id", c.Detail)
	admin.PUT("/:id", c.Update)
	admin.PUT("/reset-password/:id", c.ResetPassword)
	admin.PUT("/status/:id", c.UpdateStatus)
}

// FindList
//
//	@Summary	管理员列表
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		current		query		number													false	"当前页"	default(1)
//	@Param		page_size	query		number													false	"每页条数"	default(20)
//	@Param		order_key	query		string													false	"需要排序的列"
//	@Param		order_type	query		string													false	"排序方式"	Enums(ase,desc)
//	@Param		keyword		query		string													false	"按用户名搜索"
//	@Success	200			{object}	resp.Result{data=resp.RList{list=[]model.AdminUser}}	"resp"
//	@Router		/api/admin/user [get]
func (c *AdminUserController) FindList(ctx *gin.Context) {
	query := service.FindUserListOption{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	result, _ := c.service.FindList(query)
	ctx.JSON(resp.Succ(result))
}

// Create
//
//	@Summary	新增管理员
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/user [post]
func (c *AdminUserController) Create(ctx *gin.Context) {
	var body CreateAdminUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	if _, err := c.service.Create(model.AdminUser{
		Name:     body.Name,
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
		Avatar:   body.Avatar,
	}); err != nil {
		ctx.JSON(resp.ParseErr(err))
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
//	@Summary	修改管理员信息
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		UpdateAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/user/{id} [put]
func (c *AdminUserController) Update(ctx *gin.Context) {
	var body UpdateAdminUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Update(uint(id), model.AdminUser{
		Name:   body.Name,
		Avatar: body.Avatar,
	}); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// UpdateStatus
//
//	@Summary	修改管理员状态(封禁/解封)
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		UpdateAdminUserStatusBodyDto	true	"用户状态"
//	@Success	200	{object}	resp.Result						"resp"
//	@Router		/api/admin/user/status/{id} [put]
func (c *AdminUserController) UpdateStatus(ctx *gin.Context) {
	var body UpdateAdminUserStatusBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.UpdateStatus(uint(id), body.Status); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// ResetPassword
//
//	@Summary	重置管理员密码
//	@Tags		管理后台-管理员用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		ResetPasswordAdminUserBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result						"resp"
//	@Router		/api/admin/user/reset-password/{id} [put]
func (c *AdminUserController) ResetPassword(ctx *gin.Context) {
	var body ResetPasswordAdminUserBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.UpdatePassword(uint(id), body.Password); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
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
	Avatar   string `json:"avatar"`                                  // 头像
}

type UpdateAdminUserBodyDto struct {
	Name   string `json:"name" binding:"required" label:"姓名"` // 姓名
	Avatar string `json:"avatar"`                             // 头像
}

type UpdateAdminUserStatusBodyDto struct {
	Status consts.AdminUserStatus `json:"status" binding:"required" label:"用户状态" enums:"-1,1"` // 状态 1-解封 2-封禁
}
type ResetPasswordAdminUserBodyDto struct {
	Password string `json:"password" binding:"required" label:"密码"` // 密码
}
