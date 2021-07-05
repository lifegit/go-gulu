/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package stringconv_test

import (
	"github.com/lifegit/go-gulu/v2/conv/stringconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch(t *testing.T) {
	res := stringconv.Match("1234516", "12", "1")
	assert.Equal(t, res, "345")
}
