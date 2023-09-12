package api

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"server/consts"
	"server/lib/valid"
	"server/modules/service"
	"server/utils/resp"
)

type CaptchaController struct{}

func (c *CaptchaController) New(e *gin.Engine) {
	router := e.Group("/api/captcha")
	router.GET("", c.Create)
	router.GET("/image", c.CreateImage)
	router.GET("/verify/:id", c.Verify)
}

// Create
//
//	@Summary	获取验证码
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Param		type	query		string	true	"验证码类型， 1-手机 2-邮箱"	Enums(1,2)
//	@Param		value	query		string	true	"手机/邮箱账号"
//	@Param		scenes	query		string	true	"使用场景， 1-注册"	Enums(1)"
//	@Success	200		{object}	resp.Result{data=LoginSuccessResponse}	"请求结果"
//	@Router		/api/captcha [get]
func (c *CaptchaController) Create(ctx *gin.Context) {
	var query GetCaptchaQueryDto
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	if query.Value == "" {
		ctx.JSON(resp.ParamErr(consts.CaptchaTypeMap[query.Type] + "不能为空"))
		return
	}
	res := service.CaptchaService.Create(query.Type, query.Value, query.Scenes)
	img := captcha.NewImage(captcha.New(), []byte(res.Code), 100, 40)

	fmt.Printf("res: %v\n", img)
	ctx.JSON(resp.Succ(res))
}

// CreateImage
//
//	@Summary	获取图片验证码
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Success	200		{object}	resp.Result{data=string}	"请求结果"
//	@Router		/api/captcha/image [get]
func (c *CaptchaController) CreateImage(ctx *gin.Context) {
	ctx.JSON(resp.Succ(captcha.New()))
}

func (c *CaptchaController) Verify(ctx *gin.Context) {
	id := ctx.Param("id")

	captcha.WriteImage(ctx.Writer, id, consts.CaptchaWidth, consts.CaptchaHeight)
}

type GetCaptchaQueryDto struct {
	Type   consts.CaptchaType   `json:"type" form:"type" binding:"required"`
	Value  string               `json:"value" form:"value"`
	Scenes consts.CaptchaScenes `json:"scenes" form:"scenes"`
}
