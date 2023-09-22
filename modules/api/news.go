package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/consts/permission"
	"server/dal/model"
	"server/lib/valid"
	"server/middleware"
	"server/modules/service"
	"server/utils"
	"server/utils/resp"
	"strconv"
	"time"
)

type NewsController struct {
	service *service.NewsService
}

func NewNewsController(e *gin.Engine) {
	c := &NewsController{
		service: service.NewNewsService(),
	}
	admin := e.Group("/api/admin/news")
	admin.GET("", middleware.AdminRole(permission.AdminNewsFind), c.FindList)
	admin.GET("/:id", middleware.AdminRole(permission.AdminNewsFindDetail), c.FindDetail)
	admin.POST("", middleware.AdminRole(permission.AdminNewsCreate), c.Create)
	admin.PUT("", middleware.AdminRole(permission.AdminNewsUpdateInfo), c.Update)
	admin.DELETE("/:id", middleware.AdminRole(permission.AdminNewsDelete), c.Delete)
}

// FindList
//
//	@Summary	新闻列表
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		current		query		number											false	"当前页"	default(1)
//	@Param		page_size	query		number											false	"每页条数"	default(20)
//	@Param		order_key	query		string											false	"需要排序的列"
//	@Param		order_type	query		string											false	"排序方式"	Enums(ase,desc)
//	@Param		keyword		query		string											false	"按名称或ID搜索"
//	@Success	200			{object}	resp.Result{data=resp.RList{list=[]model.News}}	"resp"
//	@Router		/api/admin/news [get]
func (c *NewsController) FindList(ctx *gin.Context) {
	query := service.FindNewsListOption{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	result, _ := c.service.FindList(query)
	ctx.JSON(resp.Succ(result))
}

// FindDetail
//
//	@Summary	新闻详情
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		id	path		number							true	"ID"
//	@Success	200	{object}	resp.Result{data=model.News}	"resp"
//	@Router		/api/admin/news/{id} [get]
func (c *NewsController) FindDetail(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	info, err := c.service.FindDetail(uint(id))
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(info))
}

// Create
//
//	@Summary	新增新闻
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateNewsBodyDto	true	"Body"
//	@Success	200	{object}	resp.Result			"resp"
//	@Router		/api/admin/news [post]
func (c *NewsController) Create(ctx *gin.Context) {
	var body CreateNewsBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	user := utils.GetCurrentAdminUser(ctx)

	data := model.News{
		Title:     body.Title,
		Desc:      body.Desc,
		Cover:     body.Cover,
		Content:   body.Content,
		Recommend: body.Recommend,
		PushDate:  time.Now(),
		AuthorID:  user.ID,
	}

	if body.PushDate != "" {
		t, terr := time.Parse("2006-01-02 15:04:05", body.PushDate)
		if body.PushDate != "" && terr != nil {
			ctx.JSON(resp.ParamErr("发布日期格式错误"))
			return
		}
		data.PushDate = t
	}

	if _, err := c.service.Create(data); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Update
//
//	@Summary	修改新闻
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		req	body		UpdateNewsBodyDto	true	"Body"
//	@Success	200	{object}	resp.Result			"resp"
//	@Router		/api/admin/news [put]
func (c *NewsController) Update(ctx *gin.Context) {
	var body UpdateNewsBodyDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	data := model.News{
		Title:     body.Title,
		Desc:      body.Desc,
		Cover:     body.Cover,
		Content:   body.Content,
		Recommend: body.Recommend,
		PushDate:  time.Now(),
	}

	if body.PushDate != "" {
		t, terr := time.Parse("2006-01-02 15:04:05", body.PushDate)
		if body.PushDate != "" && terr != nil {
			ctx.JSON(resp.ParamErr("发布日期格式错误"))
			return
		}
		data.PushDate = t
	}
	if _, err := c.service.Update(body.ID, data); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

// Delete
//
//	@Summary	删除新闻
//	@Tags		管理后台-新闻
//	@Accept		json
//	@Produce	json
//	@Param		id	path		number		true	"ID"
//	@Success	200	{object}	resp.Result	"resp"
//	@Router		/api/admin/news/{id} [delete]
func (c *NewsController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(resp.ParseErr(err))
		return
	}
	ctx.JSON(resp.Succ(nil))
}

type CreateNewsBodyDto struct {
	Title     string `json:"title" binding:"required" label:"新闻名称"`
	Desc      string `json:"desc"`
	Cover     string `json:"cover"`
	Content   string `json:"content"`
	PushDate  string `json:"push_date"` // 发布日期 YYYY-MM-DD HH:mm:ss
	Recommend uint   `json:"recommend"` // 推荐等级 0 为不推荐，数值越大越靠前
}

type UpdateNewsBodyDto struct {
	ID uint `json:"id"  binding:"required" label:"角色 ID"` // 角色 ID
	CreateNewsBodyDto
}
