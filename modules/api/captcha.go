package api

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/robfig/cron/v3"
	"golang.org/x/exp/slog"
	"server/consts"
	"server/lib/valid"
	"server/middleware"
	"server/modules/service"
	"server/utils/resp"
)

type CaptchaController struct {
	service       *service.CaptchaService
	notifyService *service.NotifyService
}

func NewCaptchaController(e *gin.Engine) {
	c := CaptchaController{
		service:       service.NewCaptchaService(),
		notifyService: service.NewNotifyService(),
	}

	router := e.Group("/api/captcha")
	router.POST("", c.Create)
	router.GET("/image", c.CreateImage)
	router.GET("/image/:key", c.ImageFile)
	router.GET("/clear", middleware.Auth(), c.Clear)

	c.service.ClearExpiration()
	cr := cron.New()
	// 每隔五分钟清除一次过期验证码
	cr.AddFunc("@every 5m", func() {
		c.service.ClearExpiration()
	})
	cr.Start()
}

// Create
//
//	@Summary	发送验证码
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Param		req	body		CreateCaptchaBodyDto	true	"body"
//	@Success	200	{object}	resp.Result{data=int}	"验证码 ID"
//	@Router		/api/captcha [post]
func (c *CaptchaController) Create(ctx *gin.Context) {
	var body = new(CreateCaptchaBodyDto)
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		fmt.Printf("%+v\n", err)
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	// 校验验证码
	if succ := captcha.VerifyString(body.CaptchaKey, body.CaptchaValue); !succ {
		ctx.JSON(resp.ParamErr("图形验证码不正确"))
		return
	}

	// 创建邮箱验证码
	res, err := c.service.Create(body.Type, body.Value, body.Scenes)
	if err != nil {
		ctx.JSON(resp.ParamErr("验证码创建失败"))
	}

	go func() {
		if body.Type == consts.CaptchaTypeEmail {
			cru := consts.CaptchaScenesMap[body.Scenes]
			if err := c.notifyService.SendCaptchaEmail(cru.Label, cru.EmailTitle, body.Value, res.Code); err != nil {
				slog.Error("邮件发送失败", "email", body.Value, "err", err)
			}
		}
	}()

	ctx.JSON(resp.Succ(res.ID))
}

// CreateImage
//
//	@Summary	获取图片验证码的 key
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	resp.Result{data=string}	"验证码图片的key"
//	@Router		/api/captcha/image [get]
func (c *CaptchaController) CreateImage(ctx *gin.Context) {
	ctx.JSON(resp.Succ(captcha.New()))
}

// ImageFile
//
//	@Summary	获取验证码图片
//	@Tags		验证码
//	@Accept		png
//	@Produce	json
//	@Param		key	path		string	true	"图片验证码 key"
//	@Success	200	{object}	string	"图片文件流"
//	@Router		/api/captcha/image/{key} [get]
func (c *CaptchaController) ImageFile(ctx *gin.Context) {
	key := ctx.Param("key")

	captcha.WriteImage(ctx.Writer, key, consts.CaptchaWidth, consts.CaptchaHeight)
}

// Clear
//
//	@Summary	清除过期验证码
//	@Tags		验证码
//	@Accept		json
//	@Produce	json
//	@Param		X-Auth-Token	header		string			true	"Authentication token"
//	@Success	200				{object}	resp.Result{}	"请求结果"
//	@Router		/api/captcha/clear [get]
func (c *CaptchaController) Clear(ctx *gin.Context) {
	c.service.ClearExpiration()
}

type CreateCaptchaBodyDto struct {
	Type         consts.CaptchaType   `json:"type" binding:"required"`                        // 验证码类型， 1-手机 2-邮箱
	Value        string               `json:"value" binding:"required" label:"手机/邮箱账号"`       // 手机/邮箱账号
	Scenes       consts.CaptchaScenes `json:"scenes" binding:"required"`                      // 使用场景， 1-注册 2-忘记密码 3-修改手机 4-修改邮箱
	CaptchaKey   string               `json:"captcha_key" binding:"required"`                 // 图形验证码的key
	CaptchaValue string               `json:"captcha_value" binding:"required" label:"图形验证码"` // 输入的图形验证码
}
