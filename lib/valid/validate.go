package valid

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
	"golang.org/x/exp/slog"
	"reflect"
	"strings"
)

func New() {
	err := transInit("zh")
	if err != nil {
		slog.Error("验证器国际化初始化失败")
	}
}

var (
	uni   *ut.UniversalTranslator
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

		// 自定义标签
		v.RegisterTagNameFunc(customTagNameFunc)

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
	builder := strings.Builder{}
	ind := 0
	transMap := errs.Translate(trans)
	transMapLen := len(transMap)
	for _, value := range transMap {
		ind++
		builder.WriteString(value)
		if ind < transMapLen {
			builder.WriteString(", ")
		}
	}
	return builder.String()

}

// customTagNameFunc 自定义标签名称
func customTagNameFunc(field reflect.StructField) string {
	// 可以根据 label 展示
	label := field.Tag.Get("label")
	if len(label) == 0 {
		return field.Name
	}
	return label
}
