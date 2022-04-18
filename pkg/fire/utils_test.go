package fire_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
	"testing"
)

func TestCamel2Case(t *testing.T) {
	assert.Equal(t, fire.Camel2Case("name"), "name")
	assert.Equal(t, fire.Camel2Case("realName"), "real_name")
	assert.Equal(t, fire.Camel2Case("ReadName"), "read_name")
}

func TestIf(t *testing.T) {
	f := func(v ...int) int {
		return fire.If(v != nil, v, []int{0}).([]int)[0]
	}

	assert.Equal(t, f(), 0)
	assert.Equal(t, f(1), 1)
}

func TestParseDataType(t *testing.T) {
	assert.Equal(t, fire.ParseDataType("2"), int64(2))
	assert.Equal(t, fire.ParseDataType("2.6"), float64(2.6))
	assert.Equal(t, fire.ParseDataType("2 hello"), "2 hello")
}

func TestParseColumn(t *testing.T) {
	c0 := clause.Column{
		Table: clause.CurrentTable,
		Name:  "col",
	}
	v0 := fire.ParseColumn("col")
	assert.EqualValues(t, v0, c0)

	c1 := clause.Column{
		Name: "col",
	}
	v1 := fire.ParseColumn(c1)
	assert.EqualValues(t, v1, c1)
}
