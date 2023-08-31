package validate

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"selfserver/lib/logs"
)

func InitValidate() {
	err := transInit("zh")
	if err != nil {
		logs.Err().Err(err).Msg("验证器国际化初始化失败")
	}
}

var (
	uni   *ut.UniversalTranslator
	v     *validator.Validate
	trans ut.Translator
)

func transInit(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()

		uni = ut.New(zhT, enT, zhT)

		var o bool
		trans, o = uni.GetTranslator(local)
		if !o {
			return errors.New("translator init failed")
		}

		switch local {
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		}

		return err
	}
	return
}

func ErrTransform(err error) string {
	errs, ok := err.(validator.ValidationErrors)
	fmt.Println(ok)
	if !ok {
		// 非validator.ValidationErrors 类型错误直接返回
		return err.Error()
	}
	// validator.ValidationErrors 类型错误则进行翻译
	msg := ""
	for _, value := range errs.Translate(trans) {
		msg = msg + value + ";"
	}
	return msg

}
