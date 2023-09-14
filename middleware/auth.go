package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"server/dal/dao"
	"server/utils"
	"server/utils/resp"
)

// Auth 用户登录认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := isLogin(c); err != nil {
			c.JSON(resp.AuthErr(err.Error()))
			return
		}
		c.Next()
	}
}

func isLogin(c *gin.Context) error {
	token := c.GetHeader("X-Auth-Token")
	if token == "" {
		return errors.New("未登录")
	}
	if info, err := utils.ParseUserToken(token); err != nil {
		return errors.New("身份过期")
	} else {
		userId := info.User.Id
		user, err := dao.User.Where(dao.User.ID.Eq(userId)).First()
		if err != nil {
			return errors.New("用户不存在")
		}
		//fmt.Printf("USER %+v \n", user)
		c.Set("user", user)
	}
	return nil
}
