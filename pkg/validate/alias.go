package validate

import (
	"github.com/go-playground/validator/v10"
)

func NewAlias(validate *validator.Validate, date []Validate) {
	for _, v := range date {
		validate.RegisterAlias(v.Tag, v.Alias)
	}
}
