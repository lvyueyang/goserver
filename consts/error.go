package consts

const (
	Success         = 200
	ErrorParam      = 400
	ErrorAuth       = 401
	ErrorPermission = 403
	ErrorServer     = 500
)

var MsgFlags = map[int]string{
	Success:         "请求成功",
	ErrorServer:     "服务端错误",
	ErrorParam:      "请求参数错误",
	ErrorAuth:       "认证失败",
	ErrorPermission: "权限不足",
}

func GetCodeMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ErrorServer]
}
