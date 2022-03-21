package fire_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog(t *testing.T) {
	DBDryRun.WhereCompare("age", 18, fire.CompareGreaterEqual).Model(User{}).Take(&User{})
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), "SELECT * FROM `user` WHERE `age` >= 18 LIMIT 1")
}
