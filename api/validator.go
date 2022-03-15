package api

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 创建翻译器
		zhTrans := zh.New() // 中文转换器
		enTrans := en.New() // 英文转换器

		uni := ut.New(zhTrans, zhTrans, enTrans)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "zh":
			// 内置tag注册 中文翻译器
			_ = zh_trans.RegisterDefaultTranslations(v, trans)
		case "en":
			_ = en_trans.RegisterDefaultTranslations(v, trans)
		default:
			_ = zh_trans.RegisterDefaultTranslations(v, trans)
		}
		_ = v.RegisterValidation("phone", vilidPhone)
		_ = v.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
			return ut.Add("phone", "{0}必须是一个有效的手机号码！", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("phone", fe.Field())
			return t
		})
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			count := 2
			name := strings.SplitN(field.Tag.Get("json"), ",", count)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		return
	}
	return
}

var vilidPhone validator.Func = func(fl validator.FieldLevel) bool {
	if phone, ok := fl.Field().Interface().(string); ok {
		result, err := regexp.MatchString(`^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`, phone)
		if err != nil {
			return false
		}
		return result
	}
	return true
}
