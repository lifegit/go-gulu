/**
* @Author: TheLife
* @Date: 2020-10-7 11:36 下午
 */
package arrayconv

import "sync"


type ListUint struct {
	mu           sync.RWMutex
	list []uint
}

func (l *ListUint) PushList(o []uint)  {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	l.list = append(l.list, o...)
}


func (l *ListUint) Push(o uint)  {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.list = append(l.list, o)
}

func (l *ListUint) Del(o uint) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	for key, v := range l.list {
		if v == o {
			l.list = append(l.list[:key], l.list[key+1:]...)
			return true
		}
	}
	return false
}
func (l *ListUint) Exist(o uint) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, item := range l.list {
		if item == o {
			return true
		}
	}

	return false
}

func (l *ListUint) ToList() (list []uint) {
	l.mu.Lock()
	defer l.mu.Unlock()

	list = make([]uint, len(l.list))
	copy(list, l.list)

	return
}