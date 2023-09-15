package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"server/consts"
	"server/dal/model"
	"server/lib/errs"
	"time"
)

type User struct {
	Id          uint               `json:"id"`
	AccountType consts.AccountType `json:"account_type"`
}

type UserClaims struct {
	User User
	jwt.RegisteredClaims
}

func CreateUserToken(user model.User, accountType consts.AccountType, secret string) (string, error) {
	now := time.Now()
	expireTime := now.Add(7 * 24 * time.Hour)
	userClaims := UserClaims{
		User{
			user.ID,
			accountType,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	token, err := tokenClaims.SignedString([]byte(secret))
	if err != nil {
		return "", errs.CreateServerError("Token 生成失败", err, user)
	}
	return token, nil
}

func ParseUserToken(token string, secret string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errs.CreateServerError("Token 错误", err, token)
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errs.CreateServerError("Token 解析失败", errors.New(tokenClaims.Raw), tokenClaims)
}

type AdminUser struct {
	Id uint `json:"id"`
}

type AdminUserClaims struct {
	User AdminUser
	jwt.RegisteredClaims
}

func CreateAdminUserToken(userID uint, secret string) (string, error) {
	now := time.Now()
	expireTime := now.Add(7 * 24 * time.Hour)
	userClaims := AdminUserClaims{
		AdminUser{
			userID,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	token, err := tokenClaims.SignedString([]byte(secret))
	if err != nil {
		return "", errs.CreateServerError("Token 生成失败", err, userID)
	}
	return token, nil
}

func ParseAdminUserToken(token string, secret string) (*AdminUserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &AdminUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errs.CreateServerError("Token 错误", err, token)
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*AdminUserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errs.CreateServerError("Token 解析失败", errors.New(tokenClaims.Raw), tokenClaims)
}
