package main

import "sync"

type Counter struct {
	m sync.Mutex
	n uint64
}

func (c *Counter) Add(delta uint64) {
	c.m.Lock()
	c.n += delta
	c.m.Unlock()
}

func (c *Counter) Value() uint64 {
	c.m.Lock()
	defer c.m.Unlock()
	return c.n
}

func main() {
	var done = make(chan struct{})
	var c Counter
	for i := 0; i < 100; i++ {
		go func() {
			c.Add(1)
			if c.Value() == 100 {
				done <- struct{}{}

				/*
					select {
					default:
					case done <- struct{}{}:
					}
				*/
			}
		}()
	}
	<-done

}
