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
//	@Param		req	body		loginBodyDto							true	"body"
//	@Success	200	{object}	resp.Result{data=loginSuccessResponse}	"resp"
//	@Router		/api/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var body = new(loginBodyDto)
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}
	token, err := c.service.UsernameAndPasswordLogin(body.Username, body.Password)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(loginSuccessResponse{Token: token}, "登录成功"))
	}
}

// Register
//
//	@Summary	用户注册
//	@Tags		用户
//	@Accept		json
//	@Produce	json
//	@Param		req	body		registerBodyDto							true	"body"
//	@Success	200	{object}	resp.Result{data=loginSuccessResponse}	"resp"
//	@Router		/api/auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var body = new(registerBodyDto)
	if err := ctx.ShouldBindBodyWith(body, binding.JSON); err != nil {
		ctx.JSON(resp.ParamErr(valid.ErrTransform(err)))
		return
	}

	token, err := c.service.UsernameAndPasswordRegister(body.Email, body.Username, body.Password)
	if err != nil {
		ctx.JSON(resp.ParseErr(err))
	} else {
		ctx.JSON(resp.Success(loginSuccessResponse{Token: token}, "注册成功"))
	}
}

type loginBodyDto struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type registerBodyDto struct {
	Username string `json:"username" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required" label:"密码"`  // 密码
	Email    string `json:"email" binding:"required" label:"邮箱"`     // 邮箱
	Captcha  string `json:"captcha"`                                 // 邮箱验证码
}

type loginSuccessResponse struct {
	Token string `json:"token"`
}
