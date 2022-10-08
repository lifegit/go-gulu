package cacheCount

// Counter
type Counter struct {
	*Count
	max int
}

func NewCounter(max int, count *Count) *Counter {
	return &Counter{
		max:   max,
		Count: count,
	}
}

// 是否频繁
func (c *Counter) IsBusy() bool {
	res, err := c.Get()
	if err != nil {
		return false
	}
	return res >= c.max
}

// 增加一个次数,并且返回增加次数后是否频繁
func (c *Counter) AddCount() bool {
	num, _ := c.Add()

	return num >= int64(c.max)
}

// 删除
func (c *Counter) Destroy() {
	c.Destroy()
}
