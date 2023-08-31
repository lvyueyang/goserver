package console

import "github.com/gookit/color"

func Success(a ...any) {
	color.Green.Println(a...)
}

func Err(a ...any) {
	color.Error.Println(a...)

}

func Warn(a ...any) {
	color.Warn.Println(a...)
}
