package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

func init() {
	// v := validator.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		NewAlias(v, Validates)

		t, err := NewTranslate(v, Validates...)
		if err != nil {
			log.Fatal(err)
		}
		trans = t
	}
}

type Validate struct {
	Tag         string
	Translation string
	Alias       string
}

var Validates = []Validate{
	{
		Tag:         "phone",
		Translation: "{0}为11位手机号",
		Alias:       "numeric,len=11,startswith=1",
	},
	{
		Tag:         "password_simple",
		Translation: "{0}至少满足6位",
		Alias:       "min=6",
	},
}
