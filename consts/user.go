package consts

type UserStatus int8

const (
	UserStatusNormal UserStatus = 1  // 正常
	UserStatusLocked UserStatus = -1 // 锁定
)

var UserStatusMap = map[UserStatus]string{
	UserStatusNormal: "正常",
	UserStatusLocked: "封禁",
}

type AdminUserStatus int8

const (
	AdminUserStatusNormal AdminUserStatus = 1  // 正常
	AdminUserStatusLocked AdminUserStatus = -1 // 锁定
)

var AdminUserStatusMap = map[AdminUserStatus]string{
	AdminUserStatusNormal: "正常",
	AdminUserStatusLocked: "封禁",
}

var ContextUserInfoKey = "user" // 用于上下文中存储用户信息
