package consts

// Sex 性别
type Sex int

const (
	SexUnknown Sex = 0 // 未知
	SexMan     Sex = 1 // 男
	SexWoman   Sex = 2 // 女
)

func (g Sex) Label() string {
	switch g {
	case SexMan:
		return "男"
	case SexWoman:
		return "女"
	case SexUnknown:
		return "未知"
	}
	return "未知"
}
