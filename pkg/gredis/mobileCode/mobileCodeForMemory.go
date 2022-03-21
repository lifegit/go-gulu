/**
* @Author: TheLife
* @Date: 2020-2-25 9:00 下午
 */
package mobileCode

import (
	"github.com/lifegit/go-gulu/v2/nice/memoryStore"
)

// sms := mobileCode.New(mobileCode.ForMemory{ Store: memoryStore.New(1024, time.Minute * 5) } )

type ForMemory struct {
	Store *memoryStore.MemoryStore
}

func (f ForMemory) Set(key, value string) error {
	f.Store.Set(key, value)
	return nil
}

func (f ForMemory) Get(key string) (string, error) {
	return f.Store.GetString(key)
}

func (f ForMemory) Del(key string) error {
	f.Store.Del(key)
	return nil
}
