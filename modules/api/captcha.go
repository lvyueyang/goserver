package api

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/consts"
	"server/lib/valid"
	"server/modules/service"
	"server/utils/resp"
)

type CaptchaController struct {
	service *service.CaptchaService
}

func NewCaptchaController(e *gin.Engine) {
	c := CaptchaController{
		service: service.NewCaptchaService(),
	}

	router := e.Group("/api/captcha")
	router.POST("", c.Create)
	router.GET("/image", c.CreateImage)
	router.GET("/verify/:id", c.Verify)
}

// Create
//
//	@Summary	发送验证码
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateCaptchaReqDto			true	"body"
//	@Success	200	{object}	resp.Result{}	"请求结果"
//	@Router		/api/captcha [post]
func (c *CaptchaController) Create(ctx *gin.Context) {
	var body CreateCaptchaReqDto
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		fmt.Printf("%+v\n", err)
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	// 校验验证码
	if succ := captcha.Verify(body.CaptchaKey, []byte(body.CaptchaValue)); !succ {
		ctx.JSON(resp.ParamErr("图形验证码不正确"))
		return
	}

	// 创建邮箱验证码
	c.service.Create(body.Type, body.Value, body.Scenes)

	ctx.JSON(resp.Succ(nil))
}

// CreateImage
//
//	@Summary	获取图片验证码
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=string}	"请求结果"
//	@Router		/api/captcha/image [get]
func (c *CaptchaController) CreateImage(ctx *gin.Context) {
	ctx.JSON(resp.Succ(captcha.New()))
}

func (c *CaptchaController) Verify(ctx *gin.Context) {
	id := ctx.Param("id")

	captcha.WriteImage(ctx.Writer, id, consts.CaptchaWidth, consts.CaptchaHeight)
}

type CreateCaptchaReqDto struct {
	Type         consts.CaptchaType   `json:"type" binding:"required"`                        // 验证码类型， 1-手机 2-邮箱 Enums(1,2)
	Value        string               `json:"value" binding:"required" label:"手机/邮箱账号"`       // 手机/邮箱账号
	Scenes       consts.CaptchaScenes `json:"scenes" binding:"required"`                      // 使用场景， 1-注册
	CaptchaKey   string               `json:"captcha_key" binding:"required"`                 // 图形验证码的key
	CaptchaValue string               `json:"captcha_value" binding:"required" label:"图形验证码"` // 输入的图形验证码
}
