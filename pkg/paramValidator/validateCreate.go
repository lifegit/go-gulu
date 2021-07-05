/**
* @Author: TheLife
* @Date: 2021/6/17 下午7:00
 */
package paramValidator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"sync"
)

var ValidateCreate *validator.Validate

func init() {
	v := &defaultValidator{}
	ValidateCreate = v.Engine().(*validator.Validate)
}

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("bindingCreate")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
