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

type NewsController struct {
	service *service.NewsService
}

func NewNewsController(e *gin.Engine) {
	c := &NewsController{
		service: service.NewNewsService(),
	}
	admin := e.Group("/api/admin/role")
	admin.GET("/permission/codes", c.FindPermissionCodes)
	admin.GET("", middleware.News(permission.NewsFind), c.FindList)
	admin.POST("", middleware.News(permission.NewsCreate), c.Create)
	admin.PUT("", middleware.News(permission.NewsUpdateInfo), c.Update)
	admin.DELETE("/:id", middleware.News(permission.NewsDelete), c.Delete)
	admin.PUT("/permission/codes", middleware.News(permission.NewsUpdateCode), c.UpdatePermissionCodes)
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
//	@Success	200			{object}	resp.Result{data=resp.RList{list=[]model.News}}	"resp"
//	@Router		/api/admin/role [get]
func (c *NewsController) FindList(ctx *gin.Context) {
	query := service.FindNewsListOption{}
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
//	@Param		req	body		CreateNewsBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/role [post]
func (c *NewsController) Create(ctx *gin.Context) {
	var body CreateNewsBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	if _, err := c.service.Create(model.News{
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
func (c *NewsController) Delete(ctx *gin.Context) {
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
//	@Param		req	body		UpdateNewsBodyDto	true	"管理员信息"
//	@Success	200	{object}	resp.Result				"resp"
//	@Router		/api/admin/role [put]
func (c *NewsController) Update(ctx *gin.Context) {
	var body UpdateNewsBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	if _, err := c.service.Update(body.ID, model.News{
		Name: body.Name,
		Code: body.Code,
		Desc: body.Desc,
	}); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

type CreateNewsBodyDto struct {
	Name string `json:"name" binding:"required" label:"角色名称"` // 姓名
	Code string `json:"code" binding:"required" label:"角色编码"` // 用户名
	Desc string `json:"desc"`                                 // 描述
}

type UpdateNewsBodyDto struct {
	ID uint `json:"id"  binding:"required" label:"角色 ID"` // 角色 ID
	CreateNewsBodyDto
}

type UpdateNewsPermissionBodyDto struct {
	ID    uint                `json:"id"  binding:"required" label:"角色 ID"` // 角色 ID
	Codes dbtypes.StringArray `json:"codes" label:"权限码"`                    // 权限码
}
