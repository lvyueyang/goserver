package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/consts"
	"server/lib/valid"
	"server/modules/service"
	"server/utils/resp"
)

type AdminAuthController struct {
	service          *service.AdminAuthService
	adminUserService *service.AdminUserService
	captchaService   *service.CaptchaService
}

func NewAdminAuthController(e *gin.Engine) {
	c := &AdminAuthController{
		service:          service.NewAdminAuthService(),
		adminUserService: service.NewAdminUserService(),
		captchaService:   service.NewCaptchaService(),
	}
	router := e.Group("/api/admin/auth")
	router.POST("/login", c.Login)
	router.POST("/init-root-user", c.InitRootUser)
	router.POST("/forget-password", c.ForgetPassword)
}

// Login
//
//	@Summary	用户登录
//	@Tags		管理后台-用户认证
//	@Accept		json
//	@Produce	json
//	@Param		req	body		adminUserLoginBodyDto							true	"body"
//	@Success	200	{object}	resp.Result{data=adminUserLoginSuccessResponse}	"resp"
//	@Router		/api/admin/auth/login [post]
func (c *AdminAuthController) Login(ctx *gin.Context) {
	var body = new(adminUserLoginBodyDto)
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	token, err := c.service.UsernameAndPasswordLogin(body.Username, body.Password)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(adminUserLoginSuccessResponse{Token: token}, "登录成功"))
	}
}

// InitRootUser
//
//	@Summary	初始化超级管理员用户
//	@Tags		管理后台-用户认证
//	@Accept		json
//	@Produce	json
//	@Param		req	body		adminInitRootUserBodyDto						true	"body"
//	@Success	200	{object}	resp.Result{data=adminUserLoginSuccessResponse}	"resp"
//	@Router		/api/admin/auth/init-root-user [post]
func (c *AdminAuthController) InitRootUser(ctx *gin.Context) {
	var body = new(adminInitRootUserBodyDto)
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	token, err := c.adminUserService.CreateRootUser(body.Username, body.Name, body.Password, body.Email)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(adminUserLoginSuccessResponse{Token: token}, "超级管理员创建成功"))
	}
}

// ForgetPassword
//
//	@Summary	忘记密码
//	@Tags		管理后台-用户认证
//	@Accept		json
//	@Produce	json
//	@Param		req	body		adminUserForgetPasswordBodyDto	true	"body"
//	@Success	200	{object}	resp.Result{}					"resp"
//	@Router		/api/admin/auth/forget-password [post]
func (c *AdminAuthController) ForgetPassword(ctx *gin.Context) {
	var body = new(adminUserForgetPasswordBodyDto)
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	if ok, err := c.captchaService.Validate(body.Email, consts.CaptchaTypeEmail, body.Captcha, consts.CaptchaScenesForgetPassword); ok == false {
		ctx.JSON(resp.ParamErr(err.Error()))
		return
	}

	if err := c.adminUserService.ResetPassword(body.Email, body.Password); err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(nil, "密码重置成功"))
	}
}

type adminUserLoginBodyDto struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type adminUserForgetPasswordBodyDto struct {
	Email    string `json:"email" binding:"required" label:"邮箱"`    // 邮箱
	Password string `json:"password" binding:"required" label:"密码"` // 密码
	Captcha  string `json:"captcha" binding:"required" label:"验证码"` // 邮箱验证码
}

type adminUserLoginSuccessResponse struct {
	Token string `json:"token"`
}

type adminInitRootUserBodyDto struct {
	Name     string `json:"name" binding:"required" label:"用户名"`     // 昵称
	Username string `json:"username" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required" label:"密码"`  // 密码
	Email    string `json:"email" binding:"required" label:"邮箱"`     // 邮箱
}
