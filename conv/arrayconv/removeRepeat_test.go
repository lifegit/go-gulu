/**
* @Author: TheLife
* @Date: 2021/6/17 下午1:10
 */
package arrayconv_test

import (
	"github.com/lifegit/go-gulu/v2/conv/arrayconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveRepeat(t *testing.T) {
	res := arrayconv.RemoveRepeat([]int{1, 2, 2, 3, 4, 3, 5, 3, 1})
	assert.Equal(t, res, []int{1, 2, 3, 4, 5})
}
