package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"server/config"
	"server/consts"
	"server/consts/permission"
	"server/dal/dao"
	"server/dal/model"
	"server/utils"
	"server/utils/resp"
)

// AdminAuth 管理后台用户登录认证中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := isAdminLogin(c); err != nil {
			c.JSON(resp.AuthErr(err.Error()))
			c.Abort()
			return
		}
		c.Next()
	}
}

// AdminRole 管理后台用户权限中间件
func AdminRole(code string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := isAdminLogin(c)
		if err != nil {
			c.JSON(resp.AuthErr(err.Error()))
			c.Abort()
			return
		}

		if errP := isPermission(user, code); errP != nil {
			c.JSON(resp.ForbiddenErr(errP.Error()))
			c.Abort()
			return
		}

		c.Next()
	}
}

func isAdminLogin(c *gin.Context) (*model.AdminUser, error) {
	u := new(model.AdminUser)
	token := c.GetHeader("X-Auth-Token")
	if token == "" {
		return u, errors.New("未登录")
	}
	if info, err := utils.ParseAdminUserToken(token, config.Config.Auth.AdminTokenSecret); err != nil {
		return u, errors.New("身份过期")
	} else {
		userId := info.User.Id
		user, err := dao.AdminUser.Where(dao.AdminUser.ID.Eq(userId)).Preload(dao.AdminUser.Roles).Take()
		if err != nil {
			return user, errors.New("用户不存在")
		}
		if user.Status == consts.AdminUserStatusLocked {
			return user, errors.New("用户已封禁")
		}
		//fmt.Printf("USER %+v \n", user)
		c.Set(consts.ContextUserInfoKey, user)
		return user, nil
	}
}

func isPermission(user *model.AdminUser, code string) error {
	// 超管直接绕过权限认证
	if user.IsRoot == true {
		return nil
	}
	codeMap := make(map[string]bool)

	for _, role := range user.Roles {
		for _, code := range role.PermissionCodes {
			if codeMap[code] == false {
				codeMap[code] = true
			}
		}
	}
	if codeMap[code] == false {
		msg := "没有" + permission.AdminLabelMap[code].Label + "的权限"
		return errors.New(msg)
	}
	return nil
}
