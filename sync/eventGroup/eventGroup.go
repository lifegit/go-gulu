/**
* @Author: TheLife
* @Date: 2020-3-19 2:01 上午
 */
package eventGroup

import (
	"sync"
	"time"
)

type WaitGroup interface {
	// Register waits returns a chan that waits on the given ID.
	// The chan will be triggered when Trigger is called with
	// the same ID.
	Register(timeout time.Duration, count int, data interface{}, callbackFunc func(data interface{})) (id uint64)
	// Trigger triggers the waiting chans with the given ID.
	Trigger(id uint64)
	IsRegistered(id uint64) bool
}

func NewEventGroup() WaitGroup {
	return &Event{m: make(map[uint64]event, 0)}
}

type event struct {
	data interface{}
	len  int
	c    chan int
}

type Event struct {
	l sync.RWMutex
	m map[uint64]event
}

//注册
func (w *Event) Register(timeout time.Duration, count int, data interface{}, callbackFunc func(data interface{})) (id uint64) {
	w.l.Lock()
	defer w.l.Unlock()

	ev := event{data: data, len: count, c: make(chan int, 1)}
	for true {
		id = uint64(time.Now().UnixNano())
		if _, ok := w.m[id]; !ok {
			break
		}
	}
	w.m[id] = ev
	go func(id uint64, timeout time.Duration, ev event) {
		for true {
			select {
			case <-ev.c:
				callbackFunc(ev.data)
				return
			case <-time.After(timeout):
				w.Trigger(id)
			}
		}
	}(id, timeout, ev)

	return id
}

//触发
func (w *Event) trigger(id uint64, tr bool) {
	w.l.Lock()
	ev := w.m[id]
	if ev.c != nil {
		ev.len--
		w.m[id] = ev
		if ev.len <= 0 || tr {
			ev.c <- ev.len
			close(ev.c)
			delete(w.m, id)
		}
	}

	w.l.Unlock()
}

//触发
func (w *Event) Trigger(id uint64) {
	w.trigger(id, false)
}

//判断该id是否被注册
func (w *Event) IsRegistered(id uint64) bool {
	w.l.RLock()
	defer w.l.Unlock()

	_, ok := w.m[id]
	return ok
}
