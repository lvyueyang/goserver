package permission

type LabelType struct {
	Label string `json:"label"`
}

const (
	AdminUserFind           string = "admin:user:find:list"
	AdminUserCreate         string = "admin:user:create"
	AdminUserUpdateInfo     string = "admin:user:update:info"
	AdminUserUpdateStatus   string = "admin:user:update:status"
	AdminUserUpdatePassword string = "admin:user:update:password"
	AdminUserUpdateRole     string = "admin:user:update:role"

	AdminRoleFind       string = "admin:role:find:list"
	AdminRoleCreate     string = "admin:role:create"
	AdminRoleUpdateInfo string = "admin:role:update:info"
	AdminRoleFindCode   string = "admin:role:list:code"
	AdminRoleUpdateCode string = "admin:role:update:code"
	AdminRoleDelete     string = "admin:role:delete"
)

var AdminLabelMap = map[string]LabelType{
	AdminUserFind: {
		Label: "查询管理员账号列表",
	},
	AdminUserCreate: {
		Label: "创建管理员账号",
	},
	AdminUserUpdateInfo: {
		Label: "修改管理员账号",
	},
	AdminUserUpdateStatus: {
		Label: "修改管理员账号状态",
	},
	AdminUserUpdatePassword: {
		Label: "修改管理员账号密码",
	},
	AdminUserUpdateRole: {
		Label: "修改管理员账号角色",
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
	AdminRoleFindCode: {
		Label: "查询管理角色权限码列表",
	},
	AdminRoleUpdateCode: {
		Label: "修改管理角色权限码列表",
	},
}
