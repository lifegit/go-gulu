package goroutinePool

import "sync/atomic"

// GoroutinePool 协程池
type GoroutinePool struct {
	count int              // 协程的数量
	rune  uint32           // 计数器
	close chan bool        // 关闭通道
	end   chan bool        // 数据消费完通道
	ch    chan interface{} // 数据通道
}

func NewGoroutinePool(count int) *GoroutinePool {
	return &GoroutinePool{
		count: count,
		close: make(chan bool),
		end:   make(chan bool),
		ch:    make(chan interface{}, count),
	}
}

func (c *GoroutinePool) Producer(data interface{}) {
	c.ch <- data
}

func (c *GoroutinePool) Customer(callback func(tIndex int, data interface{})) {
	N := -1
	for i := range make([]int, c.count) {
		go func(i int) {
			for {
				select {
				case <-c.close:
					return
				case message := <-c.ch:
					atomic.AddUint32(&c.rune, 1)
					callback(i, message)
					atomic.AddUint32(&c.rune, ^uint32(-N-1))
					if len(c.ch) == 0 && c.rune == 0 {
						c.end <- true
					}
				}
			}
		}(i + 1)
	}
}

func (c *GoroutinePool) Close(data interface{}) {
	c.close <- true
}

func (c *GoroutinePool) Wait() {
	<-c.end
}
