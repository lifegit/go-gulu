package arrayconv_test

import (
	"github.com/lifegit/go-gulu/v2/conv/arrayconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringIn(t *testing.T) {
	res := arrayconv.StringIn("B", []string{"A", "B", "C"})
	assert.Equal(t, res, true)
}

func TestIntIn(t *testing.T) {
	res := arrayconv.IntIn(2, []int{1, 2, 3})
	assert.Equal(t, res, true)
}

func TestUIntIn(t *testing.T) {
	res := arrayconv.UIntIn(2, []uint{1, 2, 3})
	assert.Equal(t, res, true)
}
