package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"server/config"
	"server/consts"
	"server/dal/model"
	"server/lib/errs"
	"time"
)

type AuthService struct {
	accountService *AccountService
	userService    *UserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		accountService: NewAccountService(),
		userService:    NewUserService(),
	}
}

type LoginOptions struct {
	Username string
}

type UsernameAndPasswordRegisterOptions struct {
	Username string
	Email    string
	Code     string
	Password string
}

// UsernameAndPasswordRegister 使用用户名邮箱和密码注册
func (s *AuthService) UsernameAndPasswordRegister(email, username, password string) (string, error) {
	var opt = map[string]string{
		"email":    email,
		"username": username,
		"password": password,
	}
	// 验证用户名和邮箱是否已被使用
	if info, _ := s.accountService.UseEmailFindOne(email); info.ID != 0 {
		return "", &errs.ClientError{Msg: "邮箱已存在", Info: nil}
	}
	if info, _ := s.accountService.UseUsernameFindOne(username); info.ID != 0 {
		return "", &errs.ClientError{Msg: "用户名已存在", Info: nil}
	}

	var hashPassword []byte
	// 密码加盐
	if result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", &errs.ServerError{Msg: "密码加盐失败", Err: err, Info: opt}
	} else {
		hashPassword = result
	}

	var userInfo model.User
	// 创建账号
	if result, err := s.accountService.CreateEmail(username, email, string(hashPassword)); err != nil {
		return "", &errs.ServerError{Msg: "账号创建失败", Err: err, Info: opt}
	} else {
		userInfo = result
	}

	token, err := CreateToken(userInfo, consts.EmailAccountType)

	if err != nil {
		errInfo := map[string]any{"userInfo": userInfo, "opt": opt}
		return "", errs.CreateServerError("Token 生成失败", err, errInfo)
	}
	return token, nil
}

// UsernameAndPasswordLogin 使用用户名和密码登录
func (s *AuthService) UsernameAndPasswordLogin(username string, password string) (string, error) {
	info, _ := s.accountService.UseUsernameFindOne(username)
	if info.ID == 0 {
		return "", &errs.ClientError{Msg: "用户名未注册", Info: nil}
	}
	err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password))
	if err != nil {
		return "", errs.CreateClientError("密码错误", nil)
	}
	userinfo := s.userService.FindByID(info.UserID)
	token, err := CreateToken(userinfo, info.Type)
	if err != nil {
		return "", errs.CreateServerError("Token 生成失败", err, nil)
	}
	return token, nil
}

func (s *AuthService) Register() {
}

type Claims struct {
	UserID      uint               `json:"user_id"`
	AccountType consts.AccountType `json:"account_type"`
	jwt.RegisteredClaims
}

var HmacSecret = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU")

func CreateToken(user model.User, accountType consts.AccountType) (string, error) {
	now := time.Now()
	expireTime := now.Add(7 * 24 * time.Hour)
	claims := Claims{
		user.ID,
		accountType,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    config.Config.Auth.TokenSecret,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SigningString()
	if err != nil {
		return "", &errs.ServerError{Msg: "Token 生成失败", Err: err, Info: user}
	}
	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return HmacSecret, nil
	})
	if err != nil {
		return nil, &errs.ServerError{Msg: "Token 解析失败", Err: err, Info: token}
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, &errs.ServerError{Msg: "Token 解析失败", Err: errors.New(tokenClaims.Raw), Info: tokenClaims}
}
