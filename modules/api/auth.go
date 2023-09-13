package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/lib/valid"
	"server/modules/service"
	"server/utils/resp"
)

type AuthController struct {
	service *service.AuthService
}

var authService = service.NewAuthService()

func NewAuthController(e *gin.Engine) {
	c := &AuthController{
		service: service.NewAuthService(),
	}
	router := e.Group("/api/auth")
	router.POST("/login", c.Login)
	router.POST("/register", c.Register)
}

//
//func (c *AuthController) New(e *gin.Engine) {
//	router := e.Group("/api/auth")
//
//	router.POST("/login", c.Login)
//}

// Login
//
//	@Summary	用户登录
//	@Tags		用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		LoginBodyDto							true	"body"
//	@Success	200	{object}	resp.Result{data=LoginSuccessResponse}	"resp"
//	@Router		/api/auth/login [post]
func (c AuthController) Login(ctx *gin.Context) {
	var body LoginBodyDto
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	token, err := authService.UsernameAndPasswordLogin(body.Username, body.Password)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(LoginSuccessResponse{Token: token}, "登录成功"))
	}
}

// Register
//
//	@Summary	用户注册
//	@Tags		用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		RegisterBodyDto							true	"body"
//	@Success	200	{object}	resp.Result{data=LoginSuccessResponse}	"resp"
//	@Router		/api/auth/register [post]
func (c AuthController) Register(ctx *gin.Context) {
	var body RegisterBodyDto
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	token, err := c.service.UsernameAndPasswordRegister(body.Email, body.Username, body.Password)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(LoginSuccessResponse{Token: token}, "登录成功"))
	}
}

type LoginBodyDto struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type RegisterBodyDto struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Email    string `json:"email"`    // 邮箱
	Captcha  string `json:"captcha"`  // 邮箱验证码
}

type LoginSuccessResponse struct {
	Token string `json:"token"`
}
