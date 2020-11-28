/**
* @Author: TheLife
* @Date: 2020/11/25 上午1:42

 一个可以写 log 的 RWMutex 组件。
 通常使用log用于排查死锁问题。
 可直接替换 sync 包中的 RWMutex
 */
package rwMutexLog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-gulu/logging"
	"runtime"
	"strings"
	"sync"
)

const stackDeep = 2

type RWMutex struct {
	m sync.RWMutex

	once sync.Once

	Logger *logrus.Logger
	LoggerKey string
}

func (r *RWMutex) RUnlock ()  {
	r.m.RUnlock()
	r.log()
}

func (r *RWMutex) RLock ()  {
	r.m.RLock()
	r.log()
}

func (r *RWMutex) Unlock ()  {
	r.m.Unlock()
	r.log()
}

func (r *RWMutex) Lock ()  {
	r.m.Lock()
	r.log()
}



func (r *RWMutex) log ()  {
	// memory stack
	var buf [2 << 10]byte
	z := string(buf[:runtime.Stack(buf[:], true)])

	// singleton
	r.once.Do(func() {
		list := strings.Split(z, "\n")
		index := 1 + 2 * stackDeep
		if len(list) < index {
			return
		}
		if key := match(list[index]); key != "" {
			fmt.Println(key)
			r.LoggerKey = key
			r.Logger = logging.NewLogger(fmt.Sprintf("./runtime/logs/log-%s", key), 3, &logrus.TextFormatter{}, nil)
		}
	})

	// write log
	if r.Logger != nil {
		r.Logger.Println(z)
	}
}

func match(text string) (res string) {
	startIndex := strings.LastIndex(text, "/")
	if startIndex == -1 {
		return
	}

	endIndex := strings.Index(text[startIndex:],".")
	if endIndex == - 1 {
		return
	}

	return text[startIndex+1 : startIndex + endIndex]
}


