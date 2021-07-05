/**
* @Author: TheLife
* @Date: 2021/6/17 下午4:12
 */
package structureconv_test

import (
	"github.com/lifegit/go-gulu/conv/structureconv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructToMap(t *testing.T) {
	type People struct {
		Name string
		Age  int
	}
	type Class struct {
		Student []People
		Teacher People
	}

	c := Class{
		Teacher: People{
			Name: "老师",
			Age:  36,
		},
		Student: []People{{
			Name: "学生A",
			Age:  18,
		}, {
			Name: "学生B",
			Age:  18,
		}},
	}

	// map[Student:[{学生A 18} {学生B 18}] Teacher:{老师 36}]
	res := structureconv.StructToMap(c)

	m := make(map[string]interface{})
	m["Student"] = c.Student
	m["Teacher"] = c.Teacher

	assert.Equal(t, res, m)
}

func TestIsBlank(t *testing.T) {
	res1 := structureconv.IsBlank("")
	assert.Equal(t, res1, true)

	res2 := structureconv.IsBlank(0)
	assert.Equal(t, res2, true)

	res3 := structureconv.IsBlank(nil)
	assert.Equal(t, res3, true)
}
