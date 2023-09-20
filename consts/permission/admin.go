package permission

type LabelType struct {
	Label string `json:"label"`
}

const (
	AdminUserFind           string = "admin:user:find:list"
	AdminUserCreate         string = "admin:user:create"
	AdminUserDelete         string = "admin:user:delete"
	AdminUserUpdateInfo     string = "admin:user:update:info"
	AdminUserUpdateStatus   string = "admin:user:update:status"
	AdminUserUpdatePassword string = "admin:user:update:password"
	AdminUserUpdateRole     string = "admin:user:update:role"

	AdminRoleFind       string = "admin:role:find:list"
	AdminRoleCreate     string = "admin:role:create"
	AdminRoleUpdateInfo string = "admin:role:update:info"
	AdminRoleUpdateCode string = "admin:role:update:code"
	AdminRoleDelete     string = "admin:role:delete"
)

var AdminLabelMap = map[string]LabelType{
	AdminUserFind: {
		Label: "查询管理员列表",
	},
	AdminUserCreate: {
		Label: "创建管理员",
	},
	AdminUserDelete: {
		Label: "删除管理员",
	},
	AdminUserUpdateInfo: {
		Label: "修改管理员基本信息",
	},
	AdminUserUpdateStatus: {
		Label: "修改管理员状态",
	},
	AdminUserUpdatePassword: {
		Label: "修改管理员密码",
	},
	AdminUserUpdateRole: {
		Label: "修改管理员角色",
	},

	AdminRoleFind: {
		Label: "查询管理角色列表",
	},
	AdminRoleCreate: {
		Label: "创建管理角色",
	},
	AdminRoleUpdateInfo: {
		Label: "修改管理角色信息",
	},
	AdminRoleDelete: {
		Label: "删除管理角色",
	},
	AdminRoleUpdateCode: {
		Label: "修改管理角色权限码列表",
	},
}
