package utils

type EnumLabelInterface interface {
	Label() string
}

// EnumLabel 返回枚举类型的人类可读名称
func EnumLabel(e EnumLabelInterface) string {
	return e.Label()
}
