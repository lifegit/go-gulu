/**
* @Author: TheLife
* @Date: 2021/6/17 下午7:03
 */
package paramValidator_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/paramValidator"
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
