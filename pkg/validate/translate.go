package validate

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
	"strings"
)

var trans ut.Translator

func NewTranslate(validate *validator.Validate, date ...Validate) (trans ut.Translator, err error) {
	// 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		json := fld.Tag.Get("json")
		if json == "" {
			json = fld.Name
		}

		label := fld.Tag.Get("label")
		if label == "" {
			label = json
		}
		return strings.Join([]string{json, label}, "@")
	})
	// 第一个参数是备用（fallback）的语言环境,后面的参数是应该支持的语言环境（支持多个）
	uni := ut.New(en.New(), zh.New())
	trans, ok := uni.GetTranslator("zh")
	if !ok {
		return nil, fmt.Errorf("uni.GetTranslator(%s) failed", "zh")
	}
	err = zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, err
	}

	for _, v := range date {
		if err = validate.RegisterTranslation(v.Tag, trans, registrationFunc(v.Tag, v.Translation, false), translateFunc); err != nil {
			return
		}
	}

	return trans, err
}

func Translate(validationErr error) (errs validator.ValidationErrorsTranslations, first error) {
	if v, b := validationErr.(validator.ValidationErrors); b && len(v) > 0 {
		errs = make(validator.ValidationErrorsTranslations)
		for i, fe := range v {
			// get struct info
			name := strings.SplitN(fe.Namespace(), ".", 2)[1]
			json := getStructJson(name)
			// translate
			errs[json] = strings.SplitN(fe.Translate(trans), "@", 2)[1]

			if i == 0 {
				first = errors.New(errs[json])
			}
		}
		return
	}
	return nil, validationErr
}

func getStructJson(name string) (res string) {
	var buf []string
	for _, v := range strings.Split(name, ".") {
		vs := strings.Split(v, "@")
		if len(vs) > 0 {
			buf = append(buf, vs[0])
		}

	}

	return strings.Join(buf, ".")
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		return ut.Add(tag, translation, override)
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("警告: 翻译字段错误: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
