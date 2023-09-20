package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/consts/permission"
	"server/dal/dbtypes"
	"server/dal/model"
	"server/lib/valid"
	"server/middleware"
	"server/modules/service"
	"server/utils/resp"
	"strconv"
)

type AdminRoleController struct {
	service *service.AdminRoleService
}

func NewAdminRoleController(e *gin.Engine) {
	c := &AdminRoleController{
		service: service.NewAdminRoleService(),
	}
	admin := e.Group("/api/admin/role")
	admin.GET("/permission/codes", c.FindPermissionCodes)
	admin.GET("", middleware.AdminRole(permission.AdminRoleFind), c.FindList)
	admin.POST("", middleware.AdminRole(permission.AdminRoleCreate), c.Create)
	admin.PUT("", middleware.AdminRole(permission.AdminRoleUpdateInfo), c.Update)
	admin.DELETE("/:id", middleware.AdminRole(permission.AdminRoleDelete), c.Delete)
	admin.PUT("/permission/codes", middleware.AdminRole(permission.AdminRoleUpdateCode), c.UpdatePermissionCodes)
}

// FindList
//
//	@Summary	管理员角色列表
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		current		query		number													false	"当前页"	default(1)
//	@Param		page_size	query		number													false	"每页条数"	default(20)
//	@Param		order_key	query		string													false	"需要排序的列"
//	@Param		order_type	query		string													false	"排序方式"	Enums(ase,desc)
//	@Param		keyword		query		string													false	"按用户名搜索"
//	@Success	200			{object}	resp.Result{data=resp.RList{list=[]model.AdminRole}}	"resp"
//	@Router		/api/admin/role [get]
func (c *AdminRoleController) FindList(ctx *gin.Context) {
	query := service.FindAdminRoleListOption{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	result, _ := c.service.FindList(query)
	ctx.JSON(resp.Succ(result))
}

// Create
//
//	@Summary	新增管理员角色
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateAdminRoleBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/role [post]
func (c *AdminRoleController) Create(ctx *gin.Context) {
	var body CreateAdminRoleBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	if _, err := c.service.Create(model.AdminRole{
		Name: body.Name,
		Code: body.Code,
		Desc: body.Desc,
	}); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Delete
//
//	@Summary	删除角色
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		id	path		number		true	"角色 ID"
//	@Success	200	{object}	resp.Result	"resp"
//	@Router		/api/admin/role/{id} [delete]
func (c *AdminRoleController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Update
//
//	@Summary	修改管理员角色
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		req	body		UpdateAdminRoleBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/role [put]
func (c *AdminRoleController) Update(ctx *gin.Context) {
	var body UpdateAdminRoleBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	if _, err := c.service.Update(body.ID, model.AdminRole{
		Name: body.Name,
		Code: body.Code,
		Desc: body.Desc,
	}); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// UpdatePermissionCodes
//
//	@Summary	修改管理员角色权限
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Param		req	body		UpdateAdminRolePermissionBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result							"resp"
//	@Router		/api/admin/role/permission/codes [put]
func (c *AdminRoleController) UpdatePermissionCodes(ctx *gin.Context) {
	var body UpdateAdminRolePermissionBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	for _, code := range body.Codes {
		if permission.AdminLabelMap[code].Label == "" {
			ctx.JSON(resp.ParamErr("无效的权限码: " + code))
			return
		}
	}

	if _, err := c.service.UpdatePermissionCode(body.ID, body.Codes); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// FindPermissionCodes
//
//	@Summary	管理后台权限码列表
//	@Tags		管理后台-管理员角色
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=map[string]permission.LabelType}	"resp"
//	@Router		/api/admin/role/permission/codes [get]
func (c *AdminRoleController) FindPermissionCodes(ctx *gin.Context) {
	ctx.JSON(resp.Succ(permission.AdminLabelMap))
}

type CreateAdminRoleBodyDto struct {
	Name string `json:"name" binding:"required" label:"角色名称"` // 姓名
	Code string `json:"code" binding:"required" label:"角色编码"` // 用户名
	Desc string `json:"desc"`                                 // 描述
}

type UpdateAdminRoleBodyDto struct {
	ID uint `json:"id"  binding:"required" label:"角色 ID"` // 角色 ID
	CreateAdminRoleBodyDto
}

type UpdateAdminRolePermissionBodyDto struct {
	ID    uint                `json:"id"  binding:"required" label:"角色 ID"` // 角色 ID
	Codes dbtypes.StringArray `json:"codes" label:"权限码"`                    // 权限码
}
