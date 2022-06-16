package paramValidator_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/validate/paramValidator"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Tb struct {
	Type uint `bindingCreate:"required"`
}

func TestName(t *testing.T) {

	err := paramValidator.ValidateCreate.Struct(Tb{})
	assert.Error(t, err)
}
