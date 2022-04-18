package fire_test

import (
	"github.com/lifegit/go-gulu/v2/pkg/fire"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestAllow(t *testing.T) {
	// url param
	param, _ := url.ParseQuery(
		`age=18&` +
			`name=Mr&` +
			`exclude=exclude&` + //  does not exist in the allow, so exclude
			`id=1&` +
			`id=999&` +
			`tag=student&` +
			`tag=老年&` +
			`tag=youth&` +
			`sort={"age":"descend"}&`,
	)

	_, _ = DBDryRun.
		OrderByColumn("age", fire.OrderAsc).
		Allow(param, fire.Allow{
			Where: fire.Filtered{"age"},
			Like:  fire.Filtered{"name"},
			Range: fire.Filtered{"id"},
			In:    fire.Filtered{"tag"},
			Sorts: fire.Filtered{"age"},
		}).
		CrudAllPage(User{}, &[]User{}, nil)

	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(2), `SELECT count(*) FROM "user" WHERE "user"."age" = '18' AND "user"."id" >= '1' AND "user"."id" <= '999' AND "user"."tag" IN ('student','老年','youth') AND "user"."name" LIKE '%Mr%'`)
	assert.Equal(t, DBDryRun.Logger.(*fire.Diary).LastSql(), `SELECT * FROM "user" WHERE "user"."age" = '18' AND "user"."id" >= '1' AND "user"."id" <= '999' AND "user"."tag" IN ('student','老年','youth') AND "user"."name" LIKE '%Mr%' ORDER BY "user"."age" DESC LIMIT 5`)
}
