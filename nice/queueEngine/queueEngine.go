package queueEngine

import (
	"time"
)

// 开箱即用的生产者消费者队列

type Queue struct {
	Data      interface{}
	FailCount int
}
type Fail struct {
	Wait   time.Duration
	TryLen int
}

type Engine struct {
	producerLen int
	data        chan Queue

	close chan bool
	fail  Fail
}

func NewEngine(chanLen int, fail ...Fail) *Engine {
	// fail
	var f Fail
	if fail != nil {
		f = fail[0]
	} else {
		f = Fail{
			Wait:   time.Second * 3,
			TryLen: 3,
		}
	}
	// queueChan
	queueChan := make(chan Queue)
	if chanLen > 0 {
		queueChan = make(chan Queue, chanLen)
	}
	// return
	return &Engine{
		data:  queueChan,
		close: make(chan bool),
		fail:  f,
	}
}

// 生产
func (e *Engine) Producer(data Queue) {
	ch := make(chan bool)
	e.producer(data, ch)
	<-ch
}

// 生产
func (e *Engine) producer(data Queue, okChan ...chan bool) {
	e.producerLen++
	go func() {
		e.data <- data
		e.producerLen--

		if okChan != nil {
			okChan[0] <- true
		}
	}()
}

// 生产完成
func (e *Engine) ProducerFinish() {
	go func() {
		e.close <- true
	}()
}

// 等待消费数量
func (e *Engine) Len() int {
	return len(e.data)
}

// 消费
func (e *Engine) CustomerWait(callback func(Queue) bool) {
	closeProducer := false

	for {
		select {
		case <-e.close:
			closeProducer = true
		case message := <-e.data:
			if res := callback(message); !res {
				if message.FailCount++; message.FailCount <= e.fail.TryLen {
					// 等待
					time.Sleep(e.fail.Wait)
					// 重新入栈
					e.producer(message)
				}
			}
		}

		if closeProducer && e.producerLen == 0 && len(e.data) == 0 {
			return
		}
	}
}

// 消费异步
func (e *Engine) Customer(callback func(Queue) bool) {
	go func() {
		e.CustomerWait(callback)
	}()
}
