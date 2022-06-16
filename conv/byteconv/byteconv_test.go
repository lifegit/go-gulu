package byteconv_test

import (
	"github.com/lifegit/go-gulu/v2/conv/byteconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch(t *testing.T) {
	res := byteconv.Match([]byte{1, 2, 3, 4, 3, 4, 34}, []byte{3, 4}, []byte{4})
	assert.Equal(t, res, []byte{3})
}

func TestMatchLast(t *testing.T) {
	res := byteconv.MatchLast([]byte{1, 2, 3, 4, 3, 4, 34, 4}, []byte{3, 4}, []byte{4})
	assert.Equal(t, res, []byte{3, 4, 34})
}
