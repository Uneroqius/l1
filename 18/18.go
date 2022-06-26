package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mtx     sync.RWMutex
	counter int64
}

func (c *Counter) Increment() *Counter {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.counter += 1
	return c
}

func (c *Counter) GetCounter() int64 {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.counter
}

func main() {
	counter := Counter{}
	wg := sync.WaitGroup{}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.GetCounter())
}
