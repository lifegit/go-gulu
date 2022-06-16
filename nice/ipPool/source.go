package ipPool

import (
	"reflect"
	"runtime"
	"sync"
)

type Source struct {
	index int
	fun   []Func
	mtx   sync.RWMutex
}

type Func func() (res []string)

// 添加来源
func (s *Source) AddSource(source ...Func) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.fun = append(s.fun, source...)
}

// 删除来源
func (s *Source) RemoveSource(source Func) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	sourceFun := getFunctionName(source)
	for t := len(s.fun) - 1; t >= 0; t-- { // 倒序
		if sourceFun == getFunctionName(s.fun[t]) {
			s.fun = append(s.fun[:t], s.fun[t+1:]...) // 删除
		}
	}
}

// for given function fn, get the name of function.
func getFunctionName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func (s *Source) GetFuncResult() (res []Func) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if s.index >= len(s.fun) {
		s.index = 0
	}
	res = append(s.fun[s.index:], s.fun[:s.index]...)
	s.index++

	return
}
