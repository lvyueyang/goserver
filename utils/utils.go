package utils

import (
	"github.com/gin-gonic/gin"
	"server/consts"
	"server/dal/model"
)

// GetCurrentAdminUser 获取当前登录的管理员用户
func GetCurrentAdminUser(ctx *gin.Context) *model.AdminUser {
	return ctx.MustGet(consts.ContextUserInfoKey).(*model.AdminUser)
}
