package translate

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
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

func init() {
	// v := validator.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
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
		trans, ok = uni.GetTranslator("zh")
		if !ok {
			log.Fatalf("uni.GetTranslator(%s) failed", "zh")
		}
		err := zhTranslations.RegisterDefaultTranslations(v, trans)
		if err != nil {
			log.Fatal(err, "Register Default Translations With Error")
		}
	}
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
	}
	return
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
