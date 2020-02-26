/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package paramValidator

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"sync"
)

var Validate *validator.Validate
var ValidateCreate *validator.Validate

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func Setup() {
	binding.Validator = new(defaultValidator)
	binding.Validator.Engine()
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		// Validate
		Validate = validator.New()
		Validate.SetTagName("binding")
		// ValidateCreate
		ValidateCreate = validator.New()
		ValidateCreate.SetTagName("bindingCreate")
		// v.validate
		v.validate = Validate

		// add any custom validations etc. here
		addValidation("order", validationOrder)
	})
}

func addValidation(tag string, fn func(fl validator.FieldLevel) bool) {
	_ = Validate.RegisterValidation(tag, fn)
	_ = ValidateCreate.RegisterValidation(tag, fn)
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

//--------validation----------
func validationOrder(fl validator.FieldLevel) bool {
	//fmt.Println("validationTest")

	//fmt.Println("FieldName:", fl.FieldName())
	//fmt.Println("StructFieldName", fl.StructFieldName())
	//fmt.Println("Parm:", fl.Param())
	//fmt.Println("Field:", fl.Field()) //1
	//fmt.Println("Parent:",fl.Parent())
	//fmt.Println("top:",fl.Top())

	//fmt.Println("Parm:", fl.ExtractType())
	return false
}
