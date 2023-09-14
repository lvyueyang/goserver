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

type Claims struct {
	User User
	jwt.RegisteredClaims
}

var HmacSecret = []byte("aaaaaaaaaaaa")

func CreateUserToken(user model.User, accountType consts.AccountType) (string, error) {
	now := time.Now()
	expireTime := now.Add(7 * 24 * time.Hour)
	userClaims := Claims{
		User{
			user.ID,
			accountType,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	token, err := tokenClaims.SignedString(HmacSecret)
	if err != nil {
		return "", errs.CreateServerError("Token 生成失败", err, user)
	}
	return token, nil
}

func ParseUserToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return HmacSecret, nil
	})
	if err != nil {
		return nil, errs.CreateServerError("Token 错误", err, token)
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, errs.CreateServerError("Token 解析失败", errors.New(tokenClaims.Raw), tokenClaims)
}
