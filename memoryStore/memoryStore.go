/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package memoryStore

import (
	"container/list"
	"errors"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type idByTimeValue struct {
	timestamp time.Time
	id        string
}

// memoryStore is an internal store for captcha ids and their values.
type memoryStore struct {
	sync.RWMutex
	digitsById map[string]interface{}
	idByTime   *list.List
	// Number of items stored since last collection.
	numStored int
	// Number of saved items that triggers collection.
	collectNum int
	// Expiration time of captchas.
	expiration time.Duration
}

func New(maxCount, expireIn int) (Mem *memoryStore) {
	Mem = new(memoryStore)
	Mem.digitsById = make(map[string]interface{})
	Mem.idByTime = list.New()
	if maxCount <= 1024 {
		maxCount = 1024
	}
	Mem.collectNum = maxCount
	if expireIn <= 0 {
		expireIn = 30
	}
	Mem.expiration = time.Minute * time.Duration(expireIn)
	return Mem
}

func (s *memoryStore) Set(id string, value interface{}) {
	s.Lock()
	s.digitsById[id] = value
	s.idByTime.PushBack(idByTimeValue{time.Now(), id})
	s.numStored++
	s.Unlock()
	if s.numStored > s.collectNum {
		go s.collect()
	}
}

func (s *memoryStore) Get(id string, clear bool) (value interface{}, err error) {
	if !clear {
		// When we don't need to clear captcha, acquire read lock.
		s.RLock()
		defer s.RUnlock()
	} else {
		s.Lock()
		defer s.Unlock()
	}
	value, ok := s.digitsById[id]
	if !ok {
		return nil, errors.New("value not found")
	}
	if clear {
		delete(s.digitsById, id)
	}
	return
}

func (s *memoryStore) collect() {
	logrus.Warn("memory store collect function has been called some value will be lost")
	now := time.Now()
	s.Lock()
	defer s.Unlock()
	s.numStored = 0
	for e := s.idByTime.Front(); e != nil; {
		ev, ok := e.Value.(idByTimeValue)
		if !ok {
			return
		}
		if ev.timestamp.Add(s.expiration).Before(now) {
			delete(s.digitsById, ev.id)
			next := e.Next()
			s.idByTime.Remove(e)
			e = next
		} else {
			return
		}
	}
}

func (s *memoryStore) GetUint(id string) (value uint, err error) {
	vv, err := s.Get(id, false)
	if err != nil {
		return 0, err
	}
	value, ok := vv.(uint)
	if ok {
		return value, nil
	}
	return 0, errors.New("mem:has value of this id, but is not type of uint")
}
