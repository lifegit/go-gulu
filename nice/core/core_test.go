/**
* @Author: TheLife
* @Date: 2020-2-26 6:30 上午
 */
package core_test

import (
	"fmt"
	"github.com/lifegit/go-gulu/v2/nice/core"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	// Do jobs without params
	//Every(1).Second().Do(task)
	//Every(2).Seconds().Do(task)
	//Every(10).Second().Do(task)
	//Every(2).Minutes().Do(task)
	//Every(1).Hour().Do(task)
	//Every(2).Hours().Do(task)
	//Every(1).Day().Do(task)
	//Every(2).Days().Do(task)

	// Do jobs on specific weekday
	//Every(1).Monday().Do(task)
	//Every(1).Thursday().Do(task)

	// function At() take a string like 'hour:min'
	//Every(1).Day().At("10:30").Do(task)
	//Every(1).Monday().At("18:30").Do(task)

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	tasks := core.NewScheduler()
	tasks.Every(1).Seconds().Do(func() {
		fmt.Println(time.Now().Unix())
	})
	tasks.Start()

	time.Sleep(time.Second * 2)
}
