package fire_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllow(t *testing.T) {
	// url param
	param := fire.Param{
		Params: fire.Params{
			"age":     18,
			"name":    "Mr",
			"exclude": "exclude", //  does not exist in the allow, so exclude
			"id":      []int64{1, 999},
			"tag":     []string{"学生", "儿子", "青年"},
		},
		Sort: fire.Sort{
			"id":  "ascend",
			"age": "descend",
		},
	}

	_, _ = DBDryRun.OrderByColumn("age", fire.OrderAsc).
		Allow(param, fire.Allow{
			Where: []string{"age"},
			Like:  []string{"name"},
			Range: []string{"id"},
			In:    []string{"tag"},
			Sorts: []string{"age"},
		}).
		CrudAllPage(User{}, &[]User{})

	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(2), "SELECT count(*) FROM `user` WHERE `age` = 18 AND (`id` >= 1 AND `id` <= 999) AND `tag`  IN ('学生','儿子','青年') AND `name` LIKE '%Mr%'")
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), "SELECT * FROM `user` WHERE `age` = 18 AND (`id` >= 1 AND `id` <= 999) AND `tag`  IN ('学生','儿子','青年') AND `name` LIKE '%Mr%' ORDER BY `age` desc LIMIT 20")
}
