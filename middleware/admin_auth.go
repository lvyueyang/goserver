package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"server/config"
	"server/consts"
	"server/dal/dao"
	"server/utils"
	"server/utils/resp"
)

// AdminAuth 管理后台用户登录认证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := isAdminLogin(c); err != nil {
			c.JSON(resp.AuthErr(err.Error()))
			return
		}
		c.Next()
	}
}

func isAdminLogin(c *gin.Context) error {
	token := c.GetHeader("X-Auth-Token")
	if token == "" {
		return errors.New("未登录")
	}
	if info, err := utils.ParseAdminUserToken(token, config.Config.Auth.AdminTokenSecret); err != nil {
		return errors.New("身份过期")
	} else {
		userId := info.User.Id
		user, err := dao.AdminUser.Where(dao.AdminUser.ID.Eq(userId)).Preload(dao.AdminUser.Roles).Take()
		if err != nil {
			return errors.New("用户不存在")
		}
		if user.Status == consts.AdminUserStatusLocked {
			return errors.New("用户已封禁")
		}
		//fmt.Printf("USER %+v \n", user)
		c.Set(consts.ContextUserInfoKey, user)
	}
	return nil
}
